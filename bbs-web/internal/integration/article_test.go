package integration

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-29 11:28

// ArticleTestSuite
// @Description: 测试套件  单元测试的组合方式
type ArticleTestSuite struct {
	suite.Suite
}

func (s *ArticleTestSuite) TestABC() {
	s.T().Log("这里是测试套件")
}

func (s *ArticleTestSuite) TestEdit() {
	testcase := []struct {
		name string

		// 集成测试准备数据
		before func(t *testing.T)
		// 集成测试验证数据
		after func(t *testing.T)

		art ArticleReq // 预期的输入

		wantCode int // HTTP响应码

		wantRes Result[int64] // 希望新建文章后 返回文章的id

		wantErr error
	}{
		{},
	}

	for _, tc := range testcase {
		s.T().Run(tc.name, func(t *testing.T) {
			// 分成3个部分  1. 构造请求  2.执行逻辑  3.验证结果

		})
	}
}

type ArticleReq struct {
	Id          int64  `json:"id" binding:"-"`
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
	Summary     string `json:"summary" binding:"-"`
	ContentType string `json:"content_type" binding:"required"`
	Cover       string `json:"cover" binding:"-"`
}

type Result[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

func TestArticle(t *testing.T) {
	suite.Run(t, &ArticleTestSuite{})
}
