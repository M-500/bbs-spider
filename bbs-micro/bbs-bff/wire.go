//go:build wireinject

package main

import (
	article3 "bbs-micro/bbs-bff/internal/events/article"
	"bbs-micro/bbs-bff/internal/ioc"
	"bbs-micro/bbs-bff/internal/repository"
	"bbs-micro/bbs-bff/internal/repository/article"
	"bbs-micro/bbs-bff/internal/repository/cache"
	"bbs-micro/bbs-bff/internal/repository/dao"
	"bbs-micro/bbs-bff/internal/repository/dao/article_dao"
	"bbs-micro/bbs-bff/internal/service"
	article2 "bbs-micro/bbs-bff/internal/service/article"
	"bbs-micro/bbs-bff/internal/web"
	"bbs-micro/bbs-bff/internal/web/handler"
	"bbs-micro/bbs-bff/internal/web/jwtx"
	"github.com/google/wire"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-11 15:06

func InitWebServer(path string) *App {
	wire.Build(
		ioc.InitConfig,
		ioc.InitLogger,
		ioc.InitRedis,
		ioc.InitDatabase,
		jwtx.NewRedisJWTHandler,
		ioc.InitSaramaClient,
		ioc.InitConsumer,
		ioc.InitSyncProducer,
		article3.NewProducer,

		//article3.NewKafkaConsumer,
		article3.NewInteractiveReadEventBatchConsumer,

		cache.NewArticleCache,
		article_dao.NewGormArticleDao,
		article_dao.NewReadDAO,
		article_dao.NewWriteDAO,
		dao.NewUserDao,
		dao.NewInteractiveDao,

		article.NewArticleRepo,
		article.NewArticleReaderRepo,
		article.NewArtWriterRepo,
		repository.NewUserRepo,
		repository.NewInteractiveRepo,
		cache.NewRedisInteractiveCache,

		article2.NewArticleService,
		service.NewCaptchaService,
		service.NewUserService,
		service.NewInteractiveService,

		handler.NewArticleHandler,
		handler.NewCaptchaHandler,
		handler.NewUserHandler,

		web.NewRouter,

		ioc.InitMiddleware,
		ioc.InitGin,
		// wire.Struct 组装目标结构体的所有字段
		wire.Struct(new(App), "*"),
	)
	return new(App)
}
