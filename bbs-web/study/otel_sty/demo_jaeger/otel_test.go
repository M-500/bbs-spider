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

var tracer = otel.Tracer("stu-otel")

// TestJaegerSpan
//
//	@Description: 演示span在函数之间的传递
//	@param t
func TestJaegerSpan(t *testing.T) {
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

	spCtx, span := tracer.Start(ctx, "main")
	// 调用service层方法
	ServiceFun(spCtx, 123)
	span.End()
	time.Sleep(time.Second * 5) // 为了防止结束过快 系统链路信息没来及上报
}

// 假设是server层的方法
func ServiceFun(ctx context.Context, data any) {
	// 1. 获取tracer
	//tr := otel.Tracer("stu-otel") // 这个name要和上面的name一致
	// 2. 基于父span创建span，并指定当前的方法名
	spCtx, span := tracer.Start(ctx, "ServiceFun")
	defer span.End()
	time.Sleep(time.Millisecond * 500) // 模拟Server层代码执行了500毫秒
	DaoFunc(spCtx, data)               // 继续调用dao层的代码
}

// 假设是DAO层的方法
func DaoFunc(ctx context.Context, data any) {
	// 同样的方式来写span
	// 2. 基于父span创建span，并指定当前的方法名
	_, span := tracer.Start(ctx, "DaoFunc")
	defer span.End()
	time.Sleep(time.Second) // 模拟dao层执行了1秒
}
