//@Author: wulinlin
//@Description:
//@File:  interactive
//@Version: 1.0.0
//@Date: 2024/04/07 21:26

package cache

import (
	"context"
	"fmt"
	"time"

	_ "embed"
	"github.com/redis/go-redis/v9"
)

//go:embed lua/interative_incr.lua
var luaIncrCnt string

type RedisInteractiveCache interface {
	IncrReadCntIfPresent(ctx context.Context, biz string, bizId int64) error
}

type redisInteractiveCache struct {
	client     redis.Cmdable
	expiration time.Duration
}

func NewRedisInteractiveCache(client redis.Cmdable) RedisInteractiveCache {
	return &redisInteractiveCache{
		client:     client,
		expiration: time.Minute,
	}
}

func (r *redisInteractiveCache) IncrReadCntIfPresent(ctx context.Context, biz string, bizId int64) error {
	return r.client.Eval(ctx, luaIncrCnt, []string{r.key(biz, bizId)}, "read_cnt", 1).Err()
}

func (r *redisInteractiveCache) key(biz string, bizId int64) string {
	return fmt.Sprintf("interactive:%s:%d", biz, bizId)
}
