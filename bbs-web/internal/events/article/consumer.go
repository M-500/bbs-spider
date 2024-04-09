package article

import (
	"bbs-web/internal/repository"
	"bbs-web/pkg/logger"
	"bbs-web/pkg/saramax"
	"context"
	"github.com/IBM/sarama"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-09 11:31

type Consumer interface {
	// ConsumerReadEvent()
	Start() error
}

type KafkaConsumer struct {
	client sarama.Client
	l      logger.Logger
	repo   repository.InteractiveRepo
	biz    string
}

func NewKafkaConsumer(client sarama.Client, l logger.Logger, repo repository.InteractiveRepo) Consumer {
	return &KafkaConsumer{
		client: client,
		l:      l,
		repo:   repo,
		biz:    "article",
	}
}

func (c *KafkaConsumer) Consume(msg *sarama.ConsumerMessage, t ReadEvent) error {
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()
	return c.repo.IncrReadCnt(timeout, c.biz, t.Aid)
}

func (c *KafkaConsumer) Start() error {
	client, err := sarama.NewConsumerGroupFromClient("interactive", c.client)
	if err != nil {
		return err
	}
	go func() {
		err := client.Consume(context.Background(), []string{TopicString}, saramax.NewHandler[ReadEvent](c.l, c.Consume))
		if err != nil {
			c.l.Error("退出消费循环异常", logger.Error(err))
		}
	}()
	return nil
}
