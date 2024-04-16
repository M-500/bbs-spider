package service

import (
	"bbs-web/internal/domain"
	"bbs-web/internal/repository"
	"context"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-15 21:10

type JobService interface {
	// Preempt 抢占
	Preempt(ctx context.Context) (domain.Job, error)
	PreemptWithCallback(ctx context.Context) (domain.Job, func() error, error)

	Release(ctx context.Context, id int64) error
}

// preemptCronJobService
// @Description: 抢占式的分布式任务调度 基于MySQL
type preemptCronJobService struct {
	repo repository.JobRepository
}

// Preempt
//
//	@Description: 抢占式调度，不返回回调函数
func (p *preemptCronJobService) Preempt(ctx context.Context) (domain.Job, error) {
	return p.repo.Preempt(ctx)
}

// PreemptWithCallback
//
//	@Description: 抢占式调度，返回回调方法，用于释放
func (p *preemptCronJobService) PreemptWithCallback(ctx context.Context) (domain.Job, func() error, error) {
	j, err := p.repo.Preempt(ctx)
	if err != nil {
		// 抢占失败 还能咋整 妈的
	}
	j.CancleFunc = func() error {
		// 这里用来释放锁
		ctx1, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		return p.Release(ctx1, int64(j.ID))
	}

	// 是否一直抢占？是不是需要释放呢？暴漏release吗
	return j, j.CancleFunc, err
}

// Release
//
//	@Description: 用于释放锁
//	@receiver p
//	@param ctx
//	@param id
//	@return error
func (p *preemptCronJobService) Release(ctx context.Context, id int64) error {
	return p.repo.Release(ctx, id)
}
