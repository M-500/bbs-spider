package dao

import "gorm.io/gorm"

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-10 12:54

func InitTable(db *gorm.DB) error {
	return db.AutoMigrate(
		&InteractiveModel{},
		&UserLikeBizModel{},
		&UserCollectBizModel{},
	)
}
