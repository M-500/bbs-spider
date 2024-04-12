package integration

import (
	"bbs-micro/bbs-bff/internal/web/vo"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
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
		{},
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
