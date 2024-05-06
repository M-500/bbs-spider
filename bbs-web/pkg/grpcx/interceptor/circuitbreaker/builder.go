package circuitbreaker

// @Description
// @Author 代码小学生王木木

import (
	"context"
	"github.com/go-kratos/aegis/circuitbreaker"
	"google.golang.org/grpc"
)

type InterceptorBuilder struct {
	breaker circuitbreaker.CircuitBreaker
}

func (b *InterceptorBuilder) BuilderServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		if b.breaker.Allow() == nil {
			// 没有熔断 进行下一步
			resp, err = handler(ctx, req)
			if err != nil {
				// 没有区分业务错误和系统错误
				b.breaker.MarkFailed()
			} else {
				b.breaker.MarkSuccess()
			}
			return resp, err
		}
		// 触发了熔断器
		b.breaker.MarkFailed()
		return nil, err
	}
}
