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
		go func() {
			s.refresh(int64(j.ID))
		}()
		// 是否要释放？
		err = j.CancleFunc()
		if err != nil {
			s.lg.Error("释放任务失败",
				logger.Error(err),
				logger.Int64("job_id", int64(j.ID)))
		}
	}
}

// refresh
//
//	@Description: 续约
//	@receiver s
func (s *Schedule) refresh(id int64) {
	// 续约： 更新一下更新时间即可
	// 续约失败判定逻辑: 处于running状态，但是update——time在三分钟之前,说明你没有续约。

}
