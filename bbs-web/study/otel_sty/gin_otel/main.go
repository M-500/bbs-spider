package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
	"go.opentelemetry.io/otel/trace"
	"log"
	"net/http"
	"time"
)

// @Description
// @Author 代码小学生王木木

const (
	serviceName    = "Gin-Jaeger-Demo"
	jaegerEndpoint = "http://192.168.1.52:14268/api/traces"
)

var tracer = otel.Tracer("gin-server")

func main() {
	ctx := context.Background()
	// 初始化并配置 Tracer
	tp, err := initTracer(ctx)
	if err != nil {
		log.Fatalf("initTracer failed, err:%v\n", err)
	}
	defer func() {
		if err2 := tp.Shutdown(ctx); err2 != nil {
			log.Fatal(err2)
		}
	}()
	r := gin.Default()

	// 设置otel中间件
	r.Use(otelgin.Middleware(serviceName))
	// 响应头记录TRACE-ID
	r.Use(func(c *gin.Context) {
		c.Header("Trace-Id", trace.SpanFromContext(c.Request.Context()).SpanContext().TraceID().String())
	})
	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		name := getUser(c.Request.Context(), id)
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"id":   id,
		})
	})
	_ = r.Run(":8080")
}

func initTracer(ctx context.Context) (*sdktrace.TracerProvider, error) {
	//tp, err := newJaegerTraceProvider(ctx)
	tp, err := newTraceProvider()
	if err != nil {
		return nil, err
	}

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}),
	)
	return tp, nil
}

func newJaegerTraceProvider(ctx context.Context) (*sdktrace.TracerProvider, error) {
	// 创建一个使用 HTTP 协议连接本机Jaeger的 Exporter
	exp, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint(jaegerEndpoint),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}
	res, err := resource.New(ctx, resource.WithAttributes(semconv.ServiceName(serviceName)))
	if err != nil {
		return nil, err
	}
	// 创建 Provider
	traceProvider := sdktrace.NewTracerProvider(
		sdktrace.WithResource(res),
		sdktrace.WithSampler(sdktrace.AlwaysSample()), // 采样
		sdktrace.WithBatcher(exp, sdktrace.WithBatchTimeout(time.Second)),
	)
	return traceProvider, nil
}
func newTraceProvider() (*sdktrace.TracerProvider, error) {
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(jaegerEndpoint)))

	if err != nil {
		return nil, err
	}

	traceProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter,
			// Default is 5s. Set to 1s for demonstrative purposes.
			sdktrace.WithBatchTimeout(time.Second)),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
		)),
	)
	return traceProvider, nil
}

func getUser(c context.Context, id string) string {
	// 在需要时将 http.Request 中内置的 `context.Context` 对象传递给 OpenTelemetry API。
	// 可以通过 gin.Context.Request.Context() 获取。
	_, span := tracer.Start(
		c, "getUser", trace.WithAttributes(attribute.String("id", id)),
	)
	defer span.End()
	QueryFromDB(c)
	// mock 业务逻辑
	if id == "7" {
		return "Q1mi"
	}
	return "unknown"
}

func QueryFromDB(ctx context.Context) {
	_, span := tracer.Start(
		ctx, "QueryFromDB",
	)
	defer span.End()
	time.Sleep(time.Second)
}
