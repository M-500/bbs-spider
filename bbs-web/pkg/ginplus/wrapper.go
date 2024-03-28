package ginplus

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Wrap(targetFunc func(ctx *gin.Context) (any, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 请求之前可以做事情
		a, err := targetFunc(ctx)
		if err != nil {
			// 记录日志，处理Error
		}
		ctx.JSON(http.StatusOK, a)
	}
}
