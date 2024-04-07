package service

import "github.com/gin-gonic/gin"

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-07 19:13

type InteractiveService interface {
	IncrReadCnt(ctx *gin.Context, biz string, id int64) error
}
