package cache

import (
	"bbs-micro/bbs-interactive/domain"
	"context"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-10 14:50

import (
	"github.com/redis/go-redis/v9"
)

type RedisInteractiveCache interface {
	IncrReadCntIfPresent(ctx context.Context, biz string, bizId int64) error
	IncrLikeCntIfPresent(ctx context.Context, biz string, bizId int64) error
	IncrCollCntIfPresent(ctx context.Context, biz string, bizId int64) error

	DecrLikeCntIfPresent(ctx context.Context, biz string, bizId int64) error

	Get(ctx context.Context, biz string, bizId int64) (domain.Interactive, error)
	Set(ctx context.Context, biz string, bizId int64, intr domain.Interactive) error
}

type redisInteractiveCache struct {
	client redis.Cmdable
}

func NewRedisInteractiveCache(c redis.Cmdable) RedisInteractiveCache {
	return &redisInteractiveCache{
		client: c,
	}
}

func (r *redisInteractiveCache) IncrReadCntIfPresent(ctx context.Context, biz string, bizId int64) error {
	//TODO implement me
	panic("implement me")
}

func (r *redisInteractiveCache) IncrLikeCntIfPresent(ctx context.Context, biz string, bizId int64) error {
	//TODO implement me
	panic("implement me")
}

func (r *redisInteractiveCache) IncrCollCntIfPresent(ctx context.Context, biz string, bizId int64) error {
	//TODO implement me
	panic("implement me")
}

func (r *redisInteractiveCache) DecrLikeCntIfPresent(ctx context.Context, biz string, bizId int64) error {
	//TODO implement me
	panic("implement me")
}

func (r *redisInteractiveCache) Get(ctx context.Context, biz string, bizId int64) (domain.Interactive, error) {
	//TODO implement me
	panic("implement me")
}

func (r *redisInteractiveCache) Set(ctx context.Context, biz string, bizId int64, intr domain.Interactive) error {
	//TODO implement me
	panic("implement me")
}
