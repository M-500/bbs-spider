package integration

import (
	"bbs-web/internal/integration/startup"
	"bbs-web/internal/ioc"
	"bbs-web/internal/repository/dao"
	"bytes"
	"encoding/json"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-29 11:28

// ArticleTestSuite
// @Description: 测试套件  单元测试的组合方式
type ArticleTestSuite struct {
	suite.Suite
	server *gin.Engine
	db     *gorm.DB
}

// 实现 SetupAllSuite 接口
func (s *ArticleTestSuite) SetupSuite() {
	// 在所有测试开始之前，做一些事情
	var configFile = flag.String("config", "../../etc/dev.yaml", "配置文件路径")
	s.server = startup.InitArticleWebServer(*configFile)
	config := ioc.InitConfig(*configFile)
	s.db = startup.InitTestDB(config)

}

func (s *ArticleTestSuite) TearDownTest() {
	// 清空所有的数据
	s.db.Exec("TRUNCATE TABLE articles")
}

func (s *ArticleTestSuite) TestABC() {
	s.T().Log("这里是测试套件")
}

func (s *ArticleTestSuite) TestEdit() {
	t := s.T()
	now := time.Now()
	testcase := []struct {
		name string

		// 集成测试准备数据
		before func(t *testing.T)
		// 集成测试验证数据
		after func(t *testing.T)

		art ArticleReq // 预期的输入

		wantCode int // HTTP响应码

		wantRes Result[int64] // 希望新建文章后 返回文章的ID

		wantErr error
	}{
		{
			name: "新建帖子",
			before: func(t *testing.T) {

			},
			after: func(t *testing.T) {
				// 验证数据库
				var art dao.ArticleModel
				err := s.db.Where("title =?", "我的标题").First(&art).Error
				assert.NoError(t, err)
				assert.True(t, !art.CreatedAt.IsZero())
				assert.True(t, !art.UpdatedAt.IsZero())
				art.CreatedAt = time.Time{}
				art.UpdatedAt = time.Time{}
				assert.Equal(t, dao.ArticleModel{
					Model: gorm.Model{
						ID: 1,
					},
					AuthorId:    0,
					Title:       "我的标题",
					Summary:     "夏测",
					Content:     "搞事情 搞事情 搞大事情",
					ContentType: "blog",
					Cover:       "",
					Status:      1,
				}, art)
			},
			art: ArticleReq{
				Id:          0,
				Title:       "我的标题",
				Content:     "搞事情 搞事情 搞大事情",
				Summary:     "夏测",
				ContentType: "blog",
				Cover:       "",
			},
			wantCode: 200,
			wantErr:  nil,
			wantRes: Result[int64]{
				Code: 0,
				Msg:  "",
				Data: int64(1),
			},
		},
		{
			name: "修改已有的文章",
			before: func(t *testing.T) {
				// 修改之前先插入一条数据
				s.db.Create(dao.ArticleModel{
					Model: gorm.Model{
						ID:        1,
						CreatedAt: now,
						UpdatedAt: now,
					},
					AuthorId:    0,
					Title:       "我的标题",
					Summary:     "搞事情",
					Content:     "阿首都基哦亲闻风丧胆钱啊咯就卡我",
					ContentType: "blog",
					Cover:       "",
					Status:      0,
				})

			},
			after: func(t *testing.T) {
				// 验证数据库
				var art dao.ArticleModel
				err := s.db.Where("id =?", 1).First(&art).Error
				assert.NoError(t, err)
				// 验证更新时间是否成功 这里插入时间也不测试了 妈的 难测
				assert.True(t, art.UpdatedAt.After(now))
				art.UpdatedAt = time.Time{}
				art.CreatedAt = time.Time{}
				assert.Equal(t, dao.ArticleModel{
					Model: gorm.Model{
						ID:        1,
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{},
					},
					AuthorId:    0,
					Title:       "我的标题",
					Summary:     "陈冠希求救",
					Content:     "搞事情 搞事情 搞大事情 微信转账300块",
					ContentType: "blog",
					Cover:       "",
					Status:      1,
				}, art)
			},
			art: ArticleReq{
				Id:          1,
				Title:       "我的标题",
				Content:     "搞事情 搞事情 搞大事情 微信转账300块",
				Summary:     "陈冠希求救",
				ContentType: "blog",
				Cover:       "",
			},
			wantCode: 200,
			wantErr:  nil,
			wantRes: Result[int64]{
				Code: 0,
				Msg:  "",
				Data: int64(1),
			},
		},
	}

	for _, tc := range testcase {
		t.Run(tc.name, func(t *testing.T) {
			// 分成3个部分  1. 构造请求  2.执行逻辑  3.验证结果
			tc.before(t)
			reqBody, err := json.Marshal(tc.art)
			assert.NoError(t, err)
			req, err := http.NewRequest(http.MethodPost, "/articles/edit", bytes.NewBuffer(reqBody))
			require.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")
			// 这里你就可以继续使用 req
			resp := httptest.NewRecorder()
			// 这就是 HTTP 请求进去 GIN 框架的入口。
			// 当你这样调用的时候，GIN 就会处理这个请求
			// 响应写回到 resp 里
			s.server.ServeHTTP(resp, req) // 使用测试套件里的Server对象
			assert.Equal(t, tc.wantCode, resp.Code)
			if resp.Code != 200 {
				return
			}
			require.NoError(t, err)
			var webRes Result[int64]
			err = json.NewDecoder(resp.Body).Decode(&webRes)
			assert.Equal(t, tc.wantRes, webRes)
			tc.after(t)
		})
	}
}

// 预期的输入
type ArticleReq struct {
	Id          int64  `json:"id" binding:"-"`
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
	Summary     string `json:"summary" binding:"-"`
	ContentType string `json:"content_type" binding:"required"`
	Cover       string `json:"cover" binding:"-"`
}

// 预期的返回值 之所以用泛型，是防止any在反序列化的时候出问题
type Result[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

func TestArticle(t *testing.T) {
	suite.Run(t, &ArticleTestSuite{})
}
