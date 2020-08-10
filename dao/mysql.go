package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func MysqlInit(userName string, password string, ip string, dbName string) (err error) {
	dsn := fmt.Sprintf(
		"%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		userName,
		password,
		ip,
		dbName,
	)
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	//	 连接池
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)

	// 表名单数
	DB.SingularTable(true)

	return
}
