package cache

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-10 14:50
import (
	"context"
	"fmt"

	_ "embed"
	"github.com/redis/go-redis/v9"

	"bbs-micro/bbs-interactive/domain"
)

var (

	//go:embed lua/incr_cnt.lua
	luaIncrCnt string
)

const (
	readCntKey = "read_cnt"
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
	client  redis.Cmdable
	baseKey string
}

func NewRedisInteractiveCache(c redis.Cmdable) RedisInteractiveCache {
	return &redisInteractiveCache{
		client: c,
	}
}

func (r *redisInteractiveCache) key(biz string, bizId int64) string {
	return fmt.Sprintf("%s:%s:%d", r.baseKey, biz, bizId)
}

// IncrReadCntIfPresent
//
//	@Description: 如果redis存在对应的key 就对其进行+1 操作，核心 HINCRBY 命令(为了并发安全，使用lua脚本)
func (r *redisInteractiveCache) IncrReadCntIfPresent(ctx context.Context, biz string, bizId int64) error {
	return r.client.Eval(ctx, luaIncrCnt,
		[]string{r.key(biz, bizId)}, readCntKey, 1).Err()
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
