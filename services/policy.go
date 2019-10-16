package services

import (
	"github.com/astaxie/beego/logs"
	"tarobackend/models"
	"tarobackend/utils"
)

type PolicyReq struct {
	PageIndex int64 `json:"page_index"`
	PageSize  int64 `json:"page_size"`
	SearchSub string `json:"search_sub"`
}

type PolicyResp struct {
	List  []models.TaroPolicy `json:"list"`
	Count int64               `json:"count"`
}

func ListPolicy(req *PolicyReq) ([]models.TaroPolicy, int64, error) {
	engine := utils.Engine_mysql
	var (
		policies []models.TaroPolicy
		err      error
		count    int64
	)
	m := new(models.TaroPolicy)
	if len(req.SearchSub) != 0 {
		err = engine.Table("taro_policy").
			Where("policy_sub = ? ", req.SearchSub).
			Limit(int(req.PageSize), int((req.PageIndex-1)*req.PageSize)).
			Find(&policies)
		count, _ = engine.Where("policy_sub = ? ", req.SearchSub).Count(m)
	} else {
		err = engine.Table("taro_policy").
			Limit(int(req.PageSize), int((req.PageIndex-1)*req.PageSize)).
			Find(&policies)
		count, _ = engine.Count(m)
	}

	if err != nil {
		logs.Error("Table Policy Find Failed")
		return nil, 0, err
	}
	return policies, count, nil
}

func CreatePolicy(r *models.TaroPolicy) (bool, error) {
	engine := utils.Engine_mysql
	_, err := engine.InsertOne(r)
	if err != nil {
		logs.Error("Table Policy InsertOne Error")
		return false, err
	}

	enf := utils.Enforcer
	ret := enf.AddPolicy(r.PolicySub, r.PolicyObj, r.PolicyAct)
	_ = enf.SavePolicy()
	return ret, nil
}

func DeletePolicyById(id int) (bool, error) {
	engine := utils.Engine_mysql
	r := new(models.TaroPolicy)
	var ret bool

	has, err := engine.Table("taro_policy").
		Where("policy_id = ?", id).Get(r)
	if err != nil {
		logs.Error("DeletePolicyById: Table Policy Get Error")
		return false, err
	}
	if has {
		enf := utils.Enforcer
		ret = enf.RemovePolicy(r.PolicySub, r.PolicyObj, r.PolicyAct)
		_ = enf.SavePolicy()
	}

	_, err = engine.ID(id).Delete(r)
	if err != nil {
		logs.Error("DeletePolicyById: Table Policy Delete Error")
		return false, err
	}
	return ret, nil
}

func UpdatePolicy(r *models.TaroPolicy) (bool, error) {
	engine := utils.Engine_mysql
	old := new(models.TaroPolicy)
	var ret bool
	has, err := engine.Table("taro_policy").
		Where("policy_id = ?", r.PolicyId).Get(old)
	if err != nil {
		logs.Error("UpdatePolicy: Table Policy Get Error")
		return false, err
	}
	if has {
		enf := utils.Enforcer
		ret1 := enf.RemovePolicy(old.PolicySub, old.PolicyObj, old.PolicyAct)
		ret2 := enf.AddPolicy(r.PolicySub, r.PolicyObj, r.PolicyAct)
		_ = enf.SavePolicy()
		ret = ret1 && ret2
	}
	_, err = engine.ID(r.PolicyId).Update(r)
	if err != nil {
		logs.Error("UpdatePolicy: Table Policy Update Error")
		return false, err
	}
	return ret, nil
}

func CheckPolicy(r *models.TaroPolicy) (bool, error) {
	enf := utils.Enforcer
	ret := enf.Enforce(r.PolicySub, r.PolicyObj, r.PolicyAct)
	return ret, nil
}
