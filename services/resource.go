package services

import (
	"github.com/astaxie/beego/logs"
	"tarobackend/models"
	"tarobackend/utils"
)

type ResourceReq struct {
	PageIndex  int64  `json:"page_index"`
	PageSize   int64  `json:"page_size"`
	SearchType int64  `json:"search_type"`
	SearchName string `json:"search_name"`
}

type ResourceResp struct {
	List  []models.TaroResource `json:"list"`
	Count int64                 `json:"count"`
}

func ListResource(req *ResourceReq) ([]models.TaroResource, int64, error) {
	engine := utils.Engine_mysql
	var (
		resources []models.TaroResource
		err       error
		count     int64
	)
	m := new(models.TaroResource)
	if req.SearchType != -1 {
		err = engine.Table("taro_resource").
			Where("resource_name like ? ", "%"+req.SearchName+"%").
			And("resource_type = ?", req.SearchType).
			Limit(int(req.PageSize), int((req.PageIndex-1)*req.PageSize)).
			Find(&resources)
		count, _ = engine.Where("resource_name like ? ", "%"+req.SearchName+"%").
			And("resource_type = ?", req.SearchType).Count(m)
	} else {
		err = engine.Table("taro_resource").
			Limit(int(req.PageSize), int((req.PageIndex-1)*req.PageSize)).
			Find(&resources)
		count, _ = engine.Count(m)
	}

	if err != nil {
		logs.Error("ListResource: Table Resource Find Error")
		return nil, 0, err
	}
	return resources, count, nil
}

func CreateResource(r *models.TaroResource) (int64, error) {
	engine := utils.Engine_mysql
	res, err := engine.InsertOne(r)
	if err != nil {
		logs.Error("CreateResource: Table Resource InsertOne Error")
		return 0, err
	}
	if res == 0 {
		logs.Debug("Resource InsertOne failed")
	}
	return res, nil
}

func DeleteResourceById(id int) error {
	engine := utils.Engine_mysql
	r := new(models.TaroResource)
	_, err := engine.ID(id).Delete(r)
	if err != nil {
		logs.Error("DeleteResourceById: Table Resource Delete Error")
		return err
	}
	return nil
}

func UpdateResource(r *models.TaroResource) error {
	engine := utils.Engine_mysql
	_, err := engine.ID(r.ResourceId).Update(r)
	if err != nil {
		logs.Error("UpdateResource: Table Resource Update Error")
		return err
	}
	return nil
}
