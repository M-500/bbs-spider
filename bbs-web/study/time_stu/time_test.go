package time_stu

import (
	"context"
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
