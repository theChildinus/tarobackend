package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"tarobackend/models"
	"tarobackend/services"
	"tarobackend/utils"
)

type EnumController struct {
	beego.Controller
}

func (c *EnumController) GetValue() {
	var m models.TaroEnum
	//var ret string
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &m)
	if err != nil {
		utils.BuildJsonResp(c, "Error", "Json Parse Error")
		return
	}
	ret, err := services.GetEnumValue(&m)
	if err != nil {
		logs.Error("Get Enum Value Error", err.Error())
		utils.BuildJsonResp(c, "Error", "Get EnumValue Error")
		return
	}
	if len(ret.EnumValue) != 0 {
		c.Data["json"] = ret
		c.ServeJSON()
	} else {
		utils.BuildJsonResp(c, "Error", "Get EnumValue Error")
	}
	return
}

func (c *EnumController) PutValue() {
	var m models.TaroEnum
	var ret bool
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &m)
	if err != nil {
		utils.BuildJsonResp(c, "Error", "Json Parse Error")
		return
	}
	ret, err = services.PutEnumValue(&m)
	if err != nil {
		logs.Error("Put EnumValue error", err.Error())
		utils.BuildJsonResp(c, "Error", "Put EnumValue error")
		return
	}
	if ret {
		utils.BuildJsonResp(c, "Normal", "Put EnumValue Success")
	} else {
		utils.BuildJsonResp(c, "Error", "Put EnumValue Error")
	}
	return
}
