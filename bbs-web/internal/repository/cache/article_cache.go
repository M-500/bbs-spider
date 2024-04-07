package cache

import (
	"bbs-web/internal/domain"
	"context"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-07 14:41

type ArticleCache interface {
	GetFirstPage(ctx context.Context, uid int64) ([]domain.Article, error)
	SetFirstPage(ctx context.Context, uid int64, data []domain.Article) error
	DelFirstPage(ctx context.Context, uid int64) error
}
