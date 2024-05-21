package es_study

import (
	"context"
	"fmt"
	es7 "github.com/olivere/elastic/v7"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

// @Description
// @Author 代码小学生王木木

type EsTestSuite struct {
	suite.Suite
	es *es7.Client
}

func (s *EsTestSuite) SetupSuite() {
	client, err := es7.NewClient(
		es7.SetURL("http://192.168.1.52:19200"),
		es7.SetSniff(false),
	)
	require.NoError(s.T(), err)
	s.es = client
}

func (s *EsTestSuite) TestCreateIndexES() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	resp, err := s.es.CreateIndex("use_idx_demo").Body(`
    	{
  "settings": {
    "number_of_shards": 3,
    "number_of_replicas": 1,
    "analysis": {
      "analyzer": {
        "default": {
          "type": "standard"
        }
      }
    }
  },
  "mappings": {
    "properties": {
      "field1": {
        "type": "text",
        "analyzer": "standard"
      },
      "field2": {
        "type": "keyword"
      },
      "field3": {
        "type": "date",
        "format": "yyyy-MM-dd"
      }
    }
  },
  "aliases": {
    "my_index_alias": {}
  }
}
	`).Do(ctx)
	require.NoError(s.T(), err)
	fmt.Println(resp)
}

// TestPutDoc
//
//	@Description: 测试插入数据
//	@receiver s
func (s *EsTestSuite) TestQueryDoc() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	//s.es.Index().Index("use_idx_demo").BodyJson()
	query := es7.NewMatchQuery("title", "Java") // 组合查询 {"match":{"title":{"query":"     Java"}}}
	result, err := s.es.Search().Index("user").Query(query).Do(ctx)
	if err != nil {
		panic(err)
	}
	total := result.Hits.TotalHits.Value
	s.T().Log("结果总数", total)
	for _, v := range result.Hits.Hits {
		s.T().Log(v.Source)
	}
}

func TestBalancerTestSuite(t *testing.T) {
	suite.Run(t, new(EsTestSuite))
}
