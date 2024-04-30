package ratelimiter

import (
	"bbs-web/pkg/limiter"
	"bbs-web/study/grpc/hello"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

// @Description
// @Author 代码小学生王木木

type InterceptorBuilder struct {
	limiter limiter.Limiter
	key     string
}

func (b *InterceptorBuilder) BuilderServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		ok, err := b.limiter.Limit(ctx, b.key)
		if err != nil {
			// 你要用保守的 还是用激进的？  这里限流器出错，你要放行吗？
			return nil, err
		}
		if !ok {
			return nil, status.Errorf(codes.ResourceExhausted, "触发限流")
		}
		return handler(ctx, req)
	}
}

func (b *InterceptorBuilder) BuilderClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ok, err := b.limiter.Limit(ctx, b.key)
		if err != nil {
			// 这里限流器出错，你要放行吗？
			return err
		}
		if !ok {
			return status.Errorf(codes.ResourceExhausted, "触发限流")
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func (b *InterceptorBuilder) BuilderServerInterceptorV1() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		if _, ok := req.(*hello.HelloRequest); ok {
			key := "limiter:service:helloService:sayHello"
			ok, err := b.limiter.Limit(ctx, key)
			if err != nil {
				// 你要用保守的 还是用激进的？  这里限流器出错，你要放行吗？
				return nil, err
			}
			if !ok {
				return nil, status.Errorf(codes.ResourceExhausted, "触发限流")
			}
		}
		if strings.HasPrefix(info.FullMethod, "/order-service") {
			ok, err := b.limiter.Limit(ctx, b.key)
			if err != nil {
				// 你要用保守的 还是用激进的？  这里限流器出错，你要放行吗？
				return nil, err
			}
			if !ok {
				return nil, status.Errorf(codes.ResourceExhausted, "触发限流")
			}
		}
		return handler(ctx, req)
	}
}
