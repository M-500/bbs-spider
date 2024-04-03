package article_dao

import (
	"bbs-web/internal/repository/dao"
	"context"
	"gorm.io/gorm"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-02 12:16

type ReadDAO interface {
	Upsert(ctx context.Context, art dao.ArticleModel) error
}

type readDAO struct {
	db *gorm.DB
}

func (r *readDAO) Upsert(ctx context.Context, art dao.ArticleModel) error {
	//TODO implement me
	panic("implement me")
}

func NewReadDAO(db *gorm.DB) ReadDAO {
	return &readDAO{db: db}
}
