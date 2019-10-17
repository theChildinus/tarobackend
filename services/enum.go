package services

import (
	"github.com/astaxie/beego/logs"
	"tarobackend/models"
	"tarobackend/utils"
)

func GetEnumValue(r *models.TaroEnum) (*models.TaroEnum, error) {
	engine := utils.Engine_mysql
	enum := new(models.TaroEnum)
	_, err := engine.Table("taro_enum").
		Where("enum_key = ?", r.EnumKey).Get(enum)
	if err != nil {
		logs.Error("GetEnumValue: Table Enum Get Error")
		return nil, err
	}
	return enum, nil
}

func PutEnumValue(r *models.TaroEnum) (bool, error) {
	engine := utils.Engine_mysql
	_, err := engine.Id(r.EnumId).Update(r)
	if err != nil {
		logs.Error("PutEnumValue: Table Enum Update Error")
		return false, err
	}
	return true, nil
}
