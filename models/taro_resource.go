package models

import (
	"time"
)

type TaroResource struct {
	ResourceId    int       `json:"resource_id" xorm:"not null pk autoincr comment('资源id') INT(1)"`
	ResourceName  string    `json:"resource_name" xorm:"comment('资源名') VARCHAR(255)"`
	ResourceType  int       `json:"resource_type" xorm:"comment('资源类型') INT(10)"`
	ResourceCtime time.Time `json:"resource_ctime" xorm:"comment('创建时间') DATETIME"`
}
