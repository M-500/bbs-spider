package ioc

import (
	"bbs-web/internal/web"
	"bbs-web/pkg/ginplus/middlewares/metric"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 16:27

func InitGin(r *web.Router, mdls []gin.HandlerFunc) *gin.Engine {
	server := gin.Default()
	server.Use(mdls...)
	r.RegisterURL(server)
	return server
}

func InitMiddleware(cfg *Config) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		(&metric.MiddlewareBuilder{
			Namespace:  "wll",
			Subsystem:  "bbs_spider",
			Name:       "gin_http",
			Help:       "统计 GIN 的 HTTP 接口",
			InstanceID: "my-instance-1",
		}).Build(),
		otelgin.Middleware(cfg.ServiceName), // 注入otel 链路追踪
	}
}
