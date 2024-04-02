package article_dao

import (
	"bbs-web/internal/repository/dao"
	"context"
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 16:06

type ArticleDAO interface {
	Insert(ctx context.Context, art dao.ArticleModel) (int64, error)
	UpdateById(ctx context.Context, art dao.ArticleModel) error
	GetByAuthor(ctx context.Context, author int64, offset, limit int) ([]dao.ArticleModel, error)
	GetById(ctx context.Context, id int64) (dao.ArticleModel, error)
	GetPubById(ctx context.Context, id int64) (dao.ArticleModel, error)
	Sync(ctx context.Context, art dao.ArticleModel) (int64, error)
	Upsert(ctx context.Context, art dao.PublishArticleModels) error
	SyncStatus(ctx context.Context, author, id int64, status uint8) error
	ListPub(ctx context.Context, start time.Time, offset int, limit int) ([]dao.ArticleModel, error)
}

type articleDao struct {
	db *gorm.DB
}

func NewArticleDao(db *gorm.DB) ArticleDAO {
	return &articleDao{
		db: db,
	}
}

func (a *articleDao) Insert(ctx context.Context, art dao.ArticleModel) (int64, error) {
	err := a.db.WithContext(ctx).Create(&art).Error
	return int64(art.ID), err
}

func (a *articleDao) UpdateById(ctx context.Context, art dao.ArticleModel) error {
	now := time.Now()
	res := a.db.WithContext(ctx).Model(&dao.ArticleModel{}).
		Where("id = ? AND author_id = ?", art.ID, art.AuthorId).
		Updates(map[string]any{
			"title": art.Title,
			//"author_id":    art.AuthorId, // 创作者不能修改
			"summary":      art.Summary,
			"content":      art.Content,
			"content_type": art.ContentType,
			"cover":        art.Cover,
			"status":       art.Status,
			"updated_at":   now,
		})

	err := res.Error
	if err != nil {
		return err
	}
	if res.RowsAffected == 0 {
		return errors.New("更新数据失败")
	}
	return err
}

func (a *articleDao) GetByAuthor(ctx context.Context, author int64, offset, limit int) ([]dao.ArticleModel, error) {
	//TODO implement me
	panic("implement me")
}

func (a *articleDao) GetById(ctx context.Context, id int64) (dao.ArticleModel, error) {
	//TODO implement me
	panic("implement me")
}

func (a *articleDao) GetPubById(ctx context.Context, id int64) (dao.ArticleModel, error) {
	//TODO implement me
	panic("implement me")
}

// Upsert
//
//	@Description: GORM实现Upsert的语义 Insert Or Update
func (a *articleDao) Upsert(ctx context.Context, art dao.PublishArticleModels) error {
	now := time.Now()
	// OnConflict的意思是数据冲突了  用MySQL就只需要关注一个地方 DoUpdates
	var err = a.db.Clauses(clause.OnConflict{
		// MySQL只需要关心这个 其他SQL标准的就不一样啦！
		DoUpdates: clause.Assignments(map[string]interface{}{
			"title": art.Title,
			//"author_id":    art.AuthorId, // 创作者不能修改
			"summary":      art.Summary,
			"content":      art.Content,
			"content_type": art.ContentType,
			"cover":        art.Cover,
			"status":       art.Status,
			"updated_at":   now,
		}),
	}).Create(&art).Error // 最终生成的语句 INSERT xxx ON DUPLICATE KEY UPDATE xxx
	return err
}

func (a *articleDao) Sync(ctx context.Context, art dao.ArticleModel) (int64, error) {
	// 这里采用闭包形态操作事务 GORM帮我们管理了事务的生命周期
	var id = int64(art.ID)
	err := a.db.Transaction(func(tx *gorm.DB) error {
		var err error
		txDao := NewArticleDao(tx)
		if id > 0 {
			err = txDao.UpdateById(ctx, art)
		} else {
			id, err = txDao.Insert(ctx, art)
		}
		if err != nil {
			return err
		}
		// 操作线上表了
		return txDao.Upsert(ctx, dao.PublishArticleModels{art})
	})
	return id, err
}

func (a *articleDao) SyncStatus(ctx context.Context, author, id int64, status uint8) error {
	//TODO implement me
	panic("implement me")
}

func (a *articleDao) ListPub(ctx context.Context, start time.Time, offset int, limit int) ([]dao.ArticleModel, error) {
	//TODO implement me
	panic("implement me")
}
