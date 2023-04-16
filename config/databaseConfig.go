package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// //定义全局的db对象，我们执行数据库操作主要通过他实现。
var DB *gorm.DB

const (
	//定义连接数据库的参数
	host         = "localhost" // 数据库主机名或IP地址
	port         = "3306"      // 数据库端口号
	username     = "root"      // 数据库用户名
	password     = "200142"    // 数据库密码
	databaseName = "douyin"    // 数据库名
)

func getDatabaseURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, databaseName)
}

func InitDB() (*gorm.DB, error) {
	var err error
	DB, err = gorm.Open("mysql", getDatabaseURL())
	if err != nil {
		return nil, err
	}
	return DB, nil
}
