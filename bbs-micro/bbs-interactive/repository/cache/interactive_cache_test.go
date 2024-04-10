package cache

import (
	"bbs-micro/bbs-interactive/domain"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-10 16:32

func TestNewRedisInteractiveCache(t *testing.T) {
	type args struct {
		c redis.Cmdable
	}
	tests := []struct {
		name string
		args args
		want RedisInteractiveCache
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRedisInteractiveCache(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRedisInteractiveCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisInteractiveCache_DecrLikeCntIfPresent(t *testing.T) {
	type fields struct {
		client  redis.Cmdable
		baseKey string
	}
	type args struct {
		ctx   context.Context
		biz   string
		bizId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &redisInteractiveCache{
				client:  tt.fields.client,
				baseKey: tt.fields.baseKey,
			}
			if err := r.DecrLikeCntIfPresent(tt.args.ctx, tt.args.biz, tt.args.bizId); (err != nil) != tt.wantErr {
				t.Errorf("DecrLikeCntIfPresent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_redisInteractiveCache_Get(t *testing.T) {
	type fields struct {
		client  redis.Cmdable
		baseKey string
	}
	type args struct {
		ctx   context.Context
		biz   string
		bizId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Interactive
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &redisInteractiveCache{
				client:  tt.fields.client,
				baseKey: tt.fields.baseKey,
			}
			got, err := r.Get(tt.args.ctx, tt.args.biz, tt.args.bizId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisInteractiveCache_IncrCollCntIfPresent(t *testing.T) {
	type fields struct {
		client  redis.Cmdable
		baseKey string
	}
	type args struct {
		ctx   context.Context
		biz   string
		bizId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &redisInteractiveCache{
				client:  tt.fields.client,
				baseKey: tt.fields.baseKey,
			}
			if err := r.IncrCollCntIfPresent(tt.args.ctx, tt.args.biz, tt.args.bizId); (err != nil) != tt.wantErr {
				t.Errorf("IncrCollCntIfPresent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_redisInteractiveCache_IncrLikeCntIfPresent(t *testing.T) {
	testCase := []struct {
		name string
		mock func(ctrl *gomock.Controller) redis.Cmdable
		// 输入
		ctx   context.Context
		biz   string
		bizId int64
		// 输出
		wantErr error
	}{
		{},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			// 固定格式 start
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			// end
			c := NewRedisInteractiveCache(tc.mock(ctrl))
			err := c.IncrLikeCntIfPresent(tc.ctx, tc.biz, tc.bizId)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}

func Test_redisInteractiveCache_IncrReadCntIfPresent(t *testing.T) {
	type fields struct {
		client  redis.Cmdable
		baseKey string
	}
	type args struct {
		ctx   context.Context
		biz   string
		bizId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &redisInteractiveCache{
				client:  tt.fields.client,
				baseKey: tt.fields.baseKey,
			}
			if err := r.IncrReadCntIfPresent(tt.args.ctx, tt.args.biz, tt.args.bizId); (err != nil) != tt.wantErr {
				t.Errorf("IncrReadCntIfPresent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_redisInteractiveCache_Set(t *testing.T) {
	type fields struct {
		client  redis.Cmdable
		baseKey string
	}
	type args struct {
		ctx   context.Context
		biz   string
		bizId int64
		intr  domain.Interactive
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &redisInteractiveCache{
				client:  tt.fields.client,
				baseKey: tt.fields.baseKey,
			}
			if err := r.Set(tt.args.ctx, tt.args.biz, tt.args.bizId, tt.args.intr); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_redisInteractiveCache_key(t *testing.T) {
	type fields struct {
		client  redis.Cmdable
		baseKey string
	}
	type args struct {
		biz   string
		bizId int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &redisInteractiveCache{
				client:  tt.fields.client,
				baseKey: tt.fields.baseKey,
			}
			if got := r.key(tt.args.biz, tt.args.bizId); got != tt.want {
				t.Errorf("key() = %v, want %v", got, tt.want)
			}
		})
	}
}
