package models

import (
	"openpbl-go/models/db"
	"xorm.io/xorm"
)

type Student struct {
	User                         `xorm:"extends"`
	StudentEmail       string    `json:"studentEmail" xorm:"index unique"`
	StudentPhone       string    `json:"studentPhone" xorm:"index unique"`
}

type StudentInfo struct {
	UserInfo                     `xorm:"extends"`
	StudentEmail       string    `json:"studentEmail" xorm:"index unique"`
	StudentPhone       string    `json:"studentPhone" xorm:"index unique"`
}

type LearnProject struct {
	StudentId  int64    `json:"studentId" xorm:"not null index pk"`
	ProjectId  int64    `json:"projectId" xorm:"not null index pk"`
	Learning   bool     `json:"learning" xorm:"index default 0"`
}

type DetailProject struct {
	Project          `xorm:"extends"`
	Learning   bool  `json:"learning"`
}

func (s *Student) GetEngine() *xorm.Session {
	return db.GetEngine().Table(s)
}
func (l *LearnProject) GetEngine() *xorm.Session {
	return db.GetEngine().Table(l)
}
func (sp *DetailProject) GetEngine() *xorm.Session {
	return db.GetEngine().Table(sp)
}

func (s *Student) Create() (err error) {
	_, err = s.GetEngine().Insert(s)
	return
}

func GetStudentById(id string) (s Student, err error) {
	_, err = (&Student{}).GetEngine().Where("id = ?", id).Get(&s)
	return
}

func ExistStudentEmail(email string) (exist bool, err error) {
	exist, err = (&Student{}).GetEngine().Exist("student_email = ?", email)
	return
}

func ExistStudentPhone(phone string) (exist bool, err error) {
	exist, err = (&Student{}).GetEngine().Exist("student_phone = ?", phone)
	return
}

func (l *LearnProject) Create() (err error) {
	_, err = (&LearnProject{}).GetEngine().Insert(l)
	return
}

func (l *LearnProject) Update() (err error) {
	_, err = (&LearnProject{}).GetEngine().
		Where("student_id = ?", l.StudentId).
		Where("project_id = ?", l.ProjectId).
		Update(&l)
	return
}



