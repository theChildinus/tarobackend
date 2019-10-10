package services

import (
	"github.com/astaxie/beego/logs"
	"tarobackend/models"
	"tarobackend/utils"
)

func ListResource(index, offset int64) ([]models.TaroResource, int64, error) {
	engine := utils.Engine_mysql
	var resources []models.TaroResource
	err := engine.Table("taro_resource").
		Limit(int(offset), int((index - 1) * offset)).
		Find(&resources)
	if err != nil {
		logs.Debug("Resource Find failed")
		return nil, 0, err
	}
	return resources, int64(len(resources)), nil
}

func CreateResource(r *models.TaroResource) (int64, error) {
	engine := utils.Engine_mysql
	res, err := engine.InsertOne(r)
	if err != nil {
		return 0, err
	}
	if res == 0 {
		logs.Debug("Resource InsertOne failed")
	}
	return res, nil
}
