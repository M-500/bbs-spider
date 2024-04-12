package integration

import (
	"bbs-micro/bbs-bff/internal/integration/startup"
	"bbs-micro/bbs-bff/internal/ioc"
	"bbs-micro/bbs-bff/internal/repository/dao"
	"bbs-micro/bbs-bff/internal/web/jwtx"
	"bbs-micro/bbs-bff/internal/web/vo"
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
// @Date 2024-04-12 12:19

// CollectTestSuite
// @Description: 收藏夹的测试套件
type CollectTestSuite struct {
	suite.Suite
	server *gin.Engine
	db     *gorm.DB
}

func TestCollect(t *testing.T) {
	suite.Run(t, &CollectTestSuite{})
}

func (c *CollectTestSuite) SetupSuite() {
	var configFile = flag.String("config", "../../etc/dev.yaml", "配置文件路径")
	c.server = startup.InitTestWebServer(*configFile)
	c.server.Use(func(context *gin.Context) {
		// 直接设置好登录的token信息
		context.Set("users", &jwtx.UserClaims{
			Id: 3,
		})
		context.Next()
	})
	config := ioc.InitConfig(*configFile)
	c.db = startup.InitTestDB(config)
}
func (c *CollectTestSuite) TearDownTest() {
	c.db.Exec("TRUNCATE TABLE collect")
}

func (c *CollectTestSuite) TestCreate() {
	t := c.T()
	testCase := []struct {
		name string

		before func(t *testing.T)
		after  func(t *testing.T)

		collect vo.CreateCollectReq

		wantCode int

		wantRes Result[int64] // 新建成功后返回对应的ID

		wantErr error
	}{
		{
			name: "创建成功",
			before: func(t *testing.T) {

			},
			after: func(t *testing.T) {
				// 验证数据库
				var col dao.CollectionModle
				err := c.db.Where("id = ?", 1).First(&col).Error
				assert.NoError(t, err)

				// 比对数据
				assert.True(t, !col.CreatedAt.IsZero())
				assert.True(t, !col.UpdatedAt.IsZero())
				col.CreatedAt = time.Time{}
				col.UpdatedAt = time.Time{}
				assert.Equal(t, dao.CollectionModle{
					Model: gorm.Model{
						ID: 1,
					},
					UserId:      3,
					CName:       "篮球",
					Description: "关于篮球的好文章",
					IsPub:       true,
				}, col)
			},
			collect: vo.CreateCollectReq{
				CollectName: "篮球",
				Desc:        "关于篮球的好文章",
				IsPublic:    true,
			},
			wantCode: 200,
			wantErr:  nil,
			wantRes: Result[int64]{
				Code: 0,
				Msg:  "OK",
				Data: int64(1),
			},
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(t)
			// 构造gin
			reqBody, err := json.Marshal(tc.collect)
			assert.NoError(t, err)
			req, err := http.NewRequest(http.MethodPost, "/collect/create", bytes.NewBuffer(reqBody))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "\nBearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTI5ODg5MjcsIklkIjozLCJVc2VyQWdlbnQiOiIifQ.gaK2o7hHlbgJD6wOogG_rBbkbXCzCnqTOQXpjfVuJ1Rnz-tBI5y8QAbKGb8TzeEQDANXYyB1u66u4svzmm5IIA")
			// 这里你就可以继续使用 req
			resp := httptest.NewRecorder()

			c.server.ServeHTTP(resp, req) // 使用测试套件里的Server对象
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
