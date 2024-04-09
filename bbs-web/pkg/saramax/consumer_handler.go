package saramax

import (
	"bbs-web/pkg/logger"
	"encoding/json"
	"github.com/IBM/sarama"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-09 11:33

type Handler[T any] struct {
	l  logger.Logger
	fn func(msg *sarama.ConsumerMessage, t T) error
}

func NewHandler[T any](l logger.Logger, fn func(msg *sarama.ConsumerMessage, t T) error) *Handler[T] {
	return &Handler[T]{
		l:  l,
		fn: fn,
	}
}

func (h *Handler[T]) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h *Handler[T]) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h *Handler[T]) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	var barchSize = 3
	msgs := claim.Messages()
	for msg := range msgs {
		var t T
		err := json.Unmarshal(msg.Value, t)
		if err != nil {
			h.l.Error("反序列化消息失败",
				logger.Error(err),
				logger.Int64("Offset", msg.Offset),
				logger.Int32("Partition", msg.Partition),
				logger.String("Topic", msg.Topic),
			)
			continue
		}
		for i := 0; i < barchSize; i++ {
			err = h.fn(msg, t)
			// 这里要不要统一处理重试
			if err != nil {
				break
			}
			h.l.Error("处理消息失败",
				logger.Error(err),
				logger.Int64("Offset", msg.Offset),
				logger.Int32("Partition", msg.Partition),
				logger.String("Topic", msg.Topic),
			)
		}
		if err != nil {
			h.l.Error("处理消息失败-重试全部失败",
				logger.Error(err),
				logger.Int64("Offset", msg.Offset),
				logger.Int32("Partition", msg.Partition),
				logger.String("Topic", msg.Topic),
			)
		} else {
			session.MarkMessage(msg, "")
		}

	}
	return nil
}
