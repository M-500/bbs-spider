package main

import (
	event "bbs-web/internal/events/article"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-09 14:04

type App struct {
	server    *gin.Engine
	consumers []event.Consumer
	cron      *cron.Cron
}
