package ginplus

import (
	"fmt"
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

type Result struct {
	// 这个叫做业务错误码
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func WrapJson[T any](tagFn func(ctx *gin.Context, req T) (Result, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req T
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, Result{
				Code: 0,
				Msg:  "参数错误",
			})
			return
		}
		res, err := tagFn(ctx, req) // 真正的业务逻辑
		if err != nil {
			fmt.Println("执行业务逻辑错误")
			// TODO 这里要记录日志，或者监控啥的
		}
		ctx.JSON(http.StatusOK, res)
		return
	}
}
