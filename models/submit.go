package models

import (
	"time"
	"xorm.io/xorm"
)

type Submit struct {
	Id              int64     `json:"id" xorm:"not null pk autoincr"`

	StudentId       string    `json:"studentId" xorm:"not null index"`
	TaskId          int64     `json:"taskId" xorm:"not null index"`

	SubmitType      string    `json:"submitType" xorm:"index"`

	SubmitTitle     string    `json:"submitTitle"`
	SubmitIntroduce string    `json:"submitIntroduce" xorm:"text"`
	SubmitContent   string    `json:"submitContent" xorm:"text"`

	FilePath        string    `json:"filePath"`
	CreateAt        time.Time `json:"createAt"`
}

type Choice struct {
	Id            int64     `json:"id" xorm:"not null pk autoincr"`
	SubmitId      int64     `json:"submitId" xorm:"not null index"`
	ChoiceOrder   int       `json:"choiceOrder"`
	ChoiceOptions string    `json:"choiceOptions" xorm:"text"`
}

type SubmitDetail struct {
	Submit             `json:"submit" xorm:"extends"`
	Choices   []Choice `json:"choices" xorm:"extends"`
	Submitted  bool     `json:"submitted"`
}

func (p *Submit) GetEngine() *xorm.Session {
	return adapter.Engine.Table(p)
}
func (c *Choice) GetEngine() *xorm.Session {
	return adapter.Engine.Table(c)
}
func (c *Choice) Create() (err error) {
	_, err = c.GetEngine().Insert(c)
	return
}
func (c *Choice) Update() (err error) {
	_, err = c.GetEngine().ID(c.Id).Update(c)
	return
}

func (p *Submit) Create(c []Choice) (err error) {
	_, err = p.GetEngine().Insert(p)
	if p.SubmitType == "survey" {
		for i:=0; i< len(c); i++ {
			ci := &Choice{
				SubmitId:      p.Id,
				ChoiceOrder:   c[i].ChoiceOrder,
				ChoiceOptions: c[i].ChoiceOptions,
			}
			err = ci.Create()
		}
	}
	return
}
func (p *Submit) Update(c []Choice) (err error) {
	_, err = p.GetEngine().ID(p.Id).Update(p)
	if p.SubmitType == "survey" {
		for i:=0; i< len(c); i++ {
			err = (&c[i]).Update()
		}
	}
	return
}