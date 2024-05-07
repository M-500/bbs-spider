//@Author: wulinlin
//@Description:
//@File:  mq_1_test.go
//@Version: 1.0.0
//@Date: 2024/05/04 22:25

package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestMQV1_Send(t *testing.T) {
	b := MQV1{}
	// 模拟生产者
	go func() {
		for {
			msg := Msg{Context: time.Now().String()}
			err := b.Send(msg)
			if err != nil {
				return
			}
			time.Sleep(time.Second)
		}
	}()
	// 模拟三个消费者，订阅了这个消息队列
	for i := 0; i < 3; i++ {
		go func() {
			chs, err := b.Subscribe(123)
			if err != nil {
				return
			}
			for ch := range chs {
				fmt.Println(ch.Context)
			}
		}()
	}
}
