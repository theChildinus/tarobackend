package models

import (
	"time"
)

type TaroIdentity struct {
	IdentityId          int       `json:"identity_id" xorm:"not null pk autoincr comment('Fabric 注册Id') INT(1)"`
	IdentityName        string    `json:"identity_name" xorm:"comment('Fabric 参与者名') VARCHAR(255)"`
	IdentitySecret      string    `json:"identity_secret" xorm:"comment('Fabric 参与者密码') VARCHAR(255)"`
	IdentityType        string    `json:"identity_type" xorm:"comment('Fabric 参与者类型') VARCHAR(255)"`
	IdentityAffiliation string    `json:"identity_affiliation" xorm:"comment('Fabric 参与者从属关系') VARCHAR(255)"`
	IdentityAttrs       string    `json:"identity_attrs" xorm:"comment('Fabric 参与者属性') VARCHAR(1024)"`
	IdentityCtime       time.Time `json:"identity_ctime" xorm:"comment('Fabric 参与者创建时间') DATETIME"`
	IdentityStatus      int       `json:"identity_status" xorm:"comment('Fabric 参与者状态') INT(1)"`
	IdentityIp          string    `json:"identity_ip" xorm:"comment('Fabric 参与者主机IP') VARCHAR(20)"`
	IdentityUser        string    `json:"identity_user" xorm:"comment('Fabric 参与者主机名') VARCHAR(255)"`
	IdentityPw          string    `json:"identity_pw" xorm:"comment('Fabric 参与者主机密码') VARCHAR(255)"`
	IdentityPath        string    `json:"identity_path" xorm:"comment('Fabric 参与者主机路径') VARCHAR(255)"`
}
