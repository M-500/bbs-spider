package dao

import "gorm.io/gorm"

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 17:28

func InitTable(db *gorm.DB) error {
	return db.AutoMigrate(
		&UserMode{},
		&ArticleModel{},
	)
}
