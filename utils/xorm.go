package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var Engine_mysql *xorm.Engine

func init() {
	var err error
	dsource := "root:123456@tcp(localhost:3306)/taro?charset=utf8"
	Engine_mysql, err = xorm.NewEngine("mysql", dsource)
	if err != nil {
		panic(err)
	}
	err = Engine_mysql.Ping()
	if err != nil {
		panic(err)
	}
}
