package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"tarobackend/controllers"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("resource/create", &controllers.ResourceController{}, "post:Create")
	beego.Router("resource/list", &controllers.ResourceController{}, "*:List")
	beego.Router("resource/deleteOne", &controllers.ResourceController{}, "post:DeleteOne")
	beego.Router("resource/update", &controllers.ResourceController{}, "post:Update")
}
