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
	r := new(models.TaroResource)
	count, err := engine.Count(r)
	if err != nil {
		logs.Debug("Resource Count failed")
		return nil, 0, err
	}
	return resources, count, nil
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

func DeleteResourceById(id int) error {
	engine := utils.Engine_mysql
	r := new(models.TaroResource)
	_, err := engine.ID(id).Delete(r)
	if err != nil {
		return err
	}
	return nil
}