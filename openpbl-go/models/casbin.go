package models


type CasbinRule struct {
	Id    int64      `json:"id" xorm:"not null pk autoincr"`// 自增主键
	PType string     `json:"pType"`// Policy Type - 用于区分 policy和 group(role)
	V0    string     `json:"v0"`// subject
	V1    string     `json:"v1"`// object
	V2    string     `json:"v2"`// action
}