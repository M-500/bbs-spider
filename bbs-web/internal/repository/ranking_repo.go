package repository

import (
	"bbs-web/internal/domain"
	"bbs-web/internal/repository/cache"
	"context"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-15 14:53

type RankingRepository interface {
	ReplaceTopN(ctx context.Context, arts []domain.Article) error
	GetTopN(ctx context.Context) ([]domain.Article, error)
}

// onlyCacheRankinRepo
// @Description: 只有缓存的实现
type onlyCacheRankinRepo struct {
	cache cache.RankinCache
}

func NewRankingRepository(ch cache.RankinCache) RankingRepository {
	return &onlyCacheRankinRepo{
		cache: ch,
	}
}

func (o *onlyCacheRankinRepo) ReplaceTopN(ctx context.Context, arts []domain.Article) error {
	return o.cache.SetRankingCache(ctx, arts)
}

func (o *onlyCacheRankinRepo) GetTopN(ctx context.Context) ([]domain.Article, error) {
	return o.cache.GetRankingCache(ctx)
}

type cacheRankingRepo struct {
	redis cache.RankinCache
	local cache.RankinCache
}

func (c cacheRankingRepo) ReplaceTopN(ctx context.Context, arts []domain.Article) error {
	//TODO implement me
	panic("implement me")
}

func (c cacheRankingRepo) GetTopN(ctx context.Context) ([]domain.Article, error) {
	//TODO implement me
	panic("implement me")
}
