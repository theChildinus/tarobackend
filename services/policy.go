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
	"time"
)

type PolicyReq struct {
	PageIndex int64  `json:"page_index"`
	PageSize  int64  `json:"page_size"`
	SearchSub string `json:"search_sub"`
}

type PolicyCheckReq struct {
	PolicyName string `json:"policyname"`
	PolicySub  string `json:"policysub"`
	PolicyObj  string `json:"policyobj"`
	PolicyAct  string `json:"policyact"`
	UserName   string `json:"username"`
	UserHash   string `json:"userhash"`
}

type PolicyResp struct {
	List  []models.TaroPolicy `json:"list"`
	Count int64               `json:"count"`
}

type RoleAllotReq struct {
	Name  string   `json:"name"`
	Roles []string `json:"roles"`
}

type PolicyModel struct {
	PolicyName string `json:"policy_name"`
	ModelType  string `json:"model_type"`
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
			Where("policy_name like ? ", "%"+req.SearchSub+"%").
			Limit(int(req.PageSize), int((req.PageIndex-1)*req.PageSize)).
			Find(&policies)
		count, _ = engine.Where("policy_name like ? ", "%"+req.SearchSub+"%").Count(m)
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
	objs := strings.Split(r.PolicyObj, "#")
	var policys []models.TaroPolicy
	for _, obj := range objs {
		if len(obj) != 0 {
			policys = append(policys, models.TaroPolicy{
				PolicyName:  r.PolicyName,
				PolicySub:   r.PolicySub,
				PolicyObj:   obj,
				PolicyAct:   r.PolicyAct,
				PolicyType:  r.PolicyType,
			})
		}
	}
	// delete existed policy rule
	j := 0
	for _, p := range policys {
		if has := enf.HasPolicy(p.PolicySub, p.PolicyObj, p.PolicyAct); !has {
			policys[j] = p
			j++
		}
	}
	policys = policys[:j]
	// save to casbin csvfile
	for _, p := range policys {
		ret := enf.AddPolicy(p.PolicySub, p.PolicyObj, p.PolicyAct)
		if !ret {
			logs.Error("Add policy rule Error: ", p)
			return false, errors.New("Add policy rule Error")
		}
	}
	// save to db
	engine := utils.Engine_mysql
	_, err = engine.Insert(&policys)
	if err != nil {
		logs.Error("CreatePolicy: Table Policy InsertOne Error")
		return false, err
	}
	if model_type == "RBAC" {
		var users []models.TaroUser
		err = engine.Table("taro_user").
			Where("user_role like ? ", "%"+r.PolicySub+"%").Find(&users)
		//logs.Info("users:", r.PolicySub, users)
		for _, v := range users {
			_ = enf.AddRoleForUser(v.UserName, r.PolicySub)
		}
	}
	_ = enf.SavePolicy()

	// add policy to fabric configtx.yaml file when policyName matching
	if r.PolicyName == beego.AppConfig.String("fabric_policy_name") {
		tx, err := utils.ParseYamlFile(beego.AppConfig.String("fabric_configtx"))
		if err == nil {
			if len(tx.Application.ACLs) == 0 {
				tx.Application.ACLs = make(map[string]string)
			}
			for _, p := range policys {
				subStr := p.PolicyObj[strings.Index(p.PolicyObj, "/")+1:]
				tx.Application.ACLs[subStr] = p.PolicySub
			}
			if err = utils.SaveYamlFile(tx, beego.AppConfig.String("fabric_configtx")); err != nil {
				return false, err
			}
		} else {
			return false, err
		}
	}
	return true, nil
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
			subStr := r.PolicyObj[strings.Index(r.PolicyObj, "/")+1:]
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
			subStr := r.PolicyObj[strings.Index(r.PolicyObj, "/")+1:]
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
	if len(r.UserHash) == 0 || len(r.UserName) == 0 || len(r.PolicyName) == 0 {
		logs.Error("CheckPolicy: UserHash or UserName or PolicyName Empty")
		return false, errors.New("CheckPolicy: UserHash or UserName or PolicyName Empty")
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
		(strings.Contains(r.PolicySub, m.UserName) || r.PolicySub == m.UserRole) {
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
	m := &models.TaroEnum{EnumKey: "mutex_role"}
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
	req := &models.TaroEnum{EnumKey: "policy_model"}
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
		Ou   string // organization unit pel
		Iu   string // information unit pel
		Func string // function pel
	}
	var ou_func_iu []*OuFuncIu
	removeMult := make(map[string]int)
	function := make(map[string]string)
	ou := make(map[string]string)
	iu := make(map[string]string)
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
	for i := 0; i < len(arc)-1; i++ {
		for j := 1; j < len(arc); j++ {
			is, it, js, jt := arc[i].Flow.Source, arc[i].Flow.Target,
				arc[j].Flow.Source, arc[j].Flow.Target
			if _, ok := function[it]; ok && it == jt {
				_, ok1 := ou[is]
				_, ok2 := iu[js]
				_, ok3 := ou[js]
				_, ok4 := iu[is]
				if ok1 && ok2 {
					o, i, f := ou[is], iu[js], function[it]
					if _, exist := removeMult[o+"#"+i+"#"+f]; !exist {
						ou_func_iu = append(ou_func_iu, &OuFuncIu{Ou: o, Iu: i, Func: f})
						removeMult[o+"#"+i+"#"+f] = 0
					}
				} else if ok3 && ok4 {
					o, i, f := ou[js], iu[is], function[it]
					if _, exist := removeMult[o+"#"+i+"#"+f]; !exist {
						ou_func_iu = append(ou_func_iu, &OuFuncIu{Ou: o, Iu: i, Func: f})
						removeMult[o+"#"+i+"#"+f] = 0
					}
				}
			}
		}
	}
	var ret string
	for _, v := range ou_func_iu {
		// TODO: v.Ou / v.Iu in db ?
		policy_name := beego.AppConfig.String("epc_policy_name")
		_, err := CreatePolicy(&models.TaroPolicy{
			PolicyName:  policy_name,
			PolicySub:   v.Ou,
			PolicyObj:   v.Iu,
			PolicyAct:   v.Func,
			PolicyType:  "",
			PolicyCtime: time.Time{},
		})
		if err != nil {
			str := v.Ou + "->" + v.Func + "<-" + v.Iu
			logs.Error("CreatePolicy: ", str, " Existed Or Error")
			ret += str
		}
	}
	return ret, nil
}
