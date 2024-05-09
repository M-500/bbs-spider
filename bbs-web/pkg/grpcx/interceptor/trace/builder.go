//@Author: wulinlin
//@Description:
//@File:  interceptor
//@Version: 1.0.0
//@Date: 2024/05/02 16:16

package trace

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	codes2 "go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

type InterceptorBuilder struct {
	tracer    trace.Tracer
	progrator propagation.TextMapPropagator // 传播器
}

func inject(ctx context.Context, p propagation.TextMapPropagator) context.Context {
	// 先看context有没有元数据
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = metadata.New(map[string]string{})
	}
	// 把元数据返回去ctx
	p.Inject(ctx, GRPCHeaderCarrier(md))
	return metadata.NewOutgoingContext(ctx, md)
}
func (i *InterceptorBuilder) BuildClient() grpc.UnaryClientInterceptor {
	propagator := i.progrator
	if propagator == nil {
		// 如果没有 就拿全局的传播器
		propagator = otel.GetTextMapPropagator()
	}
	tracer := i.tracer
	if tracer == nil {
		tracer = otel.Tracer("pkg/grpcx/interceptor/trace/builder.go")
	}
	attrs := []attribute.KeyValue{
		semconv.RPCSystemKey.String("grpc"),
		attribute.Key("rpc.grpc.kind").String("unary"),
		attribute.Key("rpc.component").String("client"),
	}
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
		ctx, span := tracer.Start(ctx, method,
			trace.WithAttributes(attrs...),
			trace.WithSpanKind(trace.SpanKindClient))
		defer span.End()
		defer func() {
			if err != nil {
				span.RecordError(err)
				//if e := errors.FromError(err); e != nil {
				//
				//}
				span.SetStatus(codes2.Error, err.Error())
			} else {
				span.SetStatus(codes2.Ok, "ok")
			}
			span.End()
		}()
		//inject 过程，要把和 trace有关的链路元数据传递给服务端
		ctx2 := inject(ctx, propagator)
		return invoker(ctx2, method, req, reply, cc, opts...)
	}
}
func (i *InterceptorBuilder) BuildServer() grpc.UnaryServerInterceptor {
	propagator := i.progrator
	if propagator == nil {
		// 这个是全局的
		propagator = otel.GetTextMapPropagator()
	}
	tracer := i.tracer
	if tracer == nil {
		tracer = otel.Tracer("pkg/grpcx/interceptor/trace/builder.go")
	}
	attrs := []attribute.KeyValue{
		semconv.RPCSystemKey.String("grpc"),
		attribute.Key("rpc.grpc.kind").String("unary"),
		attribute.Key("rpc.component").String("server"),
	}
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		ctx = extract(ctx, propagator)
		ctx, span := tracer.Start(ctx, info.FullMethod,
			trace.WithSpanKind(trace.SpanKindServer),
			trace.WithAttributes(attrs...),
		)
		defer span.End()
		span.SetAttributes(
			semconv.RPCMethodKey.String(info.FullMethod),
			//semconv.NetPeerNameKey.String()
			//attribute.Key("net.peer.ip").String()
		)
		defer func() {
			if err != nil {

			} else {
				span.SetStatus(codes2.Code(codes.OK), "OK")
			}
		}()

		return handler(ctx, req)
	}
}

func extract(ctx context.Context, p propagation.TextMapPropagator) context.Context {
	md, ok := metadata.FromIncomingContext(ctx) // 获取到客户端过来的链路元数据
	if !ok {
		md = metadata.New(map[string]string{})
	}
	// 把这个md注入到ctx中
	return p.Extract(ctx, GRPCHeaderCarrier(md))
}

type GRPCHeaderCarrier metadata.MD

func (G GRPCHeaderCarrier) Get(key string) string {
	vals := metadata.MD(G).Get(key)
	if len(vals) > 0 {
		return vals[0]
	}
	return ""
}

func (G GRPCHeaderCarrier) Set(key string, value string) {
	metadata.MD(G).Set(key, value)
}

func (G GRPCHeaderCarrier) Keys() []string {
	keys := make([]string, 0, len(G))
	for k := range metadata.MD(G) {
		keys = append(keys, k)
	}
	return keys
}
