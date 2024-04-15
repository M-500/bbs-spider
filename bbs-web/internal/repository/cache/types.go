package cache

import (
	"context"
	"github.com/ecodeclub/ekit"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-15 18:57

type Cache interface {
	Set(ctx context.Context, key string, val any, exp time.Duration) error
	Get(ctx context.Context, key string) ekit.AnyValue
}

type localCache struct {
}

type redisCache struct {
}

type doubleCache struct {
	local Cache
	redis Cache
}
