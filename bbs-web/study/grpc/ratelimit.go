package grpc

import (
	"context"
	"github.com/ecodeclub/ekit/queue"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sync"
	"sync/atomic"
	"time"
)

// @Description
// @Author 代码小学生王木木

type CounterLimiter struct {
	cnt       *atomic.Int32
	threshold int32
}

func (c *CounterLimiter) NewServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		cnt := c.cnt.Add(1)
		defer func() {
			c.cnt.Add(-1) // 请求结束时候要将计数器-1
		}()
		if cnt > c.threshold {
			return nil, status.Errorf(codes.ResourceExhausted, "限流")
		}
		return handler(ctx, req)
	}
}

type FixedWindowLimiter struct {
	// 窗口大小
	window time.Duration
	// 上一个窗口的起始时间
	lastStart time.Time
	// 当前窗口的请求数量
	cnt int
	// 阈值
	threshold int

	lock *sync.Mutex
}

func (f *FixedWindowLimiter) NewServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		f.lock.Lock()
		now := time.Now()
		// 要换窗口了！
		if now.Before(f.lastStart.Add(f.window)) {
			f.lastStart = now
			f.cnt = 0
		}
		f.cnt++
		if f.cnt < f.threshold {
			f.lock.Unlock()
			res, err := handler(ctx, req)
			return res, err
		}
		f.lock.Unlock()
		return nil, status.Errorf(codes.ResourceExhausted, "限流")
	}
}

type SlideWindowLimiter struct {
	// 窗口大小
	window time.Duration
	// 请求的时间戳
	queue queue.PriorityQueue[time.Time]
	// 阈值
	threshold int
	lock      *sync.Mutex
}

func (s *SlideWindowLimiter) NewServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		now := time.Now()
		s.lock.Lock()
		if s.queue.Len() < s.threshold {
			_ = s.queue.Enqueue(time.Now())
			s.lock.Unlock()
			return handler(ctx, req)
		}
		// 计算窗口开始的时间点
		windowStartTime := now.Add(-s.window)
		// 一次性把所有过期的都清理掉
		for {
			// 获取队列中最早的一个时间点
			first, _ := s.queue.Peek()
			if first.After(windowStartTime) {
				_, _ = s.queue.Dequeue() // 滑动窗口往后移动左边界
			} else {
				// 退出循环
				break
			}
		}
		if s.queue.Len() < s.threshold {
			_ = s.queue.Enqueue(time.Now())
			s.lock.Unlock()
			return handler(ctx, req)
		}
		s.lock.Unlock()
		return nil, status.Errorf(codes.ResourceExhausted, "限流")
	}
}

type TokenBucketLimiter struct {
	ch      *time.Ticker
	buckets chan struct{}
	// 每个多久发一个令牌？
	interval time.Duration
}

func NewTokenBucketLimiter(interval time.Duration, caps int) *TokenBucketLimiter {
	return &TokenBucketLimiter{
		buckets:  make(chan struct{}, caps),
		interval: interval,
	}
}

func (t *TokenBucketLimiter) NewServerInterceptor() grpc.UnaryServerInterceptor {
	ticker := time.NewTicker(t.interval)
	go func() {
		for _ = range ticker.C {
			select {
			case t.buckets <- struct{}{}:
			default:
				// 桶满了
			}
		}
	}()
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		select {
		case <-t.buckets:
			// 拿到了令牌 直接访问
			return handler(ctx, req)
		case <-ctx.Done():
			// 没有令牌就等，阻塞，直到超时
			return nil, ctx.Err()
			//default:
			//	// 这种写法是没有令牌我就直接溜了，不阻塞
			//	return nil, status.Errorf(codes.ResourceExhausted, "限流")
		}
	}
}

// LeakyBucket
// @Description: 漏桶
type LeakyBucket struct {
	interval time.Duration

	closeCh chan struct{}
}

func (t *LeakyBucket) NewServerInterceptor() grpc.UnaryServerInterceptor {
	ticker := time.NewTicker(t.interval)
	go func() {
		for {
			select {
			case <-t.closeCh:
				return
			case <-ticker.C:
				select {}

			}

		}
	}()
}
