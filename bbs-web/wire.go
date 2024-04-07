//go:build wireinject

package main

import (
	"bbs-web/internal/ioc"
	"bbs-web/internal/repository"
	"bbs-web/internal/repository/article"
	"bbs-web/internal/repository/cache"
	"bbs-web/internal/repository/dao"
	"bbs-web/internal/repository/dao/article_dao"
	"bbs-web/internal/service"
	article2 "bbs-web/internal/service/article"
	"bbs-web/internal/web"
	"bbs-web/internal/web/handler"
	"bbs-web/internal/web/jwtx"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 15:57

func InitWebServer(path string) *gin.Engine {
	wire.Build(
		ioc.InitConfig,
		ioc.InitLogger,
		ioc.InitRedis,
		ioc.InitDatabase,
		jwtx.NewRedisJWTHandler,

		cache.NewArticleCache,
		article_dao.NewGormArticleDao,
		article_dao.NewReadDAO,
		article_dao.NewWriteDAO,
		dao.NewUserDao,

		article.NewArticleRepo,
		article.NewArticleReaderRepo,
		article.NewArtWriterRepo,
		repository.NewUserRepo,

		article2.NewArticleService,
		service.NewCaptchaService,
		service.NewUserService,

		handler.NewArticleHandler,
		handler.NewCaptchaHandler,
		handler.NewUserHandler,

		web.NewRouter,

		ioc.InitMiddleware,
		ioc.InitGin,
	)
	return new(gin.Engine)
}
