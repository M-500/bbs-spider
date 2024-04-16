package job

import (
	"bbs-web/internal/domain"
	"bbs-web/internal/service"
	"bbs-web/pkg/logger"
	"context"
	"fmt"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-15 21:13

type Executor interface {
	Name() string

	// Exec
	//  @Description:
	//  @param ctx 整个调度任务的上下文，当有信号的时候，就需要考虑结束执行，交给具体实现去控制 监听ctx.Done信号
	//  @param job
	//  @return error
	Exec(ctx context.Context, job domain.Job) error
}

type localFuncExecutor struct {
	functions map[string]func(ctx context.Context, j domain.Job) error
}

func (l *localFuncExecutor) Name() string {
	return "local_func"
}

func (l *localFuncExecutor) Exec(ctx context.Context, job domain.Job) error {
	fn, ok := l.functions[job.Name]
	if !ok {
		return fmt.Errorf("未找到方法名，检查是否注册, 【%s】", job.Name)
	}
	return fn(ctx, job)
}

// Schedule
// @Description: 调度器
type Schedule struct {
	svc   service.JobService
	lg    logger.Logger
	execs map[string]Executor
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
		// 执行 怎么执行？谁来执行？ 异步执行

		exec, ok := s.execs[j.Exec]
		if !ok {
			// 开发环境的话 可以考虑panic，方便定位问题
			s.lg.Error("未找到对应执行器", logger.Error(err), logger.String("Exec Name", j.Exec))
			continue
		}

		// 要不要释放
		go func() {
			defer func() {
				// 是否要释放？
				err1 := j.CancleFunc()
				if err1 != nil {
					s.lg.Error("释放任务失败",
						logger.Error(err1),
						logger.Int64("job_id", int64(j.ID)))
				}
			}()
			err2 := exec.Exec(ctx, j)
			if err2 != nil {
				// 也可以考虑重试
				s.lg.Error("任务执行失败",
					logger.Error(err2),
					logger.Int64("job_id", int64(j.ID)))
			}
			// 考虑下一次调度
			s.svc.ResetNextTime(ctx, j)
		}()

	}
}
