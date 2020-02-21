package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
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
	Roles []string `json:"roles"`
}

type PolicyModel struct {
	PolicyName string `json:"policy_name"`
	ModelType string `json:"model_type"`
}

type MutexRole struct {
	Role1 string `json:"user_role1"`
	Role2 string `json:"user_role2"`
}

type ExecutableReq struct {
	EpcCtx utils.EpcCtx `json:"epcCtx"`
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
			Where("policy_sub like ? ", "%"+req.SearchSub+"%").
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
	model_type, err := GetModelType(r.PolicyName)
	if err != nil {
		logs.Error("CreatePolicy: Get ModelType Error")
		return false, err
	}
	casbin_model := "./casbinfiles/" + strings.ToLower(model_type) + "_model.conf"
	casbin_policys := "./casbinfiles/" + r.PolicyName + ".csv"
	if ok, err := utils.FileExistAndCreate(casbin_policys); !ok {
		return false, err
	}
	enf := casbin.NewEnforcer(casbin_model, casbin_policys)
	// if policy exist, return false, else add and insert to db
	if has := enf.HasPolicy(r.PolicySub, r.PolicyObj, r.PolicyAct); has {
		return false, nil
	}
	ret := enf.AddPolicy(r.PolicySub, r.PolicyObj, r.PolicyAct)
	engine := utils.Engine_mysql
	_, err = engine.InsertOne(r)
	if err != nil {
		logs.Error("CreatePolicy: Table Policy InsertOne Error")
		return false, err
	}
	if model_type == "RBAC" {
		var users []models.TaroUser
		err = engine.Table("taro_user").
			Where("user_role like ? ", "%"+r.PolicySub+"%").Find(&users)
		fmt.Println("users:", r.PolicySub, users)
		for _, v := range users {
			_ = enf.AddRoleForUser(v.UserName, r.PolicySub)
		}
	}
	_ = enf.SavePolicy()

	// add policy to fabric configtx.yaml file when policyName matching
	if r.PolicyName == beego.AppConfig.String("fabric_policy_name") {
		tx, err := utils.ParseYamlFile(beego.AppConfig.String("fabric_configtx"))
		if err == nil {
			if len(tx.Application.ACLs) == 0{
				tx.Application.ACLs = make(map[string]string)
			}
			subStr := r.PolicyObj[strings.Index(r.PolicyObj, "/") + 1:]
			tx.Application.ACLs[subStr] = r.PolicySub
			if err = utils.SaveYamlFile(tx, beego.AppConfig.String("fabric_configtx")); err != nil {
				return false, err
			}
		} else {
			return false, err
		}
	}
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
		model_type, err := GetModelType(r.PolicyName)
		if err != nil {
			logs.Error("CreatePolicy: Get ModelType Error")
			return false, err
		}
		casbin_model := "./casbinfiles/" + strings.ToLower(model_type) + "_model.conf"
		casbin_policys := "./casbinfiles/" + r.PolicyName + ".csv"
		if ok, err := utils.FileExistAndCreate(casbin_policys); !ok {
			return false, err
		}
		enf := casbin.NewEnforcer(casbin_model, casbin_policys)
		ret = enf.RemovePolicy(r.PolicySub, r.PolicyObj, r.PolicyAct)
		if model_type == "RBAC" {
			var users []models.TaroUser
			err = engine.Table("taro_user").
				Where("user_role like ? ", "%"+r.PolicySub+"%").Find(&users)
			fmt.Println("users:", r.PolicySub, users)
			for _, v := range users {
				_ = enf.DeleteRoleForUser(v.UserName, r.PolicySub)
			}
		}
		_ = enf.SavePolicy()
	}

	_, err = engine.ID(id).Delete(r)
	if err != nil {
		logs.Error("DeletePolicyById: Table Policy Delete Error")
		return false, err
	}

	if r.PolicyName == beego.AppConfig.String("fabric_policy_name") {
		tx, err := utils.ParseYamlFile(beego.AppConfig.String("fabric_configtx"))
		if err == nil && len(tx.Application.ACLs) != 0 {
			subStr := r.PolicyObj[strings.Index(r.PolicyObj, "/") + 1:]
			delete(tx.Application.ACLs, subStr)
			err = utils.SaveYamlFile(tx, beego.AppConfig.String("fabric_configtx"))
			if err != nil {
				return false, err
			}
		} else {
			return false, err
		}
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
		model_type, err := GetModelType(r.PolicyName)
		if err != nil {
			logs.Error("CreatePolicy: Get ModelType Error")
			return false, err
		}
		casbin_model := "./casbinfiles/" + strings.ToLower(model_type) + "_model.conf"
		casbin_policys := "./casbinfiles/" + r.PolicyName + ".csv"
		if ok, err := utils.FileExistAndCreate(casbin_policys); !ok {
			return false, err
		}
		enf := casbin.NewEnforcer(casbin_model, casbin_policys)
		ret1 := enf.RemovePolicy(old.PolicySub, old.PolicyObj, old.PolicyAct)
		ret2 := enf.AddPolicy(r.PolicySub, r.PolicyObj, r.PolicyAct)
		if model_type == "RBAC" {
			// remove policy rule: users has old role
			var users []models.TaroUser
			err = engine.Table("taro_user").
				Where("user_role like ? ", "%"+old.PolicySub+"%").Find(&users)
			for _, v := range users {
				_ = enf.DeleteRoleForUser(v.UserName, old.PolicySub)
			}
			// add policy rule: users has new role
			users = users[0:0]
			err = engine.Table("taro_user").
				Where("user_role like ? ", "%"+r.PolicySub+"%").Find(&users)
			for _, v := range users {
				_ = enf.AddRoleForUser(v.UserName, r.PolicySub)
			}
		}
		_ = enf.SavePolicy()
		ret = ret1 && ret2
		_ = enf.SavePolicy()
	}
	_, err = engine.ID(r.PolicyId).Update(r)
	if err != nil {
		logs.Error("UpdatePolicy: Table Policy Update Error")
		return false, err
	}

	if r.PolicyName == beego.AppConfig.String("fabric_policy_name") {
		tx, err := utils.ParseYamlFile(beego.AppConfig.String("fabric_configtx"))
		if err == nil && len(tx.Application.ACLs) != 0 {
			subStr := r.PolicyObj[strings.Index(r.PolicyObj, "/") + 1:]
			tx.Application.ACLs[subStr] = r.PolicySub
			err = utils.SaveYamlFile(tx, beego.AppConfig.String("fabric_configtx"))
			if err != nil {
				return false, err
			}
		} else {
			return false, err
		}
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
		model_type, err := GetModelType(r.PolicyName)
		if err != nil {
			logs.Error("CheckPolicy: Get Model Type Error")
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

func RoleAllot(r *RoleAllotReq) (bool, error) {
	if len(r.Roles) < 2 {
		return true, nil
	}
	m := &models.TaroEnum{EnumKey:"mutex_role"}
	enum, err := GetEnumValue(m)
	if err != nil {
		return false, err
	}
	var mrs []MutexRole
	if err := json.Unmarshal([]byte(enum.EnumValue), &mrs); err != nil {
		return false, err
	}
	for _, v := range mrs {
		if len(r.Roles) == 2 &&
			((r.Roles[0] == v.Role1 && r.Roles[1] == v.Role2) ||
				(r.Roles[0] == v.Role2 && r.Roles[1] == v.Role1)) {
			return false, nil
		}
	}
	return true, nil
}

func GetModelType(pn string) (string, error) {
	model_type := "acl"
	req := &models.TaroEnum{EnumKey:"policy_model"}
	enum, err := GetEnumValue(req)
	if err != nil {
		return "", err
	}
	var pms []PolicyModel
	if err := json.Unmarshal([]byte(enum.EnumValue), &pms); err != nil {
		return "", err
	}
	firstStr := ""
	if strings.Index(pn, "#") != -1 {
		firstStr = strings.Split(pn, "#")[0]
	} else {
		firstStr = pn
	}
	for _, v := range pms {
		if v.PolicyName == firstStr {
			model_type = v.ModelType
			break
		}
	}
	return model_type, nil
}

func Executable(r *ExecutableReq) (string, error) {
	type OuFuncIu struct {
		Ou string   // organization unit pel
		Func string // function pel
		Iu string   // information unit pel
	}
	var ou_func_iu []*OuFuncIu
	function := make(map[string]interface{})
	ou := make(map[string]interface{})
	iu := make(map[string]interface{})
	for _, v := range r.EpcCtx.Epc.Function {
		function[v.ID] = v.Name
	}
	for _, v := range r.EpcCtx.Epc.Ou {
		ou[v.ID] = v.OuName
	}
	for _, v := range r.EpcCtx.Epc.Iu {
		iu[v.ID] = v.IuName
	}
	arc := r.EpcCtx.Epc.Arc
	for i := 0; i < len(arc) - 1; i++ {
		for j := 1; j < len(arc); j++ {
			is, it, js, jt := arc[i].Flow.Source, arc[i].Flow.Target,
				arc[j].Flow.Source, arc[j].Flow.Target
			if _, ok := function[it]; ok && it == jt {
				_, ok1 := ou[is]
				_, ok2 := iu[js]
				_, ok3 := ou[js]
				_, ok4 := iu[is]
				if ok1 && ok2 {
					ou_func_iu = append(ou_func_iu, &OuFuncIu{
						Ou:ou[is].(string),
						Iu:iu[js].(string),
						Func:function[it].(string),
					})
				} else if ok3 && ok4 {
					ou_func_iu = append(ou_func_iu, &OuFuncIu{
						Ou:ou[js].(string),
						Iu:iu[is].(string),
						Func:function[it].(string),
					})
				}
			}
		}
	}
	var ret string
	for _, v := range ou_func_iu {
		has, _ := utils.Engine_mysql.Exist(&models.TaroPolicy{PolicySub:v.Ou, PolicyObj:v.Iu})
		if !has {
			ret += v.Ou + "->" + v.Func + "<-" + v.Iu + " "
		}
	}
	return ret, nil
}
