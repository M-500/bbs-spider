package article_dao

import (
	"bbs-web/internal/repository/dao"
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

func NewWriteDAO(db *gorm.DB) WriteDAO {
	panic("implement me")
}
