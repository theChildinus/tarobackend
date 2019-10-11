package utils

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var Engine_mysql *xorm.Engine

func init() {
	var err error
	dsource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		beego.AppConfig.String("mysqluser"),
		beego.AppConfig.String("mysqlpwd"),
		beego.AppConfig.String("mysqlurl"),
		beego.AppConfig.String("mysqlport"),
		beego.AppConfig.String("mysqldb"),
	)
	Engine_mysql, err = xorm.NewEngine("mysql", dsource)
	if err != nil {
		panic(err)
	}
	err = Engine_mysql.Ping()
	if err != nil {
		panic(err)
	}
}
