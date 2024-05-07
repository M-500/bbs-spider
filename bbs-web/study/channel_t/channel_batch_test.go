package channel_t

import (
	"fmt"
	"testing"
	"time"
)

// @Description
// @Author 代码小学生王木木

var GlbCh = make(chan int, 1000) // 容量为1000的channel

func Producer() {
	i := 0
	for {
		time.Sleep(time.Millisecond * 100)
		i++
		select {
		case GlbCh <- i:
			fmt.Println("生产成功")
		default:
			fmt.Println("被阻塞了")
		}
	}
}

func Consumer() int {
	time.Sleep(time.Millisecond * 200) // 每隔200毫秒 返回一个数据
	return <-GlbCh
}

func TestDemo(t *testing.T) {
	go func() {
		Producer()
	}()
	for {
		fmt.Println(Consumer())
	}
}
