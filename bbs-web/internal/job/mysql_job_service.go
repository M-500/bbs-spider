package job

import (
	"bbs-web/internal/service"
	"bbs-web/pkg/logger"
	"context"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-15 21:13

// Schedule
// @Description: 调度器
type Schedule struct {
	svc service.JobService
	lg  logger.Logger
}

// Schedule
//
//	@Description: 这里是一个调度器
func (s *Schedule) Schedule(ctx context.Context) error {
	for {
		j, err := s.svc.Preempt(ctx)
		if err != nil {
			// 不能return 你要机选下一轮继续
			s.lg.Error("抢占任务失败", logger.Error(err))
		}
		// 执行 怎么执行？谁来执行？
		j.CancleFunc()
	}
}
