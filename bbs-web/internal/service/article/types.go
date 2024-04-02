package article

import (
	"bbs-web/internal/domain"
	"context"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-02 9:52
type IArticleService interface {
	Save(ctx context.Context, art domain.Article) (int64, error)
	Withdraw(ctx context.Context, art domain.Article) error
	Publish(ctx context.Context, art domain.Article) (int64, error)
	PublishV1(ctx context.Context, art domain.Article) (int64, error)
	List(ctx context.Context, uid int64, offset int, limit int) ([]domain.Article, error)
	ListPub(ctx context.Context, start time.Time, offset, limit int) ([]domain.Article, error)
	GetById(ctx context.Context, id int64) (domain.Article, error)
	GetByIds(ctx context.Context, biz string, ids []int64) ([]domain.Article, error)
	GetPublishedById(ctx context.Context, id, uid int64) (domain.Article, error)
}
