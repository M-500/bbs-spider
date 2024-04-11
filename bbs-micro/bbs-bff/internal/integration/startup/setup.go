package startup

import (
	"bbs-micro/bbs-bff/internal/ioc"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-29 16:25

func InitArticleWebServer(path string) *gin.Engine {

	return new(gin.Engine)
}

func InitTestDB(cfg *ioc.Config) *gorm.DB {
	config := &gorm.Config{}
	db, err := gorm.Open(mysql.Open(cfg.Database.DSN), config)
	if err != nil {
		panic(err)
	}
	return db
}
