package intergration

import (
	"github.com/stretchr/testify/suite" //测试套件
	"gorm.io/gorm"
)

// 集成测试

type intrSvcTestSuite struct {
	suite.Suite

	db *gorm.DB
}
