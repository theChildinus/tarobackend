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

type PolicyController struct {
	beego.Controller
}

func (c *PolicyController) Create() {
	var m models.TaroPolicy
	var ret bool
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &m)
	if err != nil {
		utils.BuildJsonResp(c, "Error", "Json Parse Error")
		return
	}
	m.PolicyCtime = time.Now()
	ret, err = services.CreatePolicy(&m)
	if err != nil {
		logs.Error("Create Policy error", err.Error())
		utils.BuildJsonResp(c, "Error", "Create Policy error")
		return
	}
	if ret {
		utils.BuildJsonResp(c, "Normal", "Create Policy Success")
	} else {
		utils.BuildJsonResp(c, "Normal", "Policy Existed")
	}
	return
}

func (c *PolicyController) List() {
	var req services.PolicyReq
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		utils.BuildJsonResp(c, "Error", "Json Parse Error")
		return
	}
	list, count, err := services.ListPolicy(&req)
	if err != nil {
		logs.Error("List Policy Error", err.Error())
		utils.BuildJsonResp(c, "Error", "List Policy Error")
		return
	}
	c.Data["json"] = &services.PolicyResp{
		List:  list,
		Count: count,
	}
	c.ServeJSON()
	return
}

func (c *PolicyController) DeleteOne() {
	var m models.TaroPolicy
	var ret bool
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &m)
	if err != nil {
		logs.Error("Json Parse error", err.Error())
		utils.BuildJsonResp(c, "Error", "Json Parse Error")
		return
	}

	ret, err = services.DeletePolicyById(m.PolicyId)
	if err != nil {
		logs.Error("DeletePolicyById error", err.Error())
		utils.BuildJsonResp(c, "Error", "DeletePolicyById Error")
		return
	}
	if ret {
		utils.BuildJsonResp(c, "Normal", "Delete Policy Success")
	} else {
		utils.BuildJsonResp(c, "Normal", "Policy Not Exist")
	}

	return
}

func (c *PolicyController) Update() {
	var m models.TaroPolicy
	var ret bool
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &m)
	if err != nil {
		utils.BuildJsonResp(c, "Error", "Json Parse Error")
		return
	}
	ret, err = services.UpdatePolicy(&m)
	if err != nil {
		logs.Error("Update Policy error", err.Error())
		utils.BuildJsonResp(c, "Error", "Update Policy error")
		return
	}
	if ret {
		utils.BuildJsonResp(c, "Normal", "Update Policy Success")
	} else {
		utils.BuildJsonResp(c, "Normal", "Policy Not Exist")
	}

	return
}

func (c *PolicyController) Check() {
	var m services.PolicyCheckReq
	var ret bool
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &m)
	if err != nil {
		utils.BuildJsonResp(c, "Error", "Json Parse Error")
		return
	}
	ret, err = services.CheckPolicy(&m)
	if err != nil {
		logs.Error("Check Policy error", err.Error())
		utils.BuildJsonResp(c, "Error", "Check Policy Error")
		return
	}
	if ret {
		utils.BuildJsonResp(c, "Normal", "Policy Check Success")
	} else {
		utils.BuildJsonResp(c, "Error", "Policy Check Failed")
	}

	return
}
