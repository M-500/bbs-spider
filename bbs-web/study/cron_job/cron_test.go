package cron_job

import (
	"github.com/robfig/cron/v3"
	"log"
	"testing"
	"time"
)

// @Description cron定时任务
// @Author 代码小学生王木木
// @Date 2024-03-28 16:23

func TestCronJob(t *testing.T) {
	expr := cron.New(cron.WithSeconds())
	expr.AddJob("@every 5s", myJob{}) // 现成安全的
	expr.AddFunc("@every 6s", func() {
		t.Log("AddFunc 也运行了！")
	})
	expr.Start()
	time.Sleep(time.Second * 20) // 模拟运行10s
	stop := expr.Stop()          // 发出停止信号 不会调度新的任务，但是也不会中断已经调度中的任务
	<-stop.Done()
}

type myJob struct {
}

func (m myJob) Run() {
	log.Printf("运行啦！")
}
