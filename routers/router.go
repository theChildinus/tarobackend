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

	beego.Router("policy/create", &controllers.PolicyController{}, "post:Create")
	beego.Router("policy/list", &controllers.PolicyController{}, "*:List")
	beego.Router("policy/deleteOne", &controllers.PolicyController{}, "post:DeleteOne")
	beego.Router("policy/update", &controllers.PolicyController{}, "post:Update")

	beego.Router("user/create", &controllers.UserController{}, "post:Create")
	beego.Router("user/list", &controllers.UserController{}, "*:List")
	beego.Router("user/deleteOne", &controllers.UserController{}, "post:DeleteOne")
	beego.Router("user/update", &controllers.UserController{}, "post:Update")
	beego.Router("user/register", &controllers.UserController{}, "post:Register")

}
