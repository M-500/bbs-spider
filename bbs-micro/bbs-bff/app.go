package main

import (
	"bbs-micro/bbs-bff/internal/events/article"
	"github.com/gin-gonic/gin"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-09 14:04

type App struct {
	server    *gin.Engine
	consumers []article.Consumer
}
