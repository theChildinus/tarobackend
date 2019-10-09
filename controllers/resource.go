package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
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
	c.Data["json"] = 1
	c.ServeJSON()
}