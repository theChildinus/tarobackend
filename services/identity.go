package services

import (
"github.com/astaxie/beego/logs"
"tarobackend/models"
"tarobackend/utils"
)

type IdentityReq struct {
	PageIndex  int64  `json:"page_index"`
	PageSize   int64  `json:"page_size"`
	SearchType string `json:"search_type"`
	SearchName string `json:"search_name"`
}

type IdentityResp struct {
	List  []models.TaroIdentity `json:"list"`
	Count int64                 `json:"count"`
}

func ListIdentity(req *IdentityReq) ([]models.TaroIdentity, int64, error) {
	engine := utils.Engine_mysql
	var (
		Identitys []models.TaroIdentity
		err       error
		count     int64
	)
	m := new(models.TaroIdentity)
	if len(req.SearchType) != 0 {
		err = engine.Table("taro_identity").
			Where("identity_name like ? ", "%"+req.SearchName+"%").
			And("identity_type = ?", req.SearchType).
			Limit(int(req.PageSize), int((req.PageIndex-1)*req.PageSize)).
			Find(&Identitys)
		count, _ = engine.Where("identity_name like ? ", "%"+req.SearchName+"%").
			And("identity_type = ?", req.SearchType).Count(m)
	} else {
		err = engine.Table("taro_identity").
			Limit(int(req.PageSize), int((req.PageIndex-1)*req.PageSize)).
			Find(&Identitys)
		count, _ = engine.Count(m)
	}

	if err != nil {
		logs.Error("ListIdentity: Table Identity Find Error")
		return nil, 0, err
	}
	return Identitys, count, nil
}

func CreateIdentity(r *models.TaroIdentity) (int64, error) {
	engine := utils.Engine_mysql
	res, err := engine.InsertOne(r)
	if err != nil {
		logs.Error("CreateIdentity: Table Identity InsertOne Error")
		return 0, err
	}
	if res == 0 {
		logs.Debug("CreateIdentity: Identity InsertOne failed")
	}
	return res, nil
}

func DeleteIdentityById(id int) error {
	engine := utils.Engine_mysql
	r := new(models.TaroIdentity)
	_, err := engine.ID(id).Delete(r)
	if err != nil {
		logs.Error("DeleteIdentityById: Table Identity Delete Error")
		return err
	}
	return nil
}

func UpdateIdentity(r *models.TaroIdentity) error {
	engine := utils.Engine_mysql
	_, err := engine.ID(r.IdentityId).Update(r)
	if err != nil {
		logs.Error("UpdateIdentity: Table Identity Update Error")
		return err
	}
	return nil
}
