package metric

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-28 14:13

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
	"time"
)

type MiddlewareBuilder struct {
	Namespace  string
	Subsystem  string
	Name       string
	Help       string
	InstanceID string
}

func (m *MiddlewareBuilder) Build() gin.HandlerFunc {
	// pattern 是指你命中的路由
	// 是指你的 HTTP 的 status
	// path /detail/1
	labels := []string{"method", "pattern", "status"}
	summary := prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: m.Namespace,
		Subsystem: m.Subsystem,
		Name:      m.Name + "_resp_time",
		Help:      m.Help,
		ConstLabels: map[string]string{
			"instance_id": m.InstanceID,
		},
		Objectives: map[float64]float64{
			0.5:   0.01,
			0.9:   0.01,
			0.99:  0.005,
			0.999: 0.0001,
		},
	}, labels)
	prometheus.MustRegister(summary)
	// 用gauge来记录瞬态值，主要是当前活跃的请求数
	gauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: m.Namespace,
		Subsystem: m.Subsystem,
		Name:      m.Name + "_active_req",
		Help:      m.Help,
		ConstLabels: map[string]string{
			"instance_id": m.InstanceID,
		},
	})
	prometheus.MustRegister(gauge)

	total_counter := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: m.Namespace,
		Subsystem: m.Subsystem,
		Name:      m.Name + "_total_req",
		Help:      m.Help,
		ConstLabels: map[string]string{
			"instance_id": m.InstanceID,
		},
	})
	prometheus.MustRegister(total_counter)
	return func(ctx *gin.Context) {
		start := time.Now()
		gauge.Inc()
		defer func() {
			duration := time.Since(start)
			gauge.Dec()
			// 404????
			pattern := ctx.FullPath()
			if pattern == "" {
				pattern = "unknown"
			}
			summary.WithLabelValues(
				ctx.Request.Method,
				pattern,
				strconv.Itoa(ctx.Writer.Status()),
			).Observe(float64(duration.Milliseconds()))
		}()
		// 你最终就会执行到业务里面
		ctx.Next()
	}
}
