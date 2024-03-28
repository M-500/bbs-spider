package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.Use(MiddlewareDemo())
	// 1. 在main 初始化第三方依赖 mysql redis kafka 通过参数的方式一路传递
	engine.GET("/hello", A)
	engine.GET("/hello2", B())
	engine.GET("/hello3", Wrap(C))
	engine.Run(":10086")
}

func C(ctx *gin.Context) (any, error) {
	if ctx.Request.URL.Path != "aaa" {
		return nil, errors.New("出事啦！")
	}
	return Result{
		Code: 0,
		Msg:  "搞事情",
		Data: []string{"你XX", "唱日出"},
	}, nil
}
func A(ctx *gin.Context) {

}

func D(a, b int) int {
	return a + b
}

func B() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{
			"msg": "不错",
		})
	}
}

func MiddlewareDemo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("欢迎光临")
		ctx.Next()

		fmt.Println("欢迎下次光临")
	}
}

//func WrapReq[T any](fn func(ctx *gin.Context, req T, uc jwt.UserClaims) (Result, error)) gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		// 顺便把 userClaims 也取出来
//	}
//}

type Result struct {
	// 这个叫做业务错误码
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Wrap(fn func(ctx *gin.Context) (any, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res, err := fn(ctx)
		if err != nil {

			ctx.JSON(500, "未知错误")
			return
			// 开始处理 error，其实就是记录一下日志
			//L.Error("处理业务逻辑出错",
			//	logger.String("path", ctx.Request.URL.Path),
			//	// 命中的路由
			//	logger.String("route", ctx.FullPath()),
			//	logger.Error(err))
		}
		//vector.WithLabelValues(strconv.Itoa(res.Code)).Inc()
		ctx.JSON(http.StatusOK, res)
		return
	}
}
