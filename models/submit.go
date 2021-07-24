package models

import (
	"time"
	"xorm.io/xorm"
)

type Submit struct {
	Id              int64     `json:"id" xorm:"not null pk autoincr"`

	StudentId       int64     `json:"studentId" xorm:"not null index"`
	TaskId          int64     `json:"taskId" xorm:"not null index"`

	SubmitType      string    `json:"submitType" xorm:"index"`

	SubmitTitle     string    `json:"submitTitle"`
	SubmitIntroduce string    `json:"submitIntroduce" xorm:"text"`

	FilePath        string    `json:"filePath"`
	CreateAt        time.Time `json:"createAt" xorm:"created"`

	SurveyId      int64     `json:"surveyId" xorm:"index"`
}

func (p *Submit) GetEngine() *xorm.Session {
	return adapter.Engine.Table(p)
}

func (p *Submit) Create() (err error) {
	_, err = p.GetEngine().Insert(p)
	return
}

func GetStudentTaskSubmit(tid string, sid string) (t Task, err error) {
	_, err = (&Submit{}).GetEngine().
		Where("task_id = ?", tid).
		Where("student_id = ?", sid).
		Join("LEFT OUTER", Survey{}, "survey.id = task.survey_id").
		Get(&t)
	return
}