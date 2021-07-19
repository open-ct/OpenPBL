package models

import (
	"time"
)


type User struct {
	Id           int64     `json:"id" xorm:"not null pk autoincr"`
	Password     string    `json:"password"`
	Gender       string    `json:"gender"`
	Age          int       `json:"age"`
	Address      string    `json:"address"`
	Name         string    `json:"name"`
	Image        []byte    `json:"image"`
	CreateAt     time.Time `json:"createAt" xorm:"created"`
}

type UserInfo struct {
	Id           int64     `json:"id"`
	Gender       string    `json:"gender"`
	Age          int       `json:"age"`
	Address      string    `json:"address"`
	Name         string    `json:"name"`
	Image        []byte    `json:"image"`
}
