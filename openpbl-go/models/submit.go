package models

import (
	"time"
	"xorm.io/xorm"
)

type Submit struct {
	Id              int64     `json:"id" xorm:"not null pk autoincr"`

	StudentId       int64     `json:"studentId" xorm:"not null index"`
	TaskId          int64     `json:"taskId" xorm:"not null index"`

	SubmitTitle     string    `json:"submitTitle"`
	SubmitIntroduce string    `json:"submitIntroduce" xorm:"text"`

	FilePath        string    `json:"filePath"`


	CreateAt        time.Time `json:"createAt" xorm:"created"`
}


func GetSubmitFiles(sid string, pid string) (f[] Submit, err error) {
	err = (&Submit{}).GetEngine().
		Where("project_id = ? and student_id = ?", pid, sid).
		Find(&f)
	return
}

func (p *Submit) GetEngine() *xorm.Session {
	return adapter.Engine.Table(p)
}

func (p *Submit) Create() (err error) {
	_, err = p.GetEngine().Insert(p)
	return
}