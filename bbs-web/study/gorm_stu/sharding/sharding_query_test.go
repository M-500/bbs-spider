package sharding

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/sharding"
	"strconv"
	"testing"
)

// @Description
// @Author 代码小学生王木木

func TestQuery(t *testing.T) {
	dsn := "admin:123456@tcp(192.168.1.52:3306)/sharding_demo?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据版本自动配置

	}
	DB, err := gorm.Open(mysql.New(mysqlConfig))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	DB.Use(sharding.Register(sharding.Config{
		ShardingKey:         "user_name",
		NumberOfShards:      32,
		PrimaryKeyGenerator: sharding.PKSnowflake,
	}, "users"))

	for i := 0; i < 5000; i++ {
		user := UserModel{
			//Model:    gorm.Model{ID: uint(i)},
			UserName: "测试_" + strconv.Itoa(i),
			Age:      i % 78,
		}
		DB.Create(&user)
	}
	//var orders []UserModel
	//DB.Model(&UserModel{}).Where("id = ?", int64(1791082733410193408)).Find(&orders)
	//fmt.Printf("%#v\n", orders)
}
