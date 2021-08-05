package models

import (
	"strconv"
	"strings"
	"time"
	"xorm.io/xorm"
)

type Submit struct {
	Id              int64     `json:"id" xorm:"not null pk autoincr"`

	ProjectId       int64     `json:"projectId" xorm:"not null index"`
	StudentId       string    `json:"studentId" xorm:"not null index"`
	TaskId          int64     `json:"taskId" xorm:"not null index"`

	SubmitType      string    `json:"submitType" xorm:"index"`

	SubmitTitle     string    `json:"submitTitle"`
	SubmitIntroduce string    `json:"submitIntroduce" xorm:"text"`
	SubmitContent   string    `json:"submitContent" xorm:"text"`

	FilePath        string    `json:"filePath"`
	CreateAt        time.Time `json:"createAt"`

	Score           int       `json:"score" xorm:"default 0"`
	Scored          bool      `json:"scored" xorm:"default false"`
}

type Choice struct {
	Id            int64     `json:"id" xorm:"not null pk autoincr"`
	SubmitId      int64     `json:"submitId" xorm:"not null index"`
	QuestionId    int64     `json:"questionId" xorm:"not null index"`
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
		for i := 0; i < len(c); i ++ {
			ci := &Choice{
				SubmitId:      p.Id,
				QuestionId:    c[i].QuestionId,
				ChoiceOrder:   c[i].ChoiceOrder,
				ChoiceOptions: c[i].ChoiceOptions,
			}
			err = ci.Create()
		}
		CountSubmit(c, nil)
	}
	return
}
func (p *Submit) Update(c []Choice) (err error) {
	_, err = p.GetEngine().ID(p.Id).Update(p)
	if len(c) > 0 && p.SubmitType == "survey" {
		var cs []Choice
		err = (&Choice{}).GetEngine().
			Where("submit_id = ?", p.Id).
			Find(&cs)
		for i:=0; i< len(c); i++ {
			err = (&c[i]).Update()
		}
		CountSubmit(c, cs)
	}
	return
}

func CountSubmit(c []Choice, cl []Choice) {
	if len(c) != len(cl) && cl != nil {
		return
	}
	var (
		err error
		index int
	)
	for i:=0; i< len(c); i++ {
		var question Question
		_, err = (&Question{}).GetEngine().ID(c[i].QuestionId).Get(&question)
		if err != nil {
			return
		}
		if question.QuestionType == "singleChoice" || question.QuestionType == "multipleChoice" ||
			question.QuestionType == "scale5" || question.QuestionType == "scale7" {

			l := len(strings.Split(question.QuestionOptions, ","))
			count := make([]int, l)

			if question.QuestionCount != "" {
				le := strings.Split(question.QuestionCount, ",")
				for j:=0; j< len(le); j++ {
					count[i], err = strconv.Atoi(le[i])
				}
			}
			ci := strings.Split(c[i].ChoiceOptions, ",")
			var cli []string
			if cl != nil {
				cli = strings.Split(cl[i].ChoiceOptions, ",")
			}

			for j:=0; j<len(ci); j++ {
				index, err = strconv.Atoi(ci[j])
				count[index] = count[index] + 1
			}
			if cl != nil {
				for j:=0; j<len(cli); j++ {
					index, err = strconv.Atoi(cli[j])
					count[index] = count[index] - 1
				}
			}

			var qc strings.Builder
			for j:=0; j< len(count); j++ {
				qc.WriteString(strconv.Itoa(count[j]))
				if j != len(count) - 1 {
					qc.WriteString(",")
				}
			}
			question.QuestionCount = qc.String()
			err = question.Update()
		}
	}
}