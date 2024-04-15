package job

import (
	"bbs-web/pkg/logger"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

// @Description  使用适配器模式，将job封装为corn job
// @Author 代码小学生王木木
// @Date 2024-04-15 15:50

// 这里实现了的是corn Job的接口
type RankingJobAdapter struct {
	j    Job
	logx logger.Logger
	p    prometheus.Summary
}

func NewRankingJobAdapter(j Job, l logger.Logger) *RankingJobAdapter {
	summary := prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "cron_job",
		ConstLabels: map[string]string{
			"name": j.Name(),
		},
	})
	prometheus.MustRegister(summary)
	return &RankingJobAdapter{
		j:    j,
		logx: l,
	}
}

func (r *RankingJobAdapter) Run() {
	start := time.Now()
	defer func() {
		since := time.Since(start).Milliseconds()
		r.p.Observe(float64(since))
	}()
	err := r.j.Run()
	if err != nil {
		r.logx.Error("执行错误", logger.Error(err), logger.String("名字", r.j.Name()))
	}
}
