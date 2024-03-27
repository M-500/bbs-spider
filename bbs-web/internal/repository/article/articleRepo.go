package article

import (
	"bbs-web/internal/repository/dao"
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
	artDao dao.ArticleDAO
}

func NewArticleRepo(artDao dao.ArticleDAO) ArticleRepository {
	return &articleRepo{
		artDao: artDao,
	}
}

func (repo *articleRepo) Create(ctx context.Context, art domain.Article) (int64, error) {
	//TODO implement me
	return repo.artDao.Insert(ctx, dao.ArticleModel{
		AuthorId:    0,
		Title:       art.Title,
		Summary:     art.Summary,
		Content:     art.Content,
		ContentType: art.Content,
		Cover:       art.Content,
		Status:      int(art.Status),
	})
}

func (repo *articleRepo) Update(ctx context.Context, art domain.Article) error {

	return repo.artDao.UpdateById(ctx, dao.ArticleModel{
		Model: gorm.Model{
			ID: uint(art.Id),
		},
		AuthorId:    0,
		Title:       art.Title,
		Summary:     art.Summary,
		Content:     art.Content,
		ContentType: art.Content,
		Cover:       art.Content,
		Status:      int(art.Status),
	})
}

func (repo *articleRepo) Sync(ctx context.Context, art domain.Article) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *articleRepo) SyncStatus(ctx context.Context, id int64, author int64, status domain.ArticleStatus) error {
	//TODO implement me
	panic("implement me")
}

func (repo *articleRepo) List(ctx context.Context, uid int64, offset int, limit int) ([]domain.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *articleRepo) GetByID(ctx context.Context, id int64) (domain.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *articleRepo) GetPublishedById(ctx context.Context, id int64) (domain.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *articleRepo) ListPub(ctx context.Context, start time.Time, offset int, limit int) ([]domain.Article, error) {
	//TODO implement me
	panic("implement me")
}
