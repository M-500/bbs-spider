package cache

import (
	"bbs-web/internal/domain"
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-15 15:03

type RankinCache interface {
	SetrankingRedisCache(ctx context.Context, articles []domain.Article) error
	GetrankingRedisCache(ctx context.Context) ([]domain.Article, error)
}

const TOPN_REDIS_KEY = "articles:ranking"

type rankingRedisCache struct {
	client redis.Cmdable
}

func NewRankinCache(rds redis.Cmdable) RankinCache {
	return &rankingRedisCache{
		client: rds,
	}
}
func (r *rankingRedisCache) SetrankingRedisCache(ctx context.Context, articles []domain.Article) error {
	// 这个过期时间应当适当长一点，并且超过一次热榜计算任务的总时间  甚至可以用不过期
	for _, art := range articles {
		// 剔除一些不需要的大字段
		art.Content = ""
		art.Summary = ""
	}
	marshalArt, err := json.Marshal(articles)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, TOPN_REDIS_KEY, marshalArt, time.Minute*10).Err()
}

func (r *rankingRedisCache) GetrankingRedisCache(ctx context.Context) ([]domain.Article, error) {
	bytes, err := r.client.Get(ctx, TOPN_REDIS_KEY).Bytes()
	if err != nil {
		return nil, err
	}
	data := bytes
	var res []domain.Article
	err = json.Unmarshal(data, &res)
	return res, err
}

//func (r *rankingRedisCache) key() string {
//	return fmt.Sprintf("%s",TOPN_REDIS_KEY)
//}
