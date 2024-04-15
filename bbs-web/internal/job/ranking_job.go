package job

import (
	"bbs-web/internal/service"
	"context"
	rlock "github.com/gotomicro/redis-lock" // 分布式锁的实现
	"sync"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-15 15:43

type RankingJob struct {
	svc       service.RankingService
	timeout   time.Duration // 一次运行的时间
	client    *rlock.Client
	rKey      string
	lock      *rlock.Lock
	localLock *sync.Mutex
}

func NewRankingJob(svc service.RankingService, t time.Duration, client *rlock.Client) *RankingJob {
	return &RankingJob{
		svc:       svc,
		timeout:   t,
		client:    client,
		rKey:      "rlock:cron_job:ranking",
		localLock: &sync.Mutex{},
	}
}

func (r *RankingJob) Name() string {
	return "ranking"
}

func (r *RankingJob) Run() error {
	r.localLock.Lock()
	defer r.localLock.Unlock()
	if r.lock == nil {
		// 没有抢到分布式锁，你得试图去抢锁
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		lock, err := r.client.Lock(ctx, r.rKey, r.timeout, &rlock.FixIntervalRetry{
			Interval: time.Millisecond * 100,
			Max:      3,
		}, time.Second) //
		if err != nil { // 重试多次无果，估计是别人抢到了锁，那我就不干了，等下次有机会再干
			return nil
		}
		// 抢锁成功
		r.lock = lock
		// 如何保证一直持有这个锁？
		go func() {
			r.localLock.Lock()
			defer r.localLock.Unlock()
			err2 := lock.AutoRefresh(r.timeout/2, time.Second)
			if err2 != nil {
				// 这里说明退出了续约机制 续约失败了怎么办？？ 不管他
			}
			r.lock = nil
		}()
	}
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	//lock, err := r.client.Lock(ctx, r.rKey, r.timeout+time.Second*60, &rlock.FixIntervalRetry{
	//	Interval: time.Millisecond * 100,
	//	Max:      3,
	//}, time.Second) // 过期时间要比redis的锁时间要久一点
	//if err != nil {
	//	return err
	//}
	//defer func() {
	//	ctx, cancel = context.WithTimeout(context.Background(), time.Minute)
	//	defer cancel()
	//	err1 := lock.Unlock(ctx)
	//	if err1 != nil {
	//		// 记录日志咯 分布式锁释放失败 也可以选择重试几次
	//
	//	}
	//}()
	return r.svc.TopN(ctx)
}
