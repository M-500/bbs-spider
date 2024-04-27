// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	article2 "bbs-web/internal/events/article"
	"bbs-web/internal/ioc"
	"bbs-web/internal/repository"
	"bbs-web/internal/repository/article"
	"bbs-web/internal/repository/cache"
	"bbs-web/internal/repository/dao"
	"bbs-web/internal/repository/dao/article_dao"
	"bbs-web/internal/service"
	article3 "bbs-web/internal/service/article"
	"bbs-web/internal/web"
	"bbs-web/internal/web/handler"
	"bbs-web/internal/web/jwtx"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitWebServer(path string) *App {
	config := ioc.InitConfig(path)
	db := ioc.InitDatabase(config)
	articleDAO := article_dao.NewGormArticleDao(db)
	cmdable := ioc.InitRedis(config)
	articleCache := cache.NewArticleCache(cmdable)
	logger := ioc.InitLogger()
	iUserDao := dao.NewUserDao(db)
	iUserRepo := repository.NewUserRepo(iUserDao)
	articleRepository := article.NewArticleRepo(articleDAO, articleCache, logger, iUserRepo)
	writeDAO := article_dao.NewWriteDAO(db)
	artWriterRepo := article.NewArtWriterRepo(writeDAO)
	readDAO := article_dao.NewReadDAO(db)
	articleReaderRepository := article.NewArticleReaderRepo(readDAO)
	client := ioc.InitSaramaClient(config)
	syncProducer := ioc.InitSyncProducer(client)
	producer := article2.NewProducer(syncProducer)
	iArticleService := article3.NewArticleService(articleRepository, logger, artWriterRepo, articleReaderRepository, producer)
	interactiveDao := dao.NewInteractiveDao(db)
	redisInteractiveCache := cache.NewRedisInteractiveCache(cmdable)
	interactiveRepo := repository.NewInteractiveRepo(interactiveDao, redisInteractiveCache, logger)
	interactiveService := service.NewInteractiveService(interactiveRepo, redisInteractiveCache)
	articleHandler := handler.NewArticleHandler(iArticleService, interactiveService, logger)
	iCaptchaSvc := service.NewCaptchaService()
	captchaHandler := handler.NewCaptchaHandler(iCaptchaSvc)
	iUserService := service.NewUserService(iUserRepo)
	jwtHandler := jwtx.NewRedisJWTHandler(cmdable)
	userHandler := handler.NewUserHandler(iUserService, iCaptchaSvc, jwtHandler)
	router := web.NewRouter(articleHandler, captchaHandler, userHandler)
	v := ioc.InitMiddleware(config, jwtHandler)
	engine := ioc.InitGin(router, v)
	consumer := article2.NewInteractiveReadEventBatchConsumer(client, logger, interactiveRepo)
	v2 := ioc.InitConsumer(consumer)
	rankinCache := cache.NewRankinCache(cmdable)
	rankingRepository := repository.NewRankingRepository(rankinCache)
	rankingService := service.NewBatchRankingService(iArticleService, interactiveService, rankingRepository)
	rankingJob := ioc.InitRankingJob(rankingService)
	cron := ioc.InitCronJobs(logger, rankingJob)
	app := &App{
		server:    engine,
		consumers: v2,
		cron:      cron,
	}
	return app
}

// wire.go:

var rankingServiceSet = wire.NewSet(repository.NewRankingRepository, cache.NewRankinCache, service.NewBatchRankingService)
