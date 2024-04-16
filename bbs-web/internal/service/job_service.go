package service

import (
	"bbs-web/internal/domain"
	"bbs-web/internal/repository"
	"bbs-web/pkg/logger"
	"context"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-15 21:10

type JobService interface {
	// Preempt 抢占
	Preempt(ctx context.Context) (domain.Job, error)
	//PreemptWithCallback(ctx context.Context) (domain.Job, func() error, error)

	Release(ctx context.Context, id int64) error

	ResetNextTime(ctx context.Context, j domain.Job) error
}

// preemptCronJobService
// @Description: 抢占式的分布式任务调度 基于MySQL
type preemptCronJobService struct {
	repo            repository.JobRepository
	refreshInterval time.Duration
	lg              logger.Logger
}

// Preempt
//
//	@Description: 抢占式调度，不返回回调函数
func (p *preemptCronJobService) Preempt(ctx context.Context) (domain.Job, error) {
	j, err := p.repo.Preempt(ctx)
	if err != nil {
		// 抢占失败 还能咋整 妈的
		return domain.Job{}, err
	}

	ticker := time.NewTicker(p.refreshInterval)
	go func() {
		// 续约  定时去续约
		for range ticker.C {
			p.refresh(int64(j.ID))
		}
	}()
	// 是否一直抢占？是不是需要释放呢？暴漏release吗
	j.CancleFunc = func() error {
		// 这里用来释放锁
		ticker.Stop()
		ctx1, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		err = p.Release(ctx1, int64(j.ID))
		if err != nil {
			p.lg.Error("释放 job 失败",
				logger.Error(err),
				logger.Uint("job_id", j.ID))
		}
		return err
	}
	return j, err
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

func (p *preemptCronJobService) refresh(id int64) {
	// 续约： 更新一下更新时间即可
	// 续约失败判定逻辑: 处于running状态，但是update——time在三分钟之前,说明你没有续约。
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err := p.repo.UpdateUpdateTime(ctx, id)
	if err != nil {
		// 续约失败 理论上要通知调用方 续约失败
		p.lg.Error("续约失败",
			logger.Error(err),
			logger.Int64("job_id", id))
	}
}

func (p *preemptCronJobService) ResetNextTime(ctx context.Context, j domain.Job) error {
	t := j.NextTime()
	return p.repo.UpdateNextTime(ctx, int64(j.ID), t)
}
