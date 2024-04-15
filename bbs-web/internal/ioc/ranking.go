package ioc

import (
	"bbs-web/internal/job"
	"bbs-web/internal/service"
	"bbs-web/pkg/logger"
	"github.com/robfig/cron/v3"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-15 17:20

func InitRankingJob(svc service.RankingService) *job.RankingJob {
	return job.NewRankingJob(svc, time.Minute)
}

func InitCronJobs(l logger.Logger, rankingJob *job.RankingJob) *cron.Cron {
	res := cron.New(cron.WithSeconds())
	cdb := job.NewCronJonBuilder(l).Build(rankingJob)
	_, err := res.AddJob("0 */3 * * *", cdb) // 每隔3分钟执行一次  参考文章: https://help.aliyun.com/document_detail/133509.html
	if err != nil {
		panic(err)
	}
	return res
}
