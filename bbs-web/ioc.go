//@Author: wulinlin
//@Description:
//@File:  ioc
//@Version: 1.0.0
//@Date: 2024/03/27 21:21

package main

import (
	"bbs-web/internal/ioc"
	"bbs-web/internal/repository/article"
	"bbs-web/internal/repository/dao"
	"bbs-web/internal/service"
	"bbs-web/internal/web"
	"bbs-web/internal/web/handler"
	"github.com/gin-gonic/gin"
)

// Injectors from wire.go:

func InitWebServer(path string) *gin.Engine {
	config := ioc.InitConfig(path)
	db := ioc.InitDatabase(config)
	articleDAO := dao.NewArticleDao(db)
	articleRepository := article.NewArticleRepo(articleDAO)
	iArticleService := service.NewArticleService(articleRepository)
	articleHandler := handler.NewArticleHandler(iArticleService)
	router := web.NewRouter(articleHandler)
	v := ioc.InitMiddleware(config)
	engine := ioc.InitGin(router, v)
	return engine
}
