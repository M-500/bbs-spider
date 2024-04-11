package article

import (
	"context"
	"encoding/json"

	"github.com/IBM/sarama"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-09 11:20

type KafkaProducer struct {
	producer sarama.SyncProducer
}

func NewProducer(p sarama.SyncProducer) Producer {
	return &KafkaProducer{
		producer: p,
	}
}

func (k *KafkaProducer) ProduceReadEvent(ctx context.Context, evt ReadEvent) error {
	data, err := json.Marshal(evt)
	if err != nil {
		return err
	}
	_, _, err = k.producer.SendMessage(&sarama.ProducerMessage{
		Topic: TopicString,
		Value: sarama.ByteEncoder(data),
	})

	return err
}
