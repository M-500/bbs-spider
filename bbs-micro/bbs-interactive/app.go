package main

import (
	"bbs-micro/pkg/grpcx"
	"bbs-micro/pkg/saramax"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-10 17:07

type App struct {
	// 所有需要main函数控制启动，关闭的 都会在这里有一个
	server   *grpcx.ServerX
	consumer []saramax.Consumer
}
