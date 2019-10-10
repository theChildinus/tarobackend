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
	Id int `json:"resource_id"`
}

type ResourceResp struct {
	List  []models.TaroResource `json:"list"`
	Count int64                 `json:"count"`
}

func (c ResourceController) URLMapping() {
	c.Mapping("Create", c.Create)
}

func (c *ResourceController) List() {
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
		c.Data["json"] = &utils.ErrorResp{
			ErrMsg: err.Error(),
		}
	}
	c.ServeJSON()
}

func (c *ResourceController) Create() {
	var m models.TaroResource
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &m)
	m.ResourceCtime = time.Now()
	if err == nil {
		_, err := services.CreateResource(&m)
		if err != nil {
			logs.Error("CreateResource error", err.Error())
		}
		c.Data["json"] = &utils.NormalResp{
			NorMsg: "Crete Resource Success",
		}
	} else {
		c.Data["json"] = &utils.ErrorResp{
			ErrMsg: err.Error(),
		}
	}
	c.ServeJSON()
}

func (c *ResourceController) DeleteOne() {
	var req ResourceReq
	var err error
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		c.Data["json"] = &utils.ErrorResp{
			ErrMsg: "Json Parse Error",
		}
		c.ServeJSON()
		return
	}

	err = services.DeleteResourceById(req.Id)
	if err != nil {
		logs.Error("DeleteResourceById error", err.Error())
		c.Data["json"] = &utils.ErrorResp{
			ErrMsg: "DeleteResourceById error",
		}
		c.ServeJSON()
		return
	}
	c.Data["json"] = &utils.NormalResp{
		NorMsg: "DeleteResourceById success",
	}
	c.ServeJSON()
	return
}
