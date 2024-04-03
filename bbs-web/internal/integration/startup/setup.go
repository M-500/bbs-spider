package startup

import (
	"bbs-web/internal/ioc"
	"bbs-web/internal/repository"
	"bbs-web/internal/repository/article"
	"bbs-web/internal/repository/dao"
	"bbs-web/internal/repository/dao/article_dao"
	"bbs-web/internal/service"
	article2 "bbs-web/internal/service/article"
	"bbs-web/internal/web"
	"bbs-web/internal/web/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-29 16:25

func InitArticleWebServer(path string) *gin.Engine {
	config := ioc.InitConfig(path)
	db := ioc.InitDatabase(config)
	articleDAO := article_dao.NewArticleDao(db)
	articleRepository := article.NewArticleRepo(articleDAO)
	logger := ioc.InitLogger()
	writeDAO := article_dao.NewWriteDAO(db)
	artWriterRepo := article.NewArtWriterRepo(writeDAO)
	readDAO := article_dao.NewReadDAO(db)
	articleReaderRepository := article.NewArticleReaderRepo(readDAO)
	iArticleService := article2.NewArticleService(articleRepository, logger, artWriterRepo, articleReaderRepository)
	articleHandler := handler.NewArticleHandler(iArticleService, logger)
	iCaptchaSvc := service.NewCaptchaService()
	captchaHandler := handler.NewCaptchaHandler(iCaptchaSvc)
	iUserDao := dao.NewUserDao(db)
	iUserRepo := repository.NewUserRepo(iUserDao)
	iUserService := service.NewUserService(iUserRepo)
	userHandler := handler.NewUserHandler(iUserService, iCaptchaSvc)
	router := web.NewRouter(articleHandler, captchaHandler, userHandler)
	v := ioc.InitMiddleware(config)
	engine := ioc.InitGin(router, v)
	return engine
}

func InitTestDB(cfg *ioc.Config) *gorm.DB {
	config := &gorm.Config{}
	db, err := gorm.Open(mysql.Open(cfg.Database.DSN), config)
	if err != nil {
		panic(err)
	}
	return db
}
