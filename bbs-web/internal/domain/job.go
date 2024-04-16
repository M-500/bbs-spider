package domain

import (
	"github.com/robfig/cron/v3"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-15 21:11

type Job struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Cfg       string
	Name      string // 任务名称

	Exec   string
	Status int // 用来表示状态
	// Cron 表达式
	Expression   string
	NextExecTime int64 // 定时任务的下一次执行的时间

	Version int //MySQL乐观锁 实现并发安全

	CancleFunc func() error
}

func (j Job) NextTime() time.Time {
	c := cron.NewParser(cron.Second | cron.Minute | cron.Hour |
		cron.Dom | cron.Month | cron.Dow | cron.Descriptor)
	s, _ := c.Parse(j.Expression)
	return s.Next(time.Now())
}
