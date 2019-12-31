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

	beego.Router("/index", &controllers.MainController{}, "*:Index")
	beego.Router("resource/create", &controllers.ResourceController{}, "post:Create")
	beego.Router("resource/list", &controllers.ResourceController{}, "*:List")
	beego.Router("resource/deleteOne", &controllers.ResourceController{}, "post:DeleteOne")
	beego.Router("resource/update", &controllers.ResourceController{}, "post:Update")

	beego.Router("policy/create", &controllers.PolicyController{}, "post:Create")
	beego.Router("policy/list", &controllers.PolicyController{}, "*:List")
	beego.Router("policy/deleteOne", &controllers.PolicyController{}, "post:DeleteOne")
	beego.Router("policy/update", &controllers.PolicyController{}, "post:Update")
	beego.Router("policy/check", &controllers.PolicyController{}, "post:Check")

	beego.Router("user/create", &controllers.UserController{}, "post:Create")
	beego.Router("user/list", &controllers.UserController{}, "*:List")
	beego.Router("user/deleteOne", &controllers.UserController{}, "post:DeleteOne")
	beego.Router("user/update", &controllers.UserController{}, "post:Update")
	beego.Router("user/listNameAndRole", &controllers.UserController{}, "post:ListNameAndRole")
	beego.Router("user/register", &controllers.UserController{}, "post:Register")
	beego.Router("user/downloadCert", &controllers.UserController{}, "post:Download")
	beego.Router("user/verifyIdentity", &controllers.UserController{}, "post:VerifyIdentity")
	beego.Router("user/revoke", &controllers.UserController{}, "post:Revoke")
	beego.Router("user/verifyCert", &controllers.UserController{}, "post:VerifyCert")
	beego.Router("user/login", &controllers.UserController{}, "post:Login")
	beego.Router("user/logout", &controllers.UserController{}, "post:Logout")

	beego.Router("identity/create", &controllers.IdentityController{}, "post:Create")
	beego.Router("identity/list", &controllers.IdentityController{}, "*:List")
	beego.Router("identity/deleteOne", &controllers.IdentityController{}, "post:DeleteOne")
	beego.Router("identity/update", &controllers.IdentityController{}, "post:Update")
	beego.Router("identity/register", &controllers.IdentityController{}, "post:Register")
	beego.Router("identity/enroll", &controllers.IdentityController{}, "post:Enroll")
	beego.Router("identity/revoke", &controllers.IdentityController{}, "post:Revoke")

	beego.Router("enum/getValue", &controllers.EnumController{}, "*:GetValue")
	beego.Router("enum/putValue", &controllers.EnumController{}, "post:PutValue")

}
