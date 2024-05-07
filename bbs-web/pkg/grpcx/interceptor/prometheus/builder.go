//@Author: wulinlin
//@Description:
//@File:  builder
//@Version: 1.0.0
//@Date: 2024/05/02 18:56

package prometheus

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"strings"
	"time"
)

type InterceptorBuilder struct {
	Namespace string
	Subsystem string
}

func (b *InterceptorBuilder) BuilderServer() grpc.UnaryServerInterceptor {
	sumary := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: b.Namespace,
			Subsystem: b.Subsystem,
			Name:      "server_handler_seconds",
			Objectives: map[float64]float64{
				0.5:   0.01,
				0.9:   0.01,
				0.95:  0.01,
				0.99:  0.001,
				0.999: 0.0001,
			},
		}, []string{"type", "service", "method", "peer", "code"})
	prometheus.MustRegister(sumary)
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		start := time.Now()
		defer func() {
			s, m := b.splitMethodName(info.FullMethod)
			since := float64(time.Since(start))
			if err == nil {
				sumary.WithLabelValues("unary", s, m, "", "ok").Observe(since)
			} else {
				st, ok := status.FromError(err)
				if !ok {

				}
				sumary.WithLabelValues("unary", s, m, "", st.Code().String()).Observe(since)
			}

		}()
		resp, err = handler(ctx, req)
		return
	}
}

func (b *InterceptorBuilder) splitMethodName(fullMethodName string) (string, string) {
	fullMethodName = strings.TrimPrefix(fullMethodName, "/")
	if i := strings.Index(fullMethodName, "/"); i >= 0 {
		return fullMethodName[:i], fullMethodName[i+1:]
	}
	return "unknown", "unknown"
}
