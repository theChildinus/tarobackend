package models

type TaroUser struct {
	UserId         int    `json:"user_id" xorm:"not null pk autoincr comment('用户id') INT(1)"`
	UserName       string `json:"user_name" xorm:"comment('用户名') VARCHAR(255)"`
	UserRole       string `json:"user_role" xorm:"comment('用户角色') VARCHAR(255)"`
	UserDepartment string `json:"user_department" xorm:"comment('用户部门') VARCHAR(255)"`
	UserAddress    string `json:"user_address" xorm:"comment('用户地址') VARCHAR(255)"`
	UserEmail      string `json:"user_email" xorm:"comment('电子邮箱') VARCHAR(255)"`
	UserPhone      string `json:"user_phone" xorm:"comment('联系方式') VARCHAR(255)"`
	UserStatus     int    `json:"user_status" xorm:"comment('用户状态') INT(1)"`
	UserHash       string `json:"user_hash" xorm:"comment('用户哈希') VARCHAR(255)"`
	UserPath       string `json:"user_path" xorm:"comment('用户证书存储路径') VARCHAR(255)"`
}
