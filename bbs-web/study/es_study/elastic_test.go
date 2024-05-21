package es_study

import (
	"context"
	_ "embed"
	"errors"
	"github.com/olivere/elastic/v7"
	"log"
	"testing"
	"time"
)

var (
	//go:embed user.json
	userIndexJson string
)

// @Description
// @Author 代码小学生王木木

func SetupEsClient() *elastic.Client {
	timeout := 10 * time.Second
	opts := []elastic.ClientOptionFunc{
		elastic.SetURL("http://192.168.1.52:19200"), // 连接参数
		elastic.SetSniff(false),                     // 禁用嗅探
		elastic.SetHealthcheckTimeoutStartup(timeout),
	}
	client, err := elastic.NewClient(opts...)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}
func CreateIndex(client *elastic.Client, indexName, indexJson string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// 判断index是否存在
	exist, err := client.IndexExists(indexName).Do(ctx)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("index 已存在")
	}
	_, err = client.CreateIndex(indexName).BodyJson(indexJson).Do(ctx)
	return err
}

func TestUserEs(t *testing.T) {
	es := SetupEsClient()
	err := CreateIndex(es, "", userIndexJson)
	if err != nil {

	}
}
