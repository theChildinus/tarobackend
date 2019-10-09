package services

import (
	"github.com/astaxie/beego/logs"
	"tarobackend/models"
	"tarobackend/utils"
)

func CreateResource(r *models.TaroResource) (int64, error) {
	engine := utils.Engine_mysql
	logs.Debug(r)
	res, err := engine.InsertOne(r)
	if err != nil {
		return 0, err
	}
	if res == 0 {
		logs.Debug("Resource InsertOne failed")
	}
	return res, nil
}
