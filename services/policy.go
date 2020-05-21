package services

import (
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
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
	PolocyEnv  string `json:"policyenv"`
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
	EpcCtx string `json:"epc_ctx"`
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
				PolicyCtime: time.Now(),
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
	if model_type == "RBAC" || model_type == "ABAC" {
		var usernames []string
		err = engine.Table("taro_user").
			Where("user_role like ? ", "%"+r.PolicySub+"%").
			Cols("user_name").Find(&usernames)
		logs.Info("[CreatePolicy]: policysub:", r.PolicySub, "users:", usernames)
		for _, v := range usernames {
			_ = enf.AddRoleForUser(v, r.PolicySub)
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

func DeletePolicyById(ids []int) (bool, error) {
	engine := utils.Engine_mysql
	var ps []models.TaroPolicy

	err := engine.Table("taro_policy").In("policy_id", ids).Find(&ps)
	if err != nil {
		logs.Error("DeletePolicyById: Table Policy Find Error")
		return false, err
	}

	for _, r := range ps {
		model_type, err := GetModelType(r.PolicyName)
		if err != nil {
			logs.Error("DeletePolicy: Get ModelType Error")
			return false, err
		}
		casbin_model := "./casbinfiles/" + strings.ToLower(model_type) + "_model.conf"
		casbin_policys := "./casbinfiles/" + r.PolicyName + ".csv"
		if ok, err := utils.FileExistAndCreate(casbin_policys); !ok {
			return false, err
		}
		enf := casbin.NewEnforcer(casbin_model, casbin_policys)
		ret := enf.RemovePolicy(r.PolicySub, r.PolicyObj, r.PolicyAct)
		if !ret {
			logs.Error("DeletePolicy: Delete", r, " Failed")
			return false, nil
		}
		if model_type == "RBAC" || model_type == "ABAC" {
			subs, exist := enf.GetAllSubjects(), false
			for _, sub := range subs {
				if r.PolicySub == sub {
					exist = true
					break
				}
			}
			if !exist {
				usernames, _ := enf.GetUsersForRole(r.PolicySub)
				logs.Info("[DeletePolicy]: policysub:", r.PolicySub, " users:", usernames)
				for _, v := range usernames {
					_ = enf.DeleteRoleForUser(v, r.PolicySub)
				}
			}
		}
		_ = enf.SavePolicy()

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
	}

	var m models.TaroPolicy
	_, err = engine.Table("taro_policy").In("policy_id", ids).Delete(&m)
	if err != nil {
		logs.Error("DeletePolicyById: Table Policy Delete Error")
		return false, err
	}
	return true, nil
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
		if model_type == "RBAC" || model_type == "ABAC" {
			// remove policy rule: users has old role
			subs, exist := enf.GetAllSubjects(), false
			for _, sub := range subs {
				if r.PolicySub == sub {
					exist = true
					break
				}
			}
			if !exist {
				usernames, _ := enf.GetUsersForRole(r.PolicySub)
				logs.Info("[DeletePolicy]: policysub:", r.PolicySub, " users:", usernames)
				for _, v := range usernames {
					_ = enf.DeleteRoleForUser(v, r.PolicySub)
				}
			}
			// add policy rule: users has new role
			var usernames []string
			err = engine.Table("taro_user").
				Where("user_role like ? ", "%"+old.PolicySub+"%").
				Cols("user_name").Find(&usernames)
			for _, v := range usernames {
				_ = enf.AddRoleForUser(v, r.PolicySub)
			}
		}
		_ = enf.SavePolicy()
		ret = ret1 && ret2
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
	logs.Info("[CheckPolicy] Req:", r, "UserHash: ", m.UserHash)
	if has &&
		r.UserHash == m.UserHash &&
		(strings.Contains(r.PolicySub, m.UserName) || r.PolicySub == m.UserRole) {
		model_type, err := GetModelType(r.PolicyName)
		if err != nil {
			logs.Error("CheckPolicy: Get Model Type Error")
			return false, err
		}
		casbin_model := "./casbinfiles/" + strings.ToLower(model_type) + "_model.conf"
		casbin_policys := "./casbinfiles/" + r.PolicyName + ".csv"
		logs.Info("[CheckPolicy] modelfile:", casbin_model, "policyfile:", casbin_policys)
		enf := casbin.NewEnforcer(casbin_model, casbin_policys)
		if model_type == "ABAC" {
			var sa SubAttr
			var oa ObjAttr
			var aa ActAttr
			var ea EnvAttr
			if err := json.Unmarshal([]byte(r.PolicySub), &sa); err != nil {
				logs.Error("json Unmarshal PolicySub Error")
				return false, err
			}
			if err := json.Unmarshal([]byte(r.PolicyObj), &oa); err != nil {
				return false, err
			}
			if err := json.Unmarshal([]byte(r.PolicyAct), &aa); err != nil {
				return false, err
			}
			if err := json.Unmarshal([]byte(r.PolocyEnv), &ea); err != nil {
				return false, err
			}
			logs.Info("SubAttr=>", sa, "ObjAttr=>", oa, "ActAttr=>", aa, "EnvAttr=>", ea)
			enf.AddFunction("obj_func", ObjKeyMatchFunc)
			enf.AddFunction("act_func", ActKeyMatchFunc)
			ret := enf.Enforce(sa, oa, aa, ea)
			logs.Info("[CheckPolicy]: RESULT:", ret)
			return ret, nil
		} else {
			ret := enf.Enforce(r.PolicySub, r.PolicyObj, r.PolicyAct)
			logs.Info("[CheckPolicy]: RESULT:", ret)
			return ret, nil
		}
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

func Executable(epcCtx string) (int64, error) {
	bytes, _ := base64.StdEncoding.DecodeString(epcCtx)
	var e utils.EpcCtx
	var err error
	if err = xml.Unmarshal(bytes, &e); err != nil {
		logs.Error("XML Parse Error: ", err.Error())
		return -1, errors.New("XML Parse Error: " + err.Error())
	}

	type OuFuncIu struct {
		Ou   string // organization unit pel
		Iu   string // information unit pel
		Func string // function pel
	}

	type FuncEventOpr struct {
		Func string // function pel
		Event string // event pel
		Opr string // relational operator pel
	}

	var ou_func_iu []*OuFuncIu
	var func_event_opr []*FuncEventOpr

	ofi_set := make(map[string]int)
	feo_set := make(map[string]int)
	function := make(map[string]string)
	ou := make(map[string]string)
	iu := make(map[string]string)
	event := make(map[string]string)
	epc_and := make(map[string]string)
	epc_or := make(map[string]string)
	epc_xor := make(map[string]string)
	for _, v := range e.Epc.Function {
		v.Name = strings.Replace(v.Name, " ", "", -1)
		v.Name = strings.Replace(v.Name, "\n", "", -1)
		function[v.ID] = v.Name
	}
	for _, v := range e.Epc.Ou {
		v.OuName = strings.Replace(v.OuName, "\n", "", -1)
		ou[v.ID] = v.OuName
	}
	for _, v := range e.Epc.Iu {
		v.IuName = strings.Replace(v.IuName, "<div>", "", -1)
		v.IuName = strings.Replace(v.IuName, "</div>", "", -1)
		v.IuName = strings.Replace(v.IuName, " ", "", -1)
		v.IuName = strings.Replace(v.IuName, "\n", "", -1)
		iu[v.ID] = v.IuName
	}
	for _, v := range e.Epc.Event {
		v.Name = strings.Replace(v.Name, "\n", "", -1)
		event[v.ID] = v.Name
	}
	for _, v := range e.Epc.And {
		epc_and[v.ID] = "AND"
	}
	for _, v := range e.Epc.Or {
		epc_or[v.ID] = "OR"
	}
	for _, v := range e.Epc.Xor {
		epc_xor[v.ID] = "XOR"
	}
	arc := e.Epc.Arc
	for i := 0; i < len(arc)-1; i++ {
		for j := i + 1; j < len(arc); j++ {
			is, it, js, jt := arc[i].Flow.Source, arc[i].Flow.Target,
				arc[j].Flow.Source, arc[j].Flow.Target
			// find ou -> func <- iu
			if _, ok := function[it]; ok && it == jt {
				_, ok1 := ou[is]
				_, ok2 := iu[js]
				_, ok3 := ou[js]
				_, ok4 := iu[is]
				if ok1 && ok2 {
					o, i, f := ou[is], iu[js], function[it]
					if _, exist := ofi_set[o+"#"+i+"#"+f]; !exist {
						ou_func_iu = append(ou_func_iu, &OuFuncIu{Ou: o, Iu: i, Func: f})
						ofi_set[o+"#"+i+"#"+f] = 0
					}
				} else if ok3 && ok4 {
					o, i, f := ou[js], iu[is], function[it]
					if _, exist := ofi_set[o+"#"+i+"#"+f]; !exist {
						ou_func_iu = append(ou_func_iu, &OuFuncIu{Ou: o, Iu: i, Func: f})
						ofi_set[o+"#"+i+"#"+f] = 0
					}
				}
			}

			// find func -> event -> operator
			if _, ok := epc_and[jt]; ok {
				if _, ok := event[it]; ok && it == js {
					f, e, o := function[is], event[it], epc_and[jt]
					if _, exist := feo_set[f+"#"+e+"#"+o]; !exist {
						func_event_opr = append(func_event_opr, &FuncEventOpr{Func:f,Event:e,Opr:o})
						feo_set[f+"#"+e+"#"+o] = 0
					}
				}
			}

			if _, ok := epc_or[jt]; ok {
				if _, ok := event[it]; ok && it == js {
					f, e, o := function[is], event[it], epc_or[jt]
					if _, exist := feo_set[f+"#"+e+"#"+o]; !exist {
						func_event_opr = append(func_event_opr, &FuncEventOpr{Func:f,Event:e,Opr:o})
						feo_set[f+"#"+e+"#"+o] = 0
					}
				}
			}
		}
	}

	logs.Info("##### 从EPC模型中解析: 组织单元->函数<-信息单元 #####")
	for _, v := range ou_func_iu {
		fmt.Println(v.Ou + "->" + v.Func + "<-" + v.Iu)
	}
	logs.Info("##### 从EPC模型中解析:  函数->事件->关系 #####")
	for _, v := range func_event_opr {
		fmt.Println(v.Func + "->" + v.Event + "->" + v.Opr)
	}
	for i := 0; i < len(ou_func_iu); i++ {
		for j := i + 1; j < len(ou_func_iu); j++ {
			role1, role2 := ou_func_iu[i].Ou, ou_func_iu[j].Ou
			func1, func2 := ou_func_iu[i].Func, ou_func_iu[j].Func
			if role1 != role2 && func1 == func2 {
				req := &RoleAllotReq{Roles: []string{role1, role2}}
				ret, _ := RoleAllot(req)
				if !ret {
					logs.Error("The mutex role accesses same functions")
					return -1, errors.New("The mutex role accesses same functions")
				}
			}
		}
	}

	casbin_model := "./casbinfiles/" + "rbac_model.conf"
	casbin_policys := "./casbinfiles/" + beego.AppConfig.String("epc_policy_name") + ".csv"
	enf := casbin.NewEnforcer(casbin_model, casbin_policys)
	for _, ofi := range ou_func_iu {
		for _, feo := range func_event_opr {
			if ofi.Func == feo.Func && feo.Opr == "AND" {
				if has := enf.HasPolicy(ofi.Ou); !has {
					errmsg := "角色 "+ ofi.Ou + " 未被授权"
					logs.Error(errmsg)
					return -1, errors.New(errmsg)
				}
			}
		}
	}

	//for _, v := range ou_func_iu {
	//	// TODO: v.Ou / v.Iu in db ?
	//	policy_name := beego.AppConfig.String("epc_policy_name")
	//	_, err := CreatePolicy(&models.TaroPolicy{
	//		PolicyName:  policy_name,
	//		PolicySub:   v.Ou,
	//		PolicyObj:   v.Iu,
	//		PolicyAct:   v.Func,
	//		PolicyType:  "",
	//		PolicyCtime: time.Time{},
	//	})
	//	if err != nil {
	//		str := v.Ou + "->" + v.Func + "<-" + v.Iu
	//		logs.Error("CreatePolicy: ", str, " Error")
	//	}
	//}
	logs.Info("角色分配未发生冲突，且均已授权")
	return 0, nil
}
