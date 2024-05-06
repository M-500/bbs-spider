package kafka_p

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"testing"
	"time"
)

func TestBaseConsumer(t *testing.T) {
	// 江湖规矩
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	client, err := sarama.NewClient(Addr, config)
	defer client.Close() // 释放
	if err != nil {
		panic(err)
	}
	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		panic(err)
	}
	partitions, err := consumer.Partitions(TopicName)
	if err != nil {
		panic(err)
	}
	for _, partitionId := range partitions {
		// sarama.OffsetOldest 从最老的地方开始消费
		pConsumer, err := consumer.ConsumePartition(TopicName, partitionId, sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}
		// 开启协程消费
		go func(pc *sarama.PartitionConsumer) {
			defer (*pc).Close()
			for msg := range (*pc).Messages() {
				value := string(msg.Value)
				t.Log("消费成功", value)
			}
		}(&pConsumer)
	}

	time.Sleep(time.Minute)
}

type MyConsumerGroupHandler struct {
}

// 注意这里实现的时候只能用值接收者，不能用指针接收者
func (c MyConsumerGroupHandler) Setup(session sarama.ConsumerGroupSession) error {
	fmt.Println("这里执行了SetUp")
	return nil
}

func (c MyConsumerGroupHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	fmt.Println("这里执行了Cleanup")
	return nil
}

// ConsumeClaim
//
//	@Description:
//	@receiver c
//	@param session
//	@param claim
//	@return error
func (c MyConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// Messages() <-chan *ConsumerMessage  返回一个只读channel
	msgCh := claim.Messages()
	for {
		select {
		case msg, ok := <-msgCh:
			if !ok {
				fmt.Println("读取失败")
			}
			fmt.Println(string(msg.Value), msg.Topic, msg.Offset)
		}
	}
}

func TestBaseConsumerGroup(t *testing.T) {
	groupId := "test_group" // 消费者组ID
	cfg := sarama.NewConfig()
	consumer, err := sarama.NewConsumerGroup(Addr, groupId, cfg)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*15)
	defer cancel()

	err1 := consumer.Consume(ctx, []string{TopicName}, MyConsumerGroupHandler{})
	if err1 != nil {
		panic(err1)
	}
}
