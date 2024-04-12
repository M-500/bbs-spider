package integration

import (
	"bbs-micro/bbs-bff/internal/web/vo"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
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

			tc.after(t)
		})
	}
}
