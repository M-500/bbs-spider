package job

import (
	"bbs-web/internal/service"
	"context"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-15 15:43

type RankingJob struct {
	svc     service.RankingService
	timeout time.Duration // 一次运行的时间
}

func NewRankingJob(svc service.RankingService, t time.Duration) *RankingJob {
	return &RankingJob{
		svc:     svc,
		timeout: t,
	}
}

func (r *RankingJob) Name() string {
	return "ranking"
}

func (r *RankingJob) Run() error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	return r.svc.TopN(ctx)
}
