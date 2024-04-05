package ginplus

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func Wrap(targetFunc func(ctx *gin.Context) (Result, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res, err := targetFunc(ctx) // 真正的业务逻辑
		if err != nil {
			fmt.Println("执行业务逻辑错误")
			// TODO 这里要记录日志，或者监控啥的
		}
		ctx.JSON(http.StatusOK, res)
		return
	}
}

type Result struct {
	// 这个叫做业务错误码
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func WrapParam[T any](tagFn func(ctx *gin.Context, req T) (Result, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//var req T
		//err := ctx.ShouldBin(&req)
	}
}

func WrapToken[C jwt.Claims](tagFn func(ctx *gin.Context, userToken C) (Result, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		val, ok := ctx.Get("users")
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c, ok := val.(C)
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		res, err := tagFn(ctx, c)
		if err != nil {
			fmt.Println("执行业务逻辑错误")
			// TODO 这里要记录日志，或者监控啥的
			ctx.JSON(http.StatusOK, res)
			return
		}
		res.Msg = "OK"
		ctx.JSON(http.StatusOK, res)
		return
	}
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
			ctx.JSON(http.StatusOK, res)
			return
		}
		res.Msg = "OK"
		ctx.JSON(http.StatusOK, res)
		return
	}
}
