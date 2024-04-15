package job

import (
	"bbs-web/internal/service"
	"context"
	rlock "github.com/gotomicro/redis-lock" // 分布式锁的实现
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-15 15:43

type RankingJob struct {
	svc     service.RankingService
	timeout time.Duration // 一次运行的时间
	client  *rlock.Client
	rKey    string
}

func NewRankingJob(svc service.RankingService, t time.Duration, client *rlock.Client) *RankingJob {
	return &RankingJob{
		svc:     svc,
		timeout: t,
		client:  client,
		rKey:    "rlock:cron_job:ranking",
	}
}

func (r *RankingJob) Name() string {
	return "ranking"
}

func (r *RankingJob) Run() error {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	lock, err := r.client.Lock(ctx, r.rKey, r.timeout+time.Second*60, &rlock.FixIntervalRetry{
		Interval: time.Millisecond * 100,
		Max:      3,
	}, time.Second) // 过期时间要比redis的锁时间要久一点
	if err != nil {
		return err
	}
	defer func() {
		ctx, cancel = context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		_ = lock.Unlock(ctx)
	}()
	return r.svc.TopN(ctx)
}
