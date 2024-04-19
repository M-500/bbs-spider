package gormx

import "gorm.io/gorm"

// @Description
// @Author 代码小学生王木木

type DoubleWriteCallback struct {
}

func (c *DoubleWriteCallback) create() func(db *gorm.DB) {
	return func(db *gorm.DB) {
		// 在这里完成双写 但是这里只有一个DB 无法控制双写，无法动态切换
	}
}
