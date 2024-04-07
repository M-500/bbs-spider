package article

import (
	"bbs-web/internal/repository/cache"
	"bbs-web/internal/repository/dao"
	"bbs-web/internal/repository/dao/article_dao"
	"bbs-web/pkg/logger"
	"bbs-web/pkg/utils/zifo/slice"
	"context"
	"gorm.io/gorm"
	"time"

	"bbs-web/internal/domain"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 16:02
type ArticleRepository interface {
	Create(ctx context.Context, art domain.Article) (int64, error)
	Update(ctx context.Context, art domain.Article) error
	// Sync 存储并同步数据
	Sync(ctx context.Context, art domain.Article) (int64, error)

	SyncStatus(ctx context.Context, id int64, author int64, status domain.ArticleStatus) error
	List(ctx context.Context, uid int64, offset int, limit int) ([]domain.Article, error)
	GetByID(ctx context.Context, id int64) (domain.Article, error)
	GetPublishedById(ctx context.Context, id int64) (domain.Article, error)
	ListPub(ctx context.Context, start time.Time, offset int, limit int) ([]domain.Article, error)
	//FindById(ctx context.Context, id int64) domain.Article
}

type articleRepo struct {
	artDao article_dao.ArticleDAO

	// 操作两个dao
	readDAO  article_dao.ReadDAO
	writeDAO article_dao.WriteDAO

	// 引入db
	db *gorm.DB
	// 引入cache
	cache cache.ArticleCache

	l logger.Logger
}

func NewArticleRepo(artDao article_dao.ArticleDAO, c cache.ArticleCache, l logger.Logger) ArticleRepository {
	return &articleRepo{
		artDao: artDao,
		cache:  c,
		l:      l,
	}
}

func (repo *articleRepo) Create(ctx context.Context, art domain.Article) (int64, error) {
	go func() {
		err := repo.cache.DelFirstPage(ctx, art.Author.Id)
		if err != nil {
			repo.l.Error("删除文章列表页缓存失败", logger.Error(err))
		}
	}()
	return repo.artDao.Insert(ctx, dao.ArticleModel{
		AuthorId:    art.Author.Id,
		Title:       art.Title,
		Summary:     art.Summary,
		Content:     art.Content,
		ContentType: art.ContentType,
		Cover:       art.Cover,
		Status:      int(art.Status),
	})
}

func (repo *articleRepo) Update(ctx context.Context, art domain.Article) error {
	go func() {
		err := repo.cache.DelFirstPage(ctx, art.Author.Id)
		if err != nil {
			repo.l.Error("删除文章列表页缓存失败", logger.Error(err))
		}
	}()
	return repo.artDao.UpdateById(ctx, dao.ArticleModel{
		Model: gorm.Model{
			ID: uint(art.Id),
		},
		AuthorId:    art.Author.Id,
		Title:       art.Title,
		Summary:     art.Summary,
		Content:     art.Content,
		ContentType: art.ContentType,
		Cover:       art.Cover,
		Status:      int(art.Status),
	})
}

func (repo *articleRepo) Sync(ctx context.Context, art domain.Article) (int64, error) {
	// 去DAO层处理同步的问题
	return repo.artDao.Sync(ctx, repo.toEntity(art))
}

//func (repo *articleRepo) SyncV3(ctx context.Context, art domain.Article) (int64, error) {
//	// 这么写的话 是谁在控制事务()
//	repo.artDao.Transaction(ctx, func(txDao article_dao.ArticleDAO) error {
//		txDao.Sync(ctx,art)
//	})
//}

// SyncV2
//
//	@Description: 尝试在Repo层解决事务 缺陷就是耦合了DAO层的东西  不太推荐
func (repo *articleRepo) SyncV2(ctx context.Context, art domain.Article) (int64, error) {
	// 开启了一个事务
	tx := repo.db.WithContext(ctx).Begin()
	defer tx.Rollback()
	if tx.Error != nil {
		return 0, nil
	}
	// 利用 tx 来构建Dao
	writeDAO := article_dao.NewWriteDAO(tx)
	readDAO := article_dao.NewReadDAO(tx)
	var (
		id  = art.Id
		err error
	)
	// 制作库`
	articleTmp := repo.toEntity(art)
	if art.Id <= 0 {
		id, err = writeDAO.Insert(ctx, articleTmp)
	} else {
		err = writeDAO.UpdateById(ctx, articleTmp)
	}
	if err != nil {
		//tx.Rollback() // 执行失败 需要回滚
		return 0, err
	}

	art.Id = id
	// 线上库
	err = readDAO.Upsert(ctx, articleTmp)
	if err != nil {
		//tx.Rollback() // 执行失败 需要回滚
		return 0, err
	}
	tx.Commit() // 执行成功 提交
	return id, nil
}

// SyncV1
//
//	@Description: Repo层实现数据同步 非事务实现 无法保证100%
func (repo *articleRepo) SyncV1(ctx context.Context, art domain.Article) (int64, error) {
	var (
		id  = art.Id
		err error
	)
	// 制作库
	articleTmp := repo.toEntity(art)
	if art.Id <= 0 {
		id, err = repo.writeDAO.Insert(ctx, articleTmp)
	} else {
		err = repo.writeDAO.UpdateById(ctx, articleTmp)
	}
	if err != nil {
		return 0, err
	}

	art.Id = id
	// 线上库
	err = repo.readDAO.Upsert(ctx, articleTmp)
	return id, err
}

func (repo *articleRepo) toDomain(src dao.ArticleModel) domain.Article {
	return domain.Article{
		Id:      int64(src.ID),
		Title:   src.Title,
		Content: src.Content,
		Author: domain.Author{
			Id: src.AuthorId,
		},
		Status:      domain.ArticleStatus(src.Status),
		Summary:     src.Summary,
		ContentType: src.ContentType,
		Cover:       src.Cover,
		Ctime:       src.CreatedAt,
		Utime:       src.UpdatedAt,
	}
}

func (repo *articleRepo) toEntity(art domain.Article) dao.ArticleModel {
	return dao.ArticleModel{
		Model: gorm.Model{
			ID: uint(art.Id),
		},
		AuthorId:    art.Author.Id,
		Title:       art.Title,
		Summary:     art.Summary,
		Content:     art.Content,
		ContentType: art.ContentType,
		Cover:       art.Cover,
		Status:      int(art.Status),
	}
}

func (repo *articleRepo) SyncStatus(ctx context.Context, id int64, author int64, status domain.ArticleStatus) error {

	return repo.artDao.SyncStatus(ctx, author, id, uint8(status))
}

func (repo *articleRepo) List(ctx context.Context, uid int64, offset int, limit int) ([]domain.Article, error) {
	// 可以选择缓存第一页的数据 第一页一般都是热点数据
	if offset == 0 && limit <= 100 {
		// 试图从缓存中获取数据
		page, err := repo.cache.GetFirstPage(ctx, uid)
		if err == nil {
			go func() {
				repo.preCache(ctx, page)
			}()
			return page[:limit], nil
		}
		// 否则直接走DB的逻辑
	}
	res, err := repo.artDao.GetByAuthor(ctx, uid, offset, limit)
	if err != nil {
		return nil, err
	}
	data := slice.Map[dao.ArticleModel, domain.Article](res, func(idx int, src dao.ArticleModel) domain.Article {
		temp := repo.toDomain(src)
		temp.Content = "" // 强制不展示列表内容，为了让缓存更轻量，因为列表页不会展示详细信息
		return temp
	})
	// 回写缓存的时候，考虑是set 还是delete
	/**
	如果创作者不太可能高并发，那么就可以直接set，否则应该del

	*/
	if offset == 0 && limit <= 100 {
		// 这里要回写缓存  当然也可以用go func 来执行
		//_ = repo.cache.SetFirstPage(ctx, uid, data)
		go func() {
			err = repo.cache.SetFirstPage(ctx, uid, data)
			if err != nil {
				repo.l.Error("文章列表回写缓存失败", logger.Error(err))
			}
			repo.preCache(ctx, data)
		}()
	}
	return data, nil
}

// 预加载 预测加载，过期时间短一点
func (repo *articleRepo) preCache(ctx context.Context, data []domain.Article) {
	// 是不是要考虑文章内容太长就不缓存
	if len(data) > 0 {
		err := repo.cache.Set(ctx, data[0])
		repo.l.Error("文章缓存预加载失败", logger.Error(err))
	}
}

func (repo *articleRepo) GetByID(ctx context.Context, id int64) (domain.Article, error) {
	res, err := repo.artDao.GetById(ctx, id)
	return repo.toDomain(res), err
}

func (repo *articleRepo) GetPublishedById(ctx context.Context, id int64) (domain.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *articleRepo) ListPub(ctx context.Context, start time.Time, offset int, limit int) ([]domain.Article, error) {
	//TODO implement me
	panic("implement me")
}
