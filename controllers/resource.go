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

type ResourceController struct {
	beego.Controller
}

func (c *ResourceController) List() {
	var req services.ResourceReq
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		utils.BuildJsonResp(c, "Error", "Json Parse Error")
		return
	}
	list, count, err := services.ListResource(&req)
	if err != nil {
		logs.Error("List Resource Error", err.Error())
		utils.BuildJsonResp(c, "Error", "List Resource Error")
		return
	}
	c.Data["json"] = &services.ResourceResp{
		List:  list,
		Count: count,
	}
	c.ServeJSON()
	return
}

func (c *ResourceController) Create() {
	var m models.TaroResource
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &m)
	if err != nil {
		utils.BuildJsonResp(c, "Error", "Json Parse Error")
		return
	}
	m.ResourceCtime = time.Now()
	_, err = services.CreateResource(&m)
	if err != nil {
		logs.Error("Create Resource error", err.Error())
		utils.BuildJsonResp(c, "Error", "Create Resource error")
		return
	}
	utils.BuildJsonResp(c, "Normal", "Create Resource Success")
	return
}

func (c *ResourceController) DeleteOne() {
	var m models.TaroResource
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &m)
	if err != nil {
		logs.Error("Json Parse error", err.Error())
		utils.BuildJsonResp(c, "Error", "Json Parse Error")
		return
	}

	err = services.DeleteResourceById(m.ResourceId)
	if err != nil {
		logs.Error("DeleteResourceById error", err.Error())
		utils.BuildJsonResp(c, "Error", "DeleteResourceById Error")
		return
	}
	utils.BuildJsonResp(c, "Normal", "DeleteResourceById Success")
	return
}

func (c *ResourceController) Update() {
	var m models.TaroResource
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &m)
	if err != nil {
		utils.BuildJsonResp(c, "Error", "Json Parse Error")
		return
	}
	err = services.UpdateResource(&m)
	if err != nil {
		logs.Error("Update Resource error", err.Error())
		utils.BuildJsonResp(c, "Error", "Update Resource error")
		return
	}
	utils.BuildJsonResp(c, "Normal", "Update Resource Success")
	return
}
