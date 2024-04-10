package intergration

import (
	"bbs-micro/bbs-interactive/grpc"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite" //测试套件
	"gorm.io/gorm"
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

}

// TearDownTest
//
//	@Description: 结束的时候 通常要清空数据库的所有数据
//	@receiver s
func (s *intrSvcTestSuite) TearDownTest() {

}
