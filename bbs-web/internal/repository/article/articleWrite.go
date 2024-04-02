package article

import (
	"bbs-web/internal/domain"
	"context"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-02 10:00

type ArticleWriterRepo interface {
	Create(ctx context.Context, art domain.Article) (int64, error)
	Update(ctx context.Context, art domain.Article) error
}
