package ioc

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 11:44

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase(cfg *Config) *gorm.DB {
	config := &gorm.Config{}
	db, err := gorm.Open(mysql.Open(cfg.Database.DSN), config)
	if err != nil {
		panic(err)
	}
	if sqlDB, err := db.DB(); err == nil {
		sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)
		sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	}
	// TODO Prometheus监控

	return db
}
