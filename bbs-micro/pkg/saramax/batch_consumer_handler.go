package saramax

import (
	"bbs-micro/pkg/logger"
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-09 14:29

type BatchHandler[T any] struct {
	l  logger.Logger
	fn func(msg []*sarama.ConsumerMessage, t []T) error
}

func NewBatchHandler[T any](l logger.Logger, fn func(msg []*sarama.ConsumerMessage, t []T) error) *BatchHandler[T] {
	return &BatchHandler[T]{
		l:  l,
		fn: fn,
	}
}

func (b *BatchHandler[T]) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (b *BatchHandler[T]) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (b *BatchHandler[T]) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// 批量消费
	const batchSize = 10
	msgs := claim.Messages()
	for {
		batch := make([]*sarama.ConsumerMessage, 0, batchSize)
		ts := make([]T, 0, batchSize)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		done := false
		for i := 0; i < batchSize && !done; i++ {
			select {
			case <-ctx.Done():
				// 代表超时
				done = true
			case msg, ok := <-msgs:
				if !ok {
					// 说明消费者被关闭了
					cancel()
					return nil
				}
				//batch = append(batch, msg) // 追加到数组
				var t T
				err := json.Unmarshal(msg.Value, &t)
				if err != nil {
					b.l.Error("反序列化失败", logger.Error(err),
						logger.String("topic", msg.Topic),
						logger.Int32("partition", msg.Partition),
						logger.Int64("offset", msg.Offset),
						logger.Error(err))
					continue
				}
				batch = append(batch, msg)
				ts = append(ts, t)
			}
		}
		cancel()
		// 凑够了一批，处理
		err := b.fn(batch, ts)
		if err != nil {
			b.l.Error("处理消息失败",
				// 把真个 msgs 都记录下来
				logger.Error(err))
		}
		for _, msg := range batch {
			session.MarkMessage(msg, "")
		}
	}

}
