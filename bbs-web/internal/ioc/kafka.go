package ioc

import "github.com/IBM/sarama"

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
