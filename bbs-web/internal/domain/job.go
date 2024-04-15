package domain

import "time"

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-15 21:11

type Job struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Cfg       string
	Status    int // 用来表示状态

	NextExecTime int64 // 定时任务的下一次执行的时间

	Version int //MySQL乐观锁 实现并发安全

	CancleFunc func() error
}
