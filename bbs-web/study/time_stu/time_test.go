package time_stu

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-28 16:04

// TestTimer
//
//	@Description: 循环执行  等间隔循环
//	@param t
func TestTimer(t *testing.T) {
	tm := time.NewTicker(time.Second)
	defer tm.Stop()

	for now := range tm.C {
		// 每隔一秒打印一次
		t.Log(now)
	}
}

func TestTicker(t *testing.T) {
	ch := make(chan struct{})
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	go func() {
		// 5秒钟之后开始往通道中写入数据，停止定时任务的运行
		time.Sleep(time.Second * 5)
		ch <- struct{}{}
	}()
	for {
		select {
		case <-ticker.C:
			fmt.Println("执行任务")
		case <-ch:
			// break // 注意这里写break没用，这里的break只是跳出了select，但没有跳出for
			return
		}
	}
}

// TestTimerV1
//
//	@Description: 定时器
//	@param t
func TestTimerV1(t *testing.T) {
	timer := time.NewTimer(time.Second)
	defer timer.Stop()

	go func() {
		for n := range timer.C {
			t.Log(n)
		}
	}()

	time.Sleep(time.Second * 10)
}

// TestTimerV2
//
//	@Description: 控制退出
//	@param t
func TestTimerV2(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	timer := time.NewTimer(time.Second)
	defer timer.Stop()

	for {
		select {
		case now := <-timer.C:
			t.Log(now.String())
		case <-ctx.Done():
			return
		}
	}
}
