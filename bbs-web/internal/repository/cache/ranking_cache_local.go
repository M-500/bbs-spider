package cache

import (
	"bbs-web/internal/domain"
	"context"
	"errors"
	"github.com/ecodeclub/ekit/syncx/atomicx"
	"time"
)

// @Description 热榜功能本地缓存实现
// @Author 代码小学生王木木
// @Date 2024-04-15 19:03

type RankinLocalCache struct {
	topN *atomicx.Value[[]domain.Article]

	ddl *atomicx.Value[time.Time]

	expiration time.Duration
}

func NewRankinLocalCache() *RankinLocalCache {
	return &RankinLocalCache{
		topN:       atomicx.NewValue[[]domain.Article](),
		ddl:        atomicx.NewValue[time.Time](),
		expiration: time.Minute * 10, // 可以设置到永不过期
	}
}

func (r RankinLocalCache) Set(ctx context.Context, key string, arts []domain.Article, exp time.Duration) error {
	r.topN.Store(arts)
	ddl := time.Now().Add(r.expiration)
	r.ddl.Store(ddl)
	return nil
}

func (r RankinLocalCache) Get(ctx context.Context, key string) ([]domain.Article, error) {
	ddl := r.ddl.Load()
	res := r.topN.Load()
	if len(res) == 0 || ddl.Before(time.Now()) {
		return nil, errors.New("本地缓存未命中")
	}
	return res, nil
}

type item struct {
	arts []domain.Article
	ddl  time.Time
}
