package circuitbreaker

import (
	"context"
	"github.com/go-kratos/aegis/circuitbreaker"
	"google.golang.org/grpc"
	rand2 "math/rand"
	"time"
)

// @Description
// @Author 代码小学生王木木

type InterceptorBuilderV1 struct {
	breaker   circuitbreaker.CircuitBreaker
	threshold int
	rate      int
}

func (b *InterceptorBuilderV1) BuilderServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		if !b.allow() {
			// 触发了熔断
			b.threshold = 0
			time.AfterFunc(time.Minute, func() {
				b.threshold = 1
			})
		}
		rand := rand2.Int31n(100)
		if rand < int32(b.threshold) {
			resp, err = handler(ctx, req)

			if err != nil && b.threshold != 0 {
				// 考虑调小 threshold
				b.threshold--
			} else if b.threshold != 0 {
				// 考虑调大 threshold
				b.threshold++
			}
		}
		return resp, err
	}
}
func (b *InterceptorBuilderV1) allow() bool {
	// 这里就是判断节点是否健康  1. 从Prometheus获取数据 2. 从注册中心获取数据  3.考虑动态检测等等

	return false
}
