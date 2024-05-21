package sharding

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestMigrate(t *testing.T) {
	mysqlConfig := mysql.Config{
		DSN:                       "admin:123456@tcp(192.168.1.52:3306)/sharding_demo?charset=utf8mb4&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         191,                                                                                          // string 类型字段的默认长度
		SkipInitializeWithVersion: false,                                                                                        // 根据版本自动配置

	}
	DB, err := gorm.Open(mysql.New(mysqlConfig))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//生成32张project_log表，具体数量看自己需要了。
	for i := 0; i < 32; i++ {
		tableName := fmt.Sprintf("users_%0*d", 2, i) //表名
		DB.Table(tableName).AutoMigrate(&UserModel{})
	}
}

type UserModel struct {
	gorm.Model
	UserName string
	Age      int
}

func (u UserModel) TableName() string {
	return "users"
}
