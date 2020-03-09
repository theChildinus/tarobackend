package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"tarobackend/models"
	pb "tarobackend/proto"
	"tarobackend/services"
	"tarobackend/utils"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) List() {
	var req services.UserReq
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		utils.BuildJsonResp(c, "Error", "Json Parse Error")
		return
	}
	list, count, err := services.ListUser(&req)
	if err != nil {
		logs.Error("List User Error", err.Error())
		utils.BuildJsonResp(c, "Error", "List User Error")
		return
	}
	c.Data["json"] = &services.UserResp{
		List:  list,
		Count: count,
	}
	c.ServeJSON()
	return
}

func (c *UserController) Create() {
	var m models.TaroUser
	var err error
	var ret bool
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &m)
	if err != nil {
		utils.BuildJsonResp(c, "Error", "Json Parse Error")
		return
	}
	ret, err = services.CreateUser(&m)
	if err != nil {
		logs.Error("Create User error", err.Error())
		utils.BuildJsonResp(c, "Error", "Create User error")
		return
	}
	if ret {
		utils.BuildJsonResp(c, "Normal", "Create User Success")
	} else {
		utils.BuildJsonResp(c, "Error", "Create User Failed")
	}

	return
}

func (c *UserController) DeleteOne() {
	var m models.TaroUser
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &m)
	if err != nil {
		logs.Error("Json Parse error", err.Error())
		utils.BuildJsonResp(c, "Error", "Json Parse Error")
		return
	}

	ret, err := services.DeleteUserById(&m)
	if err != nil {
		logs.Error("DeleteUserById error", err.Error())
		utils.BuildJsonResp(c, "Error", "DeleteUserById Error")
		return
	}
	if ret {
		utils.BuildJsonResp(c, "Normal", "DeleteUserById Success")
	} else {
		utils.BuildJsonResp(c, "Error", "DeleteUserById Failed")
	}
	return
}

func (c *UserController) Update() {
	var m models.TaroUser
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &m)
	if err != nil {
		utils.BuildJsonResp(c, "Error", "Json Parse Error")
		return
	}
	ret, err := services.UpdateUser(&m)
	if err != nil {
		logs.Error("Update User error", err.Error())
		utils.BuildJsonResp(c, "Error", "Update User error")
		return
	}
	if ret {
		utils.BuildJsonResp(c, "Normal", "Update User Success")
	} else {
		utils.BuildJsonResp(c, "Error", "Update User Failed")
	}
	return
}

func (c *UserController) ListNameAndRole() {
	var req services.UserReq
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		utils.BuildJsonResp(c, "Error", "Json Parse Error")
		return
	}
	list, count, err := services.ListUserNameAndRole()
	if err != nil {
		logs.Error("List User Error", err.Error())
		utils.BuildJsonResp(c, "Error", "List User Error")
		return
	}
	c.Data["json"] = &services.UserNameAndRoleResp{
		List:  list,
		Count: count,
	}
	c.ServeJSON()
	return
}

func (c *UserController) Register() {
	var req pb.RegisterReq
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		utils.BuildJsonResp(c, "Error", "Json Parse Error")
		return
	}
	code, err := services.RegisterUser(&req)
	if err != nil {
		logs.Error("Regist Error", err.Error())
		utils.BuildJsonResp(c, "Error", "Regist Error")
		return
	}
	if code == 0 {
		logs.Info("Regist " + req.Name + " Success")
		utils.BuildJsonResp(c, "Normal", "Regist "+req.Name+" Success")
	} else {
		logs.Info("Regist " + req.Name + " Failed")
		utils.BuildJsonResp(c, "Error", "Regist "+req.Name+" Failed")
	}
	return
}

func (c *UserController) Download() {
	var req pb.DownloadReq
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		utils.BuildJsonResp(c, "Error", "Json Parse Error")
		return
	}

	resp, err := services.DownloadCert(&req)
	if err != nil {
		logs.Error("Download Cert Error", err.Error())
		utils.BuildJsonResp(c, "Error", "Download Cert Error")
		return
	}
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *UserController) VerifyIdentity() {
	var req pb.VerifyIdentityReq
	var err error
	// fmt.Println("requestBody: " + string(c.Ctx.Input.RequestBody))
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		utils.BuildJsonResp(c, "Error", "Json Parse Error")
		return
	}

	code, err := services.VerifyIdentity(&req)
	if err != nil {
		logs.Error("User Verify Error", err.Error())
		utils.BuildJsonResp(c, "Error", "User Verify Error")
		return
	}
	if code == 0 {
		logs.Info("VerifyIdentity " + req.Name + " Success")
		utils.BuildJsonResp(c, "Normal", "VerifyIdentity "+req.Name+" Success")
	} else if code != -1 {
		c.Data["json"] = pb.VerifyIdentityResp{Code: code}
		c.ServeJSON()
	} else {
		logs.Error("VerifyIdentity " + req.Name + " Failed")
		utils.BuildJsonResp(c, "Error", "VerifyIdentity "+req.Name+" Failed")
	}
	return
}

func (c *UserController) Revoke() {
	var req pb.RevokeReq
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		utils.BuildJsonResp(c, "Error", "Json Parse Error")
		return
	}

	code, err := services.RevokeUser(&req)
	if err != nil {
		logs.Error("User Revoke Error", err.Error())
		utils.BuildJsonResp(c, "Error", "User Revoke Error")
		return
	}
	if code == 0 {
		logs.Info("Revoke " + req.Name + " Success")
		utils.BuildJsonResp(c, "Normal", "Revoke "+req.Name+" Success")
	} else {
		logs.Error("Revoke " + req.Name + " Failed")
		utils.BuildJsonResp(c, "Error", "Revoke "+req.Name+" Failed")
	}
	return
}

func (c *UserController) VerifyCert() {
	var req pb.VerifyCertReq
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		utils.BuildJsonResp(c, "Error", "Json Parse Error")
		return
	}
	code, err := services.VerifyCert(&req)
	if err != nil {
		logs.Error("VerifyCert Error", err.Error())
		utils.BuildJsonResp(c, "Error", "VerifyCert Error")
		return
	}
	if code == 0 {
		logs.Info("VerifyCert " + req.Name + " Success")
		utils.BuildJsonResp(c, "Normal", "VerifyCert "+req.Name+" Success")
	} else {
		logs.Info("VerifyCert " + req.Name + " Failed")
		utils.BuildJsonResp(c, "Error", "VerifyCert "+req.Name+" Failed")
	}
	return
}

func (c *UserController) Login() {
	var req services.LoginReq
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		utils.BuildJsonResp(c, "Error", "Json Parse Error")
		return
	}
	token := c.Ctx.Input.Header("Authorization")
	ret, err := services.Login(&req, token)
	if err != nil {
		logs.Error("Login Error", err.Error())
		utils.BuildJsonResp(c, "Error", "Login "+req.UserName+" Failed")
		return
	}
	if ret != "0" && ret != "-1" {
		c.Data["json"] = &services.LoginResp{Code:0, Token:ret}
		c.ServeJSON()
		return
	}

	if ret == "0" {
		logs.Info("Login " + req.UserName + " Success")
		utils.BuildJsonResp(c, "Normal", "Login "+req.UserName+" Success")
	} else {
		logs.Info("Login " + req.UserName + " Failed")
		utils.BuildJsonResp(c, "Error", "Login "+req.UserName+" Failed")
	}
	return
}

func (c *UserController) Logout() {
	var req services.LogoutReq
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		utils.BuildJsonResp(c, "Error", "Json Parse Error")
		return
	}

	token := c.Ctx.Input.Header("Authorization")
	code, err := services.Logout(&req, token)
	if err != nil {
		logs.Error("Logout Error", err.Error())
		utils.BuildJsonResp(c, "Error", "Logout Error")
		return
	}
	if code == 0 {
		logs.Info("Logout " + req.UserName + " Success")
		utils.BuildJsonResp(c, "Normal", "Logout "+req.UserName+" Success")
	} else {
		logs.Info("Logout " + req.UserName + " Failed")
		utils.BuildJsonResp(c, "Error", "Logout "+req.UserName+" Failed")
	}
	return
}

func (c *UserController) Install() {
	var req models.TaroUser
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		utils.BuildJsonResp(c, "Error", "Json Parse Error")
		return
	}

	code, err := services.InstallUser(&req)
	if err != nil {
		logs.Error("Cert Install Error", err.Error())
		utils.BuildJsonResp(c, "Error", "Cert Install Error")
		return
	}
	if code == 0 {
		logs.Info("Install " + req.UserName + " Cert Success")
		utils.BuildJsonResp(c, "Normal", "InstallCert "+req.UserName+" Success")
	} else {
		logs.Error("Install " + req.UserName + " Cert Failed")
		utils.BuildJsonResp(c, "Error", "InstallCert "+req.UserName+" Failed")
	}
	return
}
