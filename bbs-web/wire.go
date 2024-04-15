//go:build wireinject

package main

import (
	article3 "bbs-web/internal/events/article"
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
	"github.com/google/wire"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 15:57

var rankingServiceSet = wire.NewSet(
	repository.NewRankingRepository,
	cache.NewRankingCache,
	service.NewBatchRankingService,
)

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

		// 任务调度
		rankingServiceSet,
		ioc.InitCronJobs,
		ioc.InitRankingJob,

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
