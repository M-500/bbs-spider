package article

import (
	"bbs-web/internal/domain"
	"context"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-02 10:02

type ArticleReaderRepository interface {
	// Save 有就更新，没有就新建，即 upsert 的语义
	Save(ctx context.Context, art domain.Article) (int64, error)
}
