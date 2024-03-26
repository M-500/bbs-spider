package ioc

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 11:44

import (
	"bbs-web/internal/repository/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"gorm.io/plugin/prometheus"
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

	err = db.Use(prometheus.New(prometheus.Config{
		DBName:          cfg.ServiceName,
		RefreshInterval: 15,    // 插件采集数据的频率
		StartServer:     false, // 无需重新启动
		MetricsCollector: []prometheus.MetricsCollector{
			&prometheus.MySQL{
				VariableNames: []string{"thread_running"},
			},
		},
	}))

	err = dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return db
}
