//@Author: wulinlin
//@Description: 接口定义
//@File:  types
//@Version: 1.0.0
//@Date: 2024/05/19 13:02

package cache_pool

import (
	"context"
	"time"
)

type Cache interface {
	Get(ctx context.Context, key string) (interface{}, error)
	//
	// Set
	//  @Description: 设置kv 并制定过期时间
	//  @return error
	//
	Set(ctx context.Context, key string, val any, exp time.Duration) error
	//
	// Delete
	//  @Description: 删除某个key-value
	//
	Delete(ctx context.Context, key string) error
	//
	// DeleteAndReturn
	//  @Description: 删除并返回的语义
	//
	DeleteAndReturn(ctx context.Context, key string) (any, error)
}
