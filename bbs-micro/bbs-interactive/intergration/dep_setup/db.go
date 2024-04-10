package dep_setup

import (
	"bbs-micro/bbs-interactive/repository/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-10 12:38

var db *gorm.DB

// InitTestDB 测试的话，不用控制并发。等遇到了并发问题再说
func InitTestDB() *gorm.DB {
	dsn := "admin:123456@tcp(192.168.1.52:3306)/bbs-test"
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
	db, err := gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		panic(err)
	}
	err = dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	//if db == nil {
	//	dsn := "admin:123456@tcp(192.168.1.52:3306)/bbs-test"
	//	sqlDB, err := sql.Open("mysql", dsn)
	//	if err != nil {
	//		panic(err)
	//	}
	//	for {
	//		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//		err = sqlDB.PingContext(ctx)
	//		cancel()
	//		if err == nil {
	//			break
	//		}
	//		log.Println("等待连接 MySQL", err)
	//	}
	//	db, err = gorm.Open(mysql.Open(dsn))
	//	if err != nil {
	//		panic(err)
	//	}
	//	err = dao.InitTable(db)
	//	if err != nil {
	//		panic(err)
	//	}
	//	db = db.Debug()
	//}
	return db
}
