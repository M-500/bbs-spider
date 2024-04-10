package dep_setup

import "bbs-micro/pkg/logger"

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-10 17:26

func InitLog() logger.Logger {
	return logger.NewNoOpLogger()
}
