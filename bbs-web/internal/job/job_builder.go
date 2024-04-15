package job

import (
	"bbs-web/pkg/logger"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/robfig/cron/v3"
	"strconv"
	"time"
)

// @Description  不用适配器模式，用builder模式
// @Author 代码小学生王木木
// @Date 2024-04-15 16:14

type CronJobBuilder struct {
	p  *prometheus.SummaryVec
	lg logger.Logger
}

func NewCronJonBuilder(lg logger.Logger) *CronJobBuilder {
	pr := prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:      "cron_job",
		Namespace: "bbs_spider",
		Help:      "统计 cron job 的执行情况",
		Subsystem: "cron_job_bbs",
	}, []string{
		"name", "success",
	})
	prometheus.MustRegister(pr)
	return &CronJobBuilder{
		p:  pr,
		lg: lg,
	}
}

func (c *CronJobBuilder) Build(job Job) cron.Job {
	name := job.Name()

	return cronJobFuncAdapter(func() error {
		start := time.Now()

		c.lg.Debug("任务开始",
			logger.String("name", name),
			logger.String("time", start.String()),
		)

		err := job.Run()
		duration := time.Since(start)
		if err != nil {
			c.lg.Error("任务执行失败",
				logger.String("name", name),
				logger.Error(err))
		}
		c.lg.Debug("任务结束",
			logger.String("name", name))

		c.p.WithLabelValues(name,
			strconv.FormatBool(err == nil)).Observe(float64(duration.Milliseconds()))
		return err
	})
}

type cronJobFuncAdapter func() error

func (c cronJobFuncAdapter) Run() {
	_ = c()
}
