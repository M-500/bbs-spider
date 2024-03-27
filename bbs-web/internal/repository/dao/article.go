package dao

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 16:06

type ArticleDAO interface {
	Insert(ctx context.Context, art ArticleModel) (int64, error)
	UpdateById(ctx context.Context, art ArticleModel) error
	GetByAuthor(ctx context.Context, author int64, offset, limit int) ([]ArticleModel, error)
	GetById(ctx context.Context, id int64) (ArticleModel, error)
	GetPubById(ctx context.Context, id int64) (ArticleModel, error)
	Sync(ctx context.Context, art ArticleModel) (int64, error)
	SyncStatus(ctx context.Context, author, id int64, status uint8) error
	ListPub(ctx context.Context, start time.Time, offset int, limit int) ([]ArticleModel, error)
}

type articleDao struct {
	db *gorm.DB
}

func NewArticleDao(db *gorm.DB) ArticleDAO {
	return &articleDao{
		db: db,
	}
}

func (dao *articleDao) Insert(ctx context.Context, art ArticleModel) (int64, error) {
	err := dao.db.WithContext(ctx).Create(&art).Error
	return int64(art.ID), err
}

func (dao *articleDao) UpdateById(ctx context.Context, art ArticleModel) error {
	res := dao.db.WithContext(ctx).
		Where("id = ? AND author_id = ?", art.ID, art.AuthorId).
		Updates(map[string]any{
			"title": art.Title,
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

func (dao *articleDao) GetByAuthor(ctx context.Context, author int64, offset, limit int) ([]ArticleModel, error) {
	//TODO implement me
	panic("implement me")
}

func (dao *articleDao) GetById(ctx context.Context, id int64) (ArticleModel, error) {
	//TODO implement me
	panic("implement me")
}

func (dao *articleDao) GetPubById(ctx context.Context, id int64) (ArticleModel, error) {
	//TODO implement me
	panic("implement me")
}

func (dao *articleDao) Sync(ctx context.Context, art ArticleModel) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (dao *articleDao) SyncStatus(ctx context.Context, author, id int64, status uint8) error {
	//TODO implement me
	panic("implement me")
}

func (dao *articleDao) ListPub(ctx context.Context, start time.Time, offset int, limit int) ([]ArticleModel, error) {
	//TODO implement me
	panic("implement me")
}
