package demo_jaeger

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
	"testing"
	"time"
)

// @Description
// @Author 代码小学生王木木

func TestJaeger(t *testing.T) {
	serviceName := "test-jaeger"
	jaegerEndpoint := "http://192.168.1.52:14268/api/traces"
	exporterJaeger, err := jaeger.New(jaeger.WithCollectorEndpoint(
		jaeger.WithEndpoint(jaegerEndpoint)))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporterJaeger, trace.WithBatchTimeout(time.Second)),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName), // 显示设置服务名称
			attribute.String("env", "prod"),            // 自定义一些参数配置
		)),
	)
	otel.SetTracerProvider(tp)

	tr := otel.Tracer("stu-otel")
	_, span := tr.Start(ctx, "第一个span")
	time.Sleep(time.Second) // 模拟执行了1秒
	span.SetAttributes(attribute.String("key1", "value1"))
	span.SetAttributes(attribute.Bool("key2", false))
	span.SetAttributes(attribute.Int("key3", 45))
	span.SetAttributes(attribute.StringSlice("key4", []string{"demo2", "demo1"}))

	span.AddEvent("这是一个事件")
	span.End()

	time.Sleep(time.Minute) // 为了防止结束过快 系统链路信息没来及上报
}
