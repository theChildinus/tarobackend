package models

import (
	"time"
)

type TaroPolicy struct {
	PolicyId    int       `json:"policy_id" xorm:"not null pk autoincr comment('策略id') INT(1)"`
	PolicySub   string    `json:"policy_sub" xorm:"comment('策略主体') VARCHAR(255)"`
	PolicyObj   string    `json:"policy_obj" xorm:"comment('策略资源') VARCHAR(255)"`
	PolicyAct   string    `json:"policy_act" xorm:"comment('策略动作') VARCHAR(255)"`
	PolicyCtime time.Time `json:"policy_ctime" xorm:"comment('创建时间') DATETIME"`
}
