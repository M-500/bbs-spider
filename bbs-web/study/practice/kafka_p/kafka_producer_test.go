package kafka_p

import (
	"github.com/IBM/sarama"
	"strconv"
	"testing"
)

var Addr = []string{"192.168.1.52:9094"}
var TopicName = "FirstTopic"

// TestBaseProducer
//
//	@Description:  普通的生产者
//	@param t
func TestBaseProducer(t *testing.T) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认才会返回
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
	// 指定分区 kafka-console-consumer --topic test --from-beginning --bootstrap-server localhost:9092
	//config.Producer.Partitioner = sarama.NewHashPartitioner()
	//config.Producer.Partitioner = sarama.NewRandomPartitioner
	syncProducer, err := sarama.NewSyncProducer(Addr, config)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 10; i++ {
		// 构造消息
		message := sarama.ProducerMessage{
			Topic: TopicName,
			Value: sarama.StringEncoder("测试消息" + strconv.Itoa(i)),
		}
		p, offset, err := syncProducer.SendMessage(&message)
		if err != nil {
			panic(err)
		}
		t.Log(p, offset)
	}
}
