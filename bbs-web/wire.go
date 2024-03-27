//go:build wireinject

package main

import (
	"bbs-web/internal/ioc"
	"bbs-web/internal/repository/article"
	"bbs-web/internal/repository/dao"
	"bbs-web/internal/service"
	"bbs-web/internal/web"
	"bbs-web/internal/web/handler"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 15:57

func InitWebServer(path string) *gin.Engine {
	wire.Build(
		ioc.InitConfig,
		ioc.InitDatabase,

		dao.NewArticleDao,

		article.NewArticleRepo,

		service.NewArticleService,

		handler.NewArticleHandler,

		web.NewRouter,

		ioc.InitMiddleware,
		ioc.InitGin,
	)
	return new(gin.Engine)
}
