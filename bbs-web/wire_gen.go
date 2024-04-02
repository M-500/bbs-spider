// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"bbs-web/internal/ioc"
	"bbs-web/internal/repository"
	"bbs-web/internal/repository/article"
	"bbs-web/internal/repository/dao"
	"bbs-web/internal/service"
	article2 "bbs-web/internal/service/article"
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
	iArticleService := article2.NewArticleService(articleRepository)
	logger := ioc.InitLogger()
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
