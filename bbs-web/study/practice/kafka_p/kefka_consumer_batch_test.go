package kafka_p

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"golang.org/x/sync/errgroup"
	"log"
	"testing"
	"time"
)

// @Description  批量消费
// @Author 代码小学生王木木

type BatchConsumerHandler struct {
}

func (b BatchConsumerHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (b BatchConsumerHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (b BatchConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	msgs := claim.Messages()
	batchSize := 10    // 10个一批消费
	var isDone = false // 是否结束
	var eg errgroup.Group
	for {
		fmt.Println("一个批次开始")
		batch := make([]*sarama.ConsumerMessage, 0, batchSize)
		// 一秒超时，如果1s之内没有攒够10个消息，那就不等了 直接消费
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		for i := 0; i < batchSize; i++ {
			select {
			case <-ctx.Done():
				isDone = true
			// 超时了
			case msg, ok := <-msgs:
				if !ok {
					cancel()
					return nil
				}
				// 添加到batch数组中
				batch = append(batch, msg)
			}
		}
		// 到这里说明攒够了一批
		for _, msg := range batch {
			eg.Go(func() error {
				// 模拟消费数据
				log.Printf("模拟消费数据", string(msg.Value))
				return nil
			})
		}

	}
}

func TestConsumerBatch(t *testing.T) {
	groupId := "test_group" // 消费者组ID
	cfg := sarama.NewConfig()
	cfg.Consumer.Return.Errors = true
	consumer, err := sarama.NewConsumerGroup(Addr, groupId, cfg)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*15)
	defer cancel()
	err1 := consumer.Consume(ctx, []string{TopicName}, BatchConsumerHandler{})
	if err1 != nil {
		panic(err1)
	}
}
