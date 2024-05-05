//@Author: wulinlin
//@Description:
//@File:  consumer
//@Version: 1.0.0
//@Date: 2024/05/05 12:24

package sms_send

import (
	"bbs-web/pkg/logger"
	"bbs-web/pkg/saramax"
	"context"
	"github.com/IBM/sarama"
)

type SMSConsumer interface {
	Start() error
}

type KafkaSMSConsumer struct {
	client     sarama.Client // sarama客户端
	smsService SMSService
	l          logger.Logger
	groupId    string // 消费者组ID
	topicName  string
}

func (k *KafkaSMSConsumer) Start() error {
	client, err := sarama.NewConsumerGroupFromClient(k.groupId, k.client)
	if err != nil {
		return err
	}
	go func() {
		err := client.Consume(context.Background(), []string{k.topicName}, saramax.NewHandler[SMSSendEvent](k.l, k.SendSMS))
		if err != nil {
			// 只能记录日志啦
		}
	}()
	return nil
}

func (k *KafkaSMSConsumer) SendSMS(msg *sarama.ConsumerMessage, t SMSSendEvent) error {
	//TODO implement me
	panic("implement me")
}

func NewKafkaSMSConsumer(client sarama.Client, smsService SMSService, gID string) *KafkaSMSConsumer {
	return &KafkaSMSConsumer{
		client:     client,
		smsService: smsService,
		groupId:    gID,
	}
}
