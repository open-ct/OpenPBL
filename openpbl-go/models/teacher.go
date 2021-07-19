package models

import (
	"openpbl-go/models/db"
	"xorm.io/xorm"
)

type Teacher struct {
	User                         `xorm:"extends"`
	TeacherEmail       string    `json:"teacherEmail" xorm:"index unique"`
	TeacherPhone       string    `json:"teacherPhone" xorm:"index unique"`
}

type ProjectDetail struct {
	Project            `xorm:"extends"`
	Name     string    `json:"name"`
	Learning bool      `json:"learning"`
}

func (t *Teacher) GetEngine() *xorm.Session {
	return db.GetEngine().Table(t)
}

func (t * Teacher) Create() (err error) {
	_, err = t.GetEngine().Insert(t)
	return
}

func ExistTeacherEmail(email string) (exist bool, err error) {
	exist, err = (&Teacher{}).GetEngine().Exist("teacher_email = ?", email)
	return
}

func ExistTeacherPhone(phone string) (exist bool, err error) {
	exist, err = (&Teacher{}).GetEngine().Exist("teacher_phone = ?", phone)
	return
}
