package ioc

import (
	"bbs-web/internal/events/article"
	"github.com/IBM/sarama"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-09 12:20

func InitSaramaClient(cfg *Config) sarama.Client {
	scfg := sarama.NewConfig()
	scfg.Producer.Return.Successes = true
	client, err := sarama.NewClient(cfg.KafkaCfg.Brokers, scfg)
	if err != nil {
		panic(err)
	}
	return client
}

func InitSyncProducer(c sarama.Client) sarama.SyncProducer {
	producer, err := sarama.NewSyncProducerFromClient(c)
	if err != nil {
		panic(err)
	}
	return producer
}

func InitConsumer(c article.Consumer) []article.Consumer {
	return []article.Consumer{c}
}
