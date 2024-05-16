package kafka_demo

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-09 16:37

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestSyncProducer(t *testing.T) {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(addr, cfg)
	cfg.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	assert.NoError(t, err)
	for i := 0; i < 100; i++ {
		u := Student{
			Name: "吴琳琳",
			Age:  i,
		}
		data, err := json.Marshal(u)
		if err != nil {
			panic(err)
		}
		_, _, err = producer.SendMessage(&sarama.ProducerMessage{
			Topic: "read_article",
			//Value: sarama.StringEncoder("这是一条消息"),
			Value: sarama.ByteEncoder(data),
			// 会在生产者和消费者之间传递的
			Headers: []sarama.RecordHeader{
				{
					Key:   []byte("key1"),
					Value: []byte("value1"),
				},
			},
			Metadata: "这是 metadata",
		})
		time.Sleep(time.Second)
	}
}
