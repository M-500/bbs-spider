package intergration

import (
	intrv1 "bbs-micro/api/proto/gen/proto/intr/v1"
	"bbs-micro/bbs-interactive/grpc"
	"bbs-micro/bbs-interactive/intergration/dep_setup"
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite" //测试套件
	"gorm.io/gorm"
	"testing"
)

// 集成测试

type intrSvcTestSuite struct {
	suite.Suite
	rdb    redis.Cmdable
	db     *gorm.DB
	server *grpc.InteractiveServiceServer
}

// SetupSuite
//
//	@Description: 开始测试套件之前
//	@receiver s
func (s *intrSvcTestSuite) SetupSuite() {
	s.db = dep_setup.InitTestDB()
	s.rdb = dep_setup.InitRedis()
}

// TearDownTest
//
//	@Description: 结束的时候 通常要清空数据库的所有数据
//	@receiver s
func (s *intrSvcTestSuite) TearDownTest() {

}

func (s *intrSvcTestSuite) TestIncrReadCnt() {
	testCases := []struct {
		name   string
		before func(t *testing.T)
		after  func(t *testing.T)

		biz   string
		bizId int64

		wantErr  error
		wantResp *intrv1.IncrReadCntResponse
	}{
		{
			name: "增加成功,db和redis", // DB 和缓存都有数据
			before: func(t *testing.T) {

			},
			after: func(t *testing.T) {

			},
		},
	}
	// 不同于 AsyncSms 服务，我们不需要 mock，所以创建一个就可以
	// 不需要每个测试都创建
	for _, tc := range testCases {
		s.T().Run(tc.name, func(t *testing.T) {
			tc.before(t)
			resp, err := s.server.IncrReadCnt(context.Background(), &intrv1.IncrReadCntRequest{
				Biz: tc.biz, BizId: tc.bizId,
			})
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantResp, resp)
			tc.after(t)
		})
	}
}
