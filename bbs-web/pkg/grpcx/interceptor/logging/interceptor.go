//@Author: wulinlin
//@Description: 日志逐渐
//@File:  interceptor
//@Version: 1.0.0
//@Date: 2024/05/02 15:53

package logging

import (
	"bbs-web/pkg/logger"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"runtime"
	"time"
)

type LogInterceptorBuilder struct {
	l logger.Logger
}

func (i *LogInterceptorBuilder) Build() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		// 1. 执行时间 是否错误 调用方法 等等
		startTime := time.Now()
		event := "normal"
		defer func() {
			since := time.Since(startTime) // 执行时间
			fields := []logger.Field{
				logger.Int64("duration", since.Milliseconds()),
				logger.String("type", "unary"),
				logger.String("method", info.FullMethod),
				//logger.String("peer"),
			}
			if rec := recover(); rec != nil {
				switch recType := rec.(type) {
				case error:
					err = recType
				default:
					err = fmt.Errorf("%v", rec)
				}
				stack := make([]byte, 4096)
				stack = stack[:runtime.Stack(stack, true)]
				event = "recover"
				err = status.New(codes.Internal, "panic err"+err.Error()).Err()
			}
			if err != nil {
				st, _ := status.FromError(err)
				fields = append(fields, logger.String("code", st.Code().String()),
					logger.String("code_msg", st.Message()),
					logger.String("event", event),
				)
			}
			i.l.Debug("RPC请求", fields...)
		}()
		resp, err = handler(ctx, req)
		return resp, err
	}
}
