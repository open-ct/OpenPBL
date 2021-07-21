package models

import (
	"openpbl-go/models/db"
	"xorm.io/xorm"
)

type LearnProject struct {
	StudentId  string   `json:"studentId" xorm:"not null index pk"`
	ProjectId  int64    `json:"projectId" xorm:"not null index pk"`
	Learning   bool     `json:"learning" xorm:"index default 0"`
}

type DetailProject struct {
	Project          `xorm:"extends"`
	Learning   bool  `json:"learning"`
}

func (l *LearnProject) GetEngine() *xorm.Session {
	return db.GetEngine().Table(l)
}
func (sp *DetailProject) GetEngine() *xorm.Session {
	return db.GetEngine().Table(sp)
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



