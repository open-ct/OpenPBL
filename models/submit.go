// Copyright 2021 The OpenPBL Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

import (
	"OpenPBL/util"
	"strconv"
	"strings"
	"time"
	"xorm.io/xorm"
)

type Submit struct {
	Id              string    `json:"id" xorm:"not null pk"`

	ProjectId       string    `json:"projectId" xorm:"not null index"`
	StudentId       string    `json:"studentId" xorm:"not null index"`
	TaskId          string    `json:"taskId" xorm:"not null index"`

	SubmitType      string    `json:"submitType" xorm:"index"`

	SubmitTitle     string    `json:"submitTitle"`
	SubmitIntroduce string    `json:"submitIntroduce" xorm:"text"`
	SubmitContent   string    `json:"submitContent" xorm:"text"`

	CreateAt        time.Time `json:"createAt"`

	Score           int       `json:"score" xorm:"default 0"`
	Scored          bool      `json:"scored" xorm:"default false"`
}

type SubmitFile struct {
	Id        string   `json:"id" xorm:"not null pk"`
	SubmitId  string   `json:"submitId" xorm:"not null index"`
	FilePath  string   `json:"filePath"`
	Name      string   `json:"name"`
	Url       string   `json:"url"`
}

type Choice struct {
	Id            string    `json:"id" xorm:"not null pk"`
	SubmitId      string    `json:"submitId" xorm:"not null index"`
	QuestionId    string    `json:"questionId" xorm:"not null index"`
	ChoiceOrder   int       `json:"choiceOrder"`
	ChoiceOptions string    `json:"choiceOptions" xorm:"text"`
}

type SubmitDetail struct {
	Submit             `json:"submit" xorm:"extends"`
	Choices   []Choice `json:"choices" xorm:"extends"`
	Submitted  bool    `json:"submitted"`
}

func (p *Submit) GetEngine() *xorm.Session {
	return adapter.Engine.Table(p)
}
func (p *SubmitFile) GetEngine() *xorm.Session {
	return adapter.Engine.Table(p)
}
func (p *SubmitFile) Create() (err error) {
	_, err = p.GetEngine().Insert(p)
	return
}

func (p *SubmitFile) Update() (err error) {
	_, err = p.GetEngine().ID(p.Id).Update(p)
	return
}

func GetSubmitFiles(sid string) (files []SectionFile, err error) {
	err = (&SubmitFile{}).GetEngine().
		Where("submit_id = ?", sid).
		Find(&files)
	return
}

func DeleteSubmitFile(fid string) (err error) {
	_, err = (&SubmitFile{}).GetEngine().ID(fid).Delete(&SubmitFile{})
	return
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
				Id:            util.NewId(),
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
	_, err = p.GetEngine().ID(p.Id).MustCols("scored").Update(p)
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
				count = string2List(question.QuestionCount)
			}
			opt := string2List(c[i].ChoiceOptions)
			for j:=0; j<len(opt); j++ {
				count[opt[j]] = count[opt[j]] + 1
			}

			if cl != nil {
				opt2 := string2List(cl[i].ChoiceOptions)
				for j:=0; j<len(opt2); j++ {
					count[opt2[j]] = count[opt2[j]] - 1
				}
			}

			question.QuestionCount = list2String(count)
			err = question.Update()
		}
	}
}

func string2List(str string) []int {
	var err error
	l := strings.Split(str, ",")
	res := make([]int, len(l))
	for i:=0; i< len(l); i++ {
		res[i], err = strconv.Atoi(l[i])
		if err != nil {
			res[i] = 0
		}
	}
	return res
}
func list2String(l []int) string {
	var s strings.Builder
	for i:=0; i< len(l); i++ {
		s.WriteString(strconv.Itoa(l[i]))
		if i != len(l)-1 {
			s.WriteString(",")
		}
	}
	return s.String()
}