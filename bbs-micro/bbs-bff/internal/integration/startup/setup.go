package startup

import (
	article2 "bbs-micro/bbs-bff/internal/events/article"
	"bbs-micro/bbs-bff/internal/ioc"
	"bbs-micro/bbs-bff/internal/repository"
	"bbs-micro/bbs-bff/internal/repository/article"
	"bbs-micro/bbs-bff/internal/repository/cache"
	"bbs-micro/bbs-bff/internal/repository/dao"
	"bbs-micro/bbs-bff/internal/repository/dao/article_dao"
	"bbs-micro/bbs-bff/internal/service"
	article3 "bbs-micro/bbs-bff/internal/service/article"
	"bbs-micro/bbs-bff/internal/web"
	"bbs-micro/bbs-bff/internal/web/handler"
	"bbs-micro/bbs-bff/internal/web/jwtx"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-29 16:25

func InitTestWebServer(path string) *gin.Engine {
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
	interactiveServiceClient := ioc.InitInterGRPCClient(interactiveService, config)
	articleHandler := handler.NewArticleHandler(iArticleService, interactiveServiceClient, logger)
	iCaptchaSvc := service.NewCaptchaService()
	captchaHandler := handler.NewCaptchaHandler(iCaptchaSvc)
	iUserService := service.NewUserService(iUserRepo)
	jwtHandler := jwtx.NewRedisJWTHandler(cmdable)
	userHandler := handler.NewUserHandler(iUserService, iCaptchaSvc, jwtHandler)
	iCollectDAO := dao.NewCollectDao(db)
	iCollectRepo := repository.NewCollectRepo(iCollectDAO)
	iCollectService := service.NewCollectService(iCollectRepo)
	collectHandler := handler.NewCollectHandler(iCollectService)
	router := web.NewRouter(articleHandler, captchaHandler, userHandler, collectHandler)
	v := ioc.InitMiddleware(config, jwtHandler)
	return ioc.InitGin(router, v)
}

func InitTestDB(cfg *ioc.Config) *gorm.DB {
	config := &gorm.Config{}
	db, err := gorm.Open(mysql.Open(cfg.Database.DSN), config)
	if err != nil {
		panic(err)
	}
	return db
}
