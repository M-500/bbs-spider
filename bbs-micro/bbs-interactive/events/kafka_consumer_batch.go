package events

import (
	"bbs-micro/bbs-interactive/repository"
	"bbs-micro/pkg/logger"
	"bbs-micro/pkg/saramax"
	"context"
	"github.com/IBM/sarama"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-10 18:49

type InteractiveReadEventBatchConsumer struct {
	client sarama.Client
	l      logger.Logger
	repo   repository.InteractiveRepo
	biz    string
}

func NewInteractiveReadEventBatchConsumer(client sarama.Client,
	l logger.Logger,
	repo repository.InteractiveRepo) Consumer {
	return &InteractiveReadEventBatchConsumer{
		client: client,
		l:      l,
		repo:   repo,
		biz:    "article",
	}
}

func (c *InteractiveReadEventBatchConsumer) Start() error {
	client, err := sarama.NewConsumerGroupFromClient("interactive", c.client)
	if err != nil {
		return err
	}
	go func() {
		err := client.Consume(context.Background(), []string{TopicString}, saramax.NewBatchHandler[ReadEvent](c.l, c.BatchConsume))
		if err != nil {
			c.l.Error("退出消费循环异常", logger.Error(err))
		}
	}()
	return nil
}

func (c *InteractiveReadEventBatchConsumer) BatchConsume(msgs []*sarama.ConsumerMessage,
	events []ReadEvent) error {
	bizs := make([]string, 0, len(events))
	bizIds := make([]int64, 0, len(events))
	for _, evt := range events {
		bizs = append(bizs, "article")
		bizIds = append(bizIds, evt.Aid)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return c.repo.BatchIncrReadCnt(ctx, bizs, bizIds)
}
