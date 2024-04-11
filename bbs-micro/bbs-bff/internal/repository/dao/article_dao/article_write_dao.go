package article_dao

import (
	"bbs-micro/bbs-bff/internal/repository/dao"
	"context"
	"gorm.io/gorm"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-02 12:16

type WriteDAO interface {
	Insert(ctx context.Context, art dao.ArticleModel) (int64, error)
	UpdateById(ctx context.Context, art dao.ArticleModel) error
}

type writerDAO struct {
	db *gorm.DB
}

func (w *writerDAO) Insert(ctx context.Context, art dao.ArticleModel) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (w *writerDAO) UpdateById(ctx context.Context, art dao.ArticleModel) error {
	//TODO implement me
	panic("implement me")
}

func NewWriteDAO(db *gorm.DB) WriteDAO {
	return &writerDAO{
		db: db,
	}
}
