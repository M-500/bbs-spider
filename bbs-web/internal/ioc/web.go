package ioc

import (
	"bbs-web/internal/web"
	"github.com/gin-gonic/gin"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 16:27

func InitGin(r *web.Router) *gin.Engine {

	server := gin.Default()
	r.RegisterURL(server)
	return server
}
