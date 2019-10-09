package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"tarobackend/models"
	"tarobackend/services"
	"time"
)

type ResourceController struct {
	beego.Controller
}

func (c ResourceController) URLMapping() {
	c.Mapping("Create", c.Create)
}

func (c *ResourceController) List() {
	c.Data["json"] = "kong"
	c.ServeJSON()
}

func (c *ResourceController) Create() {
	logs.Debug("enter ResourceController")
	var m models.TaroResource
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &m)
	m.ResourceCtime = time.Now()
	if err == nil {
		i, err := services.CreateResource(&m)
		if err != nil {
			logs.Debug("CreateResource error", err.Error())
		}
		c.Data["json"] = i
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}