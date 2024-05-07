//@Author: wulinlin
//@Description:
//@File:  kafka_producer
//@Version: 1.0.0
//@Date: 2024/05/05 12:16

package sms_send

import (
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
)

type SMSKafkaProducer struct {
	producer  sarama.SyncProducer
	topicName string
}

func NewSMSKafkaProducer(producer sarama.SyncProducer, topicName string) SMSProducer {
	return &SMSKafkaProducer{
		producer:  producer,
		topicName: topicName,
	}
}

func (s *SMSKafkaProducer) ProducerSMSSendEvent(ctx context.Context, evt SMSSendEvent) error {
	data, err := json.Marshal(evt)
	if err != nil {
		return err
	}
	_, _, err = s.producer.SendMessage(&sarama.ProducerMessage{
		Topic: s.topicName,
		Value: sarama.ByteEncoder(data),
	})
	return err
}
