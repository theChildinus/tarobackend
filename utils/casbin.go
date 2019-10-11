package utils

import (
	"github.com/astaxie/beego"
	"github.com/casbin/casbin"
)

var Enforcer *casbin.Enforcer

func init() {
	Enforcer = casbin.NewEnforcer(beego.AppConfig.String("casbinmodel"), beego.AppConfig.String("casbinpolicy"))
}
