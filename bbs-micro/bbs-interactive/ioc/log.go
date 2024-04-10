package ioc

import (
	"bbs-micro/pkg/logger"
	"go.uber.org/zap"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-01 17:28

func InitLogger() logger.Logger {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer zapLogger.Sync() // 刷新 buffer，保证日志最终会被输出
	return logger.NewZapLogger(zapLogger)
}
