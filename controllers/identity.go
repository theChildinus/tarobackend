package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"tarobackend/models"
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
