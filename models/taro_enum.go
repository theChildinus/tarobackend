package models

type TaroEnum struct {
	EnumId    int    `json:"enum_id" xorm:"not null pk autoincr INT(1)"`
	EnumKey   string `json:"enum_key" xorm:"VARCHAR(255)"`
	EnumValue string `json:"enum_value" xorm:"VARCHAR(2048)"`
}
