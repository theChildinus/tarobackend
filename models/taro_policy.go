package models

import (
	"time"
)

type TaroPolicy struct {
	PolicyId    int       `json:"policy_id" xorm:"not null pk autoincr comment('策略id') INT(1)"`
	PolicyName  string    `json:"policy_name" xorm:"comment('策略名称') VARCHAR(255)"`
	PolicySub   string    `json:"policy_sub" xorm:"comment('策略规则主体') VARCHAR(255)"`
	PolicyObj   string    `json:"policy_obj" xorm:"comment('策略规则资源') VARCHAR(255)"`
	PolicyAct   string    `json:"policy_act" xorm:"comment('策略规则动作') VARCHAR(255)"`
	PolicyType  string    `json:"policy_type" xorm:"comment('策略类型') VARCHAR(255)"`
	PolicyCtime time.Time `json:"policy_ctime" xorm:"comment('创建时间') DATETIME"`
}
