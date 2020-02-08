package services

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/casbin/casbin"
	"strings"
	"tarobackend/models"
	"tarobackend/utils"
)

type PolicyReq struct {
	PageIndex int64  `json:"page_index"`
	PageSize  int64  `json:"page_size"`
	SearchSub string `json:"search_sub"`
}

type PolicyCheckReq struct {
	PolicyName string `json:"policyname"`
	PolicySub string `json:"policysub"`
	PolicyObj string `json:"policyobj"`
	PolicyAct string `json:"policyact"`
	UserName  string `json:"username"`
	UserHash  string `json:"userhash"`
}

type PolicyResp struct {
	List  []models.TaroPolicy `json:"list"`
	Count int64               `json:"count"`
}

type RoleAllotReq struct {
	Name string `json:"name"`
	Role []string `json:"role"`
}

type PolicyModel struct {
	PolicyName string `json:"policy_name"`
	ModelType string `json:"model_type"`
}

type ExecutableReq struct {
	EpcCtx map[string]string `json:"epcCtx"`
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
		logs.Error("ListPolicy: Table Policy List Failed")
		return nil, 0, err
	}
	return policies, count, nil
}

func CreatePolicy(r *models.TaroPolicy) (bool, error) {
	engine := utils.Engine_mysql
	_, err := engine.InsertOne(r)
	if err != nil {
		logs.Error("CreatePolicy: Table Policy InsertOne Error")
		return false, err
	}
	casbin_model := "./casbinfiles/rbac_model.conf"
	casbin_policys := "./casbinfiles/" + r.PolicyName + ".csv"
	if ok, err := utils.FileExistAndCreate(casbin_policys); !ok {
		return false, err
	}
	enf := casbin.NewEnforcer(casbin_model, casbin_policys)
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
		casbin_model := "./casbinfiles/rbac_model.conf"
		casbin_policys := "./casbinfiles/" + r.PolicyName + ".csv"
		if ok, err := utils.FileExistAndCreate(casbin_policys); !ok {
			return false, err
		}
		enf := casbin.NewEnforcer(casbin_model, casbin_policys)
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
		casbin_model := "./casbinfiles/rbac_model.conf"
		casbin_policys_old := "./casbinfiles/" + old.PolicyName + ".csv"
		casbin_policys_new := "./casbinfiles/" + r.PolicyName + ".csv"
		if ok, err := utils.FileExistAndCreate(casbin_policys_new); !ok {
			return false, err
		}
		if old.PolicyName == r.PolicyName {
			enf1 := casbin.NewEnforcer(casbin_model, casbin_policys_new)
			ret1 := enf1.RemovePolicy(old.PolicySub, old.PolicyObj, old.PolicyAct)
			ret2 := enf1.AddPolicy(r.PolicySub, r.PolicyObj, r.PolicyAct)
			_ = enf1.SavePolicy()
			ret = ret1 && ret2
		} else {
			enf1 := casbin.NewEnforcer(casbin_model, casbin_policys_old)
			enf2 := casbin.NewEnforcer(casbin_model, casbin_policys_new)
			ret1 := enf1.RemovePolicy(old.PolicySub, old.PolicyObj, old.PolicyAct)
			ret2 := enf2.AddPolicy(r.PolicySub, r.PolicyObj, r.PolicyAct)
			_ = enf1.SavePolicy()
			_ = enf2.SavePolicy()
			ret = ret1 && ret2
		}
	}
	_, err = engine.ID(r.PolicyId).Update(r)
	if err != nil {
		logs.Error("UpdatePolicy: Table Policy Update Error")
		return false, err
	}
	return ret, nil
}

func CheckPolicy(r *PolicyCheckReq) (bool, error) {
	if len(r.UserHash) == 0 {
		logs.Error("CheckPolicy: UserHase Empty")
		return false, errors.New("CheckPolicy: UserHase Empty")
	}
	engine := utils.Engine_mysql
	m := new(models.TaroUser)
	has, err := engine.Table("taro_user").
		Where("user_name = ?", r.UserName).
		And("user_hash = ?", r.UserHash).Get(m)
	if err != nil {
		logs.Error("CheckPolicy: Table User Get Error")
		return false, err
	}
	// fmt.Println("r.policysub:", r.PolicySub, "m.username:", m.UserName, "m.userrole:", m.UserRole)
	if has &&
		r.UserHash == m.UserHash &&
		(r.PolicySub == m.UserName || r.PolicySub == m.UserRole) {
		req := &models.TaroEnum{EnumKey:"policy_model"}
		enum, err := GetEnumValue(req)
		if err != nil {
			return false, err
		}
		var pms []PolicyModel
		if err := json.Unmarshal([]byte(enum.EnumValue), &pms); err != nil {
			return false, err
		}
		model_type := "acl"
		if strings.Index(r.PolicyName, "#") != -1 {
			firstStr := strings.Split(r.PolicyName, "#")[0]
			for _, v := range pms {
				if v.PolicyName == firstStr {
					model_type = v.ModelType
					break
				}
			}
		}
		casbin_model := "./casbinfiles/" + strings.ToLower(model_type) + "_model.conf"
		casbin_policys := "./casbinfiles/" + r.PolicyName + ".csv"
		logs.Info("[CheckPolicy] modelfile:", casbin_model, "policyfile:", casbin_policys)
		enf := casbin.NewEnforcer(casbin_model, casbin_policys)
		ret := enf.Enforce(r.PolicySub, r.PolicyObj, r.PolicyAct)
		return ret, nil
	} else {
		return false, nil
	}
}
