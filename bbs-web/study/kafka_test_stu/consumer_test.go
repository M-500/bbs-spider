//@Author: wulinlin
//@Description:
//@File:  consumer_test
//@Version: 1.0.0
//@Date: 2024/05/05 14:05

package kafka_test_stu

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"testing"
	"time"
)

const GroupID = "test_group"

var topicName = ""

// TestNormalConsumerGroup
//
//	@Description: 消费者组
//	@param t
func TestNormalConsumerGroup(t *testing.T) {
	cfg := sarama.NewConfig()
	// 创建消费者组
	consumer, err := sarama.NewConsumerGroup(Addr, GroupID, cfg)
	if err != nil {
		panic(err)
	}
	// 创建超时响应
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 开始消费
	consumer.Consume(ctx, []string{topicName}, nil)

}

// TestNormalConsumer
//
//	@Description: 普通消费者
//	@param t
func TestNormalConsumer(t *testing.T) {
	cfg := sarama.NewConfig()
	// 连接kafka
	consumer, err := sarama.NewConsumer(Addr, cfg)
	if err != nil {
		fmt.Println("消费者连接kafka失败")
		return
	}
	// 选择要消费的topic，通过topic获取到所有的分区
	partitions, err := consumer.Partitions(topicName)
	if err != nil {
		panic(err)
	}
	for part := range partitions {

	}
}
