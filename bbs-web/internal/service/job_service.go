package service

import (
	"bbs-web/internal/domain"
	"bbs-web/internal/repository"
	"context"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-15 21:10

type JobService interface {
	// Preempt 抢占
	Preempt(ctx context.Context) (domain.Job, error)
}

// preemptCronJobService
// @Description: 抢占式的分布式任务调度 基于MySQL
type preemptCronJobService struct {
	repo repository.JobRepository
}

func (p *preemptCronJobService) Preempt(ctx context.Context) (domain.Job, error) {
	return p.repo.Preempt(ctx)
}
