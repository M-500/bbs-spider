//@Author: wulinlin
//@Description:
//@File:  interactive
//@Version: 1.0.0
//@Date: 2024/04/07 21:26

package cache

import (
	"bbs-web/internal/domain"
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	_ "embed"
	"github.com/redis/go-redis/v9"
)

//go:embed lua/interative_incr.lua
var luaIncrCnt string

const (
	fieldReadCnt       = "read_cnt"
	fieldLikeCnt       = "like_cnt"
	fieldCollectionCnt = "collect_cnt"
)

var (
	ErrKeyNotExist = errors.New("缓存不存在")
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
	client     redis.Cmdable
	expiration time.Duration
}

func NewRedisInteractiveCache(client redis.Cmdable) RedisInteractiveCache {
	return &redisInteractiveCache{
		client:     client,
		expiration: time.Minute,
	}
}

func (r *redisInteractiveCache) Get(ctx context.Context, biz string, bizId int64) (domain.Interactive, error) {
	// TODO HgetALL 和 HMGet的区别
	result, err := r.client.HGetAll(ctx, r.key(biz, bizId)).Result()
	if err != nil {
		return domain.Interactive{}, err
	}
	// 判空
	if len(result) == 0 {
		// 缓存不存在
		return domain.Interactive{}, ErrKeyNotExist
	}
	var res domain.Interactive
	res.CollectCnt, _ = strconv.ParseInt(result[fieldCollectionCnt], 10, 64)
	res.LikeCnt, _ = strconv.ParseInt(result[fieldLikeCnt], 10, 64)
	res.ReadCnt, _ = strconv.ParseInt(result[fieldReadCnt], 10, 64)

	return res, nil
}

func (r *redisInteractiveCache) Set(ctx context.Context, biz string, bizId int64, intr domain.Interactive) error {
	key := r.key(biz, bizId)
	err := r.client.HSet(ctx, key,
		fieldReadCnt, intr.ReadCnt,
		fieldLikeCnt, intr.LikeCnt,
		fieldCollectionCnt, intr.CollectCnt).Err()
	if err != nil {
		return err
	}
	return r.client.Expire(ctx, key, time.Minute*1).Err()
}

func (r *redisInteractiveCache) IncrReadCntIfPresent(ctx context.Context, biz string, bizId int64) error {
	return r.client.Eval(ctx, luaIncrCnt, []string{r.key(biz, bizId)}, fieldReadCnt, 1).Err()
}

func (r *redisInteractiveCache) IncrLikeCntIfPresent(ctx context.Context, biz string, bizId int64) error {
	return r.client.Eval(ctx, luaIncrCnt, []string{r.key(biz, bizId)}, fieldLikeCnt, 1).Err()
}

func (r *redisInteractiveCache) IncrCollCntIfPresent(ctx context.Context, biz string, bizId int64) error {
	return r.client.Eval(ctx, luaIncrCnt, []string{r.key(biz, bizId)}, fieldCollectionCnt, 1).Err()
}

func (r *redisInteractiveCache) DecrLikeCntIfPresent(ctx context.Context, biz string, bizId int64) error {
	return r.client.Eval(ctx, luaIncrCnt, []string{r.key(biz, bizId)}, fieldLikeCnt, -1).Err()
}

func (r *redisInteractiveCache) key(biz string, bizId int64) string {
	return fmt.Sprintf("interactive:%s:%d", biz, bizId)
}
