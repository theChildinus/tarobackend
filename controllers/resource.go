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

type ResourceReq struct {
}

type ResourceResp struct {
	List  []models.TaroResource `json:"list"`
	Count int64                 `json:"count"`
}

func (c ResourceController) URLMapping() {
	c.Mapping("Create", c.Create)
}

func (c *ResourceController) List() {
	logs.Debug("enter ListResource")
	var page utils.Page
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &page)
	if err == nil {
		list, count, err := services.ListResource(page.PageIndex, page.PageSize)
		if err != nil {
			logs.Error("ListResource error", err.Error())
		}
		c.Data["json"] = &ResourceResp{
			List:  list,
			Count: count,
		}
	} else {
		c.Data["json"] = &utils.Error{
			ErrMsg: err.Error(),
		}
	}
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
			logs.Error("CreateResource error", err.Error())
		}
		c.Data["json"] = i
	} else {
		c.Data["json"] = &utils.Error{
			ErrMsg: err.Error(),
		}
	}
	c.ServeJSON()
}
