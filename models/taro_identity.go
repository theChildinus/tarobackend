package models

import (
	"time"
)

type TaroIdentity struct {
	IdentityId          int       `json:"identity_id" xorm:"not null pk autoincr comment('Fabric 注册Id') INT(1)"`
	IdentityName        string    `json:"identity_name" xorm:"comment('Fabric 身份名') VARCHAR(255)"`
	IdentitySecret      string    `json:"identity_secret" xorm:"comment('Fabric 身份密码') VARCHAR(255)"`
	IdentityType        string    `json:"identity_type" xorm:"comment('Fabric 身份类型') VARCHAR(255)"`
	IdentityAffiliation string    `json:"identity_affiliation" xorm:"comment('Fabric 身份从属关系') VARCHAR(255)"`
	IdentityAttrs       string    `json:"identity_attrs" xorm:"comment('Fabric 身份属性') VARCHAR(1024)"`
	IdentityCtime       time.Time `json:"identity_ctime" xorm:"comment('Fabric 身份创建时间') DATETIME"`
	IdentityStatus      int       `json:"identity_status" xorm:"comment('Fabric 身份状态') INT(1)"`
}
