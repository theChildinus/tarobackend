package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"tarobackend/models"
	pb "tarobackend/proto"
	"tarobackend/services"
	"tarobackend/utils"
	"time"
)

type IdentityController struct {
	beego.Controller
}

func (c *IdentityController) List() {
	var req services.IdentityReq
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		utils.BuildJsonResp(c, "Error", "Json Parse Error")
		return
	}
	list, count, err := services.ListIdentity(&req)
	if err != nil {
		logs.Error("List Identity Error", err.Error())
		utils.BuildJsonResp(c, "Error", "List Identity Error")
		return
	}
	c.Data["json"] = &services.IdentityResp{
		List:  list,
		Count: count,
	}
	c.ServeJSON()
	return
}

func (c *IdentityController) Create() {
	var m models.TaroIdentity
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &m)
	if err != nil {
		utils.BuildJsonResp(c, "Error", "Json Parse Error")
		return
	}
	m.IdentityCtime = time.Now()
	_, err = services.CreateIdentity(&m)
	if err != nil {
		logs.Error("Create Identity error", err.Error())
		utils.BuildJsonResp(c, "Error", "Create Identity error")
		return
	}
	utils.BuildJsonResp(c, "Normal", "Create Identity Success")
	return
}

func (c *IdentityController) DeleteOne() {
	var m models.TaroIdentity
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &m)
	if err != nil {
		logs.Error("Json Parse error", err.Error())
		utils.BuildJsonResp(c, "Error", "Json Parse Error")
		return
	}

	err = services.DeleteIdentityById(m.IdentityId)
	if err != nil {
		logs.Error("DeleteIdentityById error", err.Error())
		utils.BuildJsonResp(c, "Error", "DeleteIdentityById Error")
		return
	}
	utils.BuildJsonResp(c, "Normal", "DeleteIdentityById Success")
	return
}

func (c *IdentityController) Update() {
	var m models.TaroIdentity
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &m)
	if err != nil {
		utils.BuildJsonResp(c, "Error", "Json Parse Error")
		return
	}
	err = services.UpdateIdentity(&m)
	if err != nil {
		logs.Error("Update Identity error", err.Error())
		utils.BuildJsonResp(c, "Error", "Update Identity error")
		return
	}
	utils.BuildJsonResp(c, "Normal", "Update Identity Success")
	return
}

func (c *IdentityController) Register() {
	var req pb.RegisterReq
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		utils.BuildJsonResp(c, "Error", "Json Parse Error")
		return
	}
	code, err := services.RegisterIdentity(&req)
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

func (c *IdentityController) Enroll() {
	var req pb.EnrollReq
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		utils.BuildJsonResp(c, "Error", "Json Parse Error")
		return
	}
	code, err := services.EnrollIdentity(&req)
	if err != nil {
		logs.Error("Enroll Error", err.Error())
		utils.BuildJsonResp(c, "Error", "Enroll Error")
		return
	}
	if code == 0 {
		logs.Info("Enroll " + req.Name + " Success")
		utils.BuildJsonResp(c, "Normal", "Enroll "+req.Name+" Success")
	} else {
		logs.Info("Enroll " + req.Name + " Failed")
		utils.BuildJsonResp(c, "Error", "Enroll "+req.Name+" Failed")
	}
	return
}

func (c *IdentityController) Revoke() {
	var req pb.RevokeReq
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		utils.BuildJsonResp(c, "Error", "Json Parse Error")
		return
	}

	code, err := services.RevokeIdentity(&req)
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
