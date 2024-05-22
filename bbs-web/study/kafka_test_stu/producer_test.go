//@Author: wulinlin
//@Description:
//@File:  producer_test
//@Version: 1.0.0
//@Date: 2024/05/05 12:41

package kafka_test_stu

import (
	"github.com/IBM/sarama"
	"testing"
)

var Addr = []string{"127.0.0.1:9094"}

// TestNormalProducer
//
//	@Description: 普通的生产者程序
//	@param t
func TestNormalProducer(t *testing.T) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	syncProducer, err := sarama.NewSyncProducer(Addr, config)
	if err != nil {
		panic(err)
	}
	pt, offset, err := syncProducer.SendMessage(&sarama.ProducerMessage{
		Topic: "test",
		Value: sarama.StringEncoder("测试发送普通消息"),
	})
	if err != nil {
		panic(err)
	}
	t.Log(pt, offset)
}

func TestAsyncProducer(t *testing.T) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	producer, err := sarama.NewAsyncProducer(Addr, config)
	if err != nil {
		panic(err)
	}
	producer.AsyncClose()
}
