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
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &m)
	if err != nil {
		utils.BuildJsonResp(c, "Error", "Json Parse Error")
		return
	}
	_, err = services.CreateUser(&m)
	if err != nil {
		logs.Error("Create User error", err.Error())
		utils.BuildJsonResp(c, "Error", "Create User error")
		return
	}
	utils.BuildJsonResp(c, "Normal", "Create User Success")
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

	err = services.DeleteUserById(m.UserId)
	if err != nil {
		logs.Error("DeleteUserById error", err.Error())
		utils.BuildJsonResp(c, "Error", "DeleteUserById Error")
		return
	}
	utils.BuildJsonResp(c, "Normal", "DeleteUserById Success")
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
	err = services.UpdateUser(&m)
	if err != nil {
		logs.Error("Update User error", err.Error())
		utils.BuildJsonResp(c, "Error", "Update User error")
		return
	}
	utils.BuildJsonResp(c, "Normal", "Update User Success")
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
		logs.Info("Regist " + req.Username + " Success")
		utils.BuildJsonResp(c, "Normal", "Regist "+req.Username+" Success")
	} else {
		logs.Info("Regist " + req.Username + " Failed")
		utils.BuildJsonResp(c, "Error", "Regist "+req.Username+" Failed")
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

	resp, err := services.DownloadCard(&req)
	if err != nil {
		logs.Error("Download Card Error", err.Error())
		utils.BuildJsonResp(c, "Error", "Download Card Error")
		return
	}
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *UserController) Login() {
	var req pb.LoginReq
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		utils.BuildJsonResp(c, "Error", "Json Parse Error")
		return
	}

	code, err := services.Login(&req)
	if err != nil {
		logs.Error("User Login Error", err.Error())
		utils.BuildJsonResp(c, "Error", "User Login Error")
		return
	}
	if code == 0 {
		logs.Info("Login " + req.Username + " Success")
		utils.BuildJsonResp(c, "Normal", "Login "+req.Username+" Success")
	} else {
		logs.Info("Login " + req.Username + " Failed")
		utils.BuildJsonResp(c, "Error", "Login "+req.Username+" Failed")
	}
	return
}
