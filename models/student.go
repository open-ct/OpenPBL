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
	"time"
	"xorm.io/xorm"
)

type LearnProject struct {
	Avatar     string    `json:"avatar" xorm:"text"`
	Name       string    `json:"name"`
	StudentId  string    `json:"studentId" xorm:"not null index pk"`
	ProjectId  string    `json:"projectId" xorm:"not null index pk"`
	Learning   bool      `json:"learning" xorm:"index default 0"`
	JoinTime   time.Time `json:"joinTime" xorm:"created"`
}

type LearnSection struct {
	StudentId     string    `json:"studentId" xorm:"not null pk"`
	SectionId     string    `json:"sectionId" xorm:"not null pk"`

	LearnMinute   int       `json:"learnMinute" xorm:"default 0"`
	LearnSecond   int       `json:"learnSecond" xorm:"default 0"`
}

type LastLearn struct {
	StudentId     string    `json:"studentId" xorm:"not null pk"`
	ProjectId     string    `json:"projectId" xorm:"not null pk"`
	SectionId     string    `json:"sectionId" xorm:"not null index"`
	ExitAt        time.Time `json:"exitAt" xorm:"updated"`
}

type LastLearnSection struct {
	LastLearn             `xorm:"extends"`
	Id            string  `json:"id"`
	SectionName   string  `json:"sectionName"`
	ChapterNumber int     `json:"chapterNumber"`
	SectionNumber int     `json:"sectionNumber"`
	Last          bool    `json:"last"`
}

func (l *LearnProject) GetEngine() *xorm.Session {
	return adapter.Engine.Table(l)
}
func (l *LearnSection) GetEngine() *xorm.Session {
	return adapter.Engine.Table(l)
}
func (l *LastLearn) GetEngine() *xorm.Session {
	return adapter.Engine.Table(l)
}

func (l *LearnProject) Create() (err error) {
	_, err = (&LearnProject{}).GetEngine().Insert(l)
	_, err = adapter.Engine.
		Exec("update project set join_num = join_num + 1 where id = ?", l.ProjectId)
	return
}

func (l *LearnProject) Update() (err error) {
	_, err = (&LearnProject{}).GetEngine().
		Where("student_id = ?", l.StudentId).
		Where("project_id = ?", l.ProjectId).
		MustCols("learning").
		Update(l)
	return
}

func (l *LearnProject) Delete() (err error) {
	_, err = (&LearnProject{}).GetEngine().
		Where("student_id = ?", l.StudentId).
		Where("project_id = ?", l.ProjectId).
		Delete(l)
	_, err = adapter.Engine.
		Exec("update project set join_num = join_num - 1 where id = ?", l.ProjectId)
	return
}

func IsLearningProject(pid string, uid string) (e bool) {
	var err error
	e, err = (&LearnProject{}).GetEngine().Exist(&LearnProject{
		StudentId: uid,
		ProjectId: pid,
		Learning:  true,
	})
	if err != nil {
		e = false
	}
	return
}

func GetProjectStudents(pid string, from int, size int) (s []LearnProject, rows int64, err error) {
	err = (&LearnProject{}).GetEngine().
		Where("project_id = ?", pid).
		Desc("join_time").
		Limit(size, from).
		Find(&s)
	rows, err = (&LearnProject{}).GetEngine().
		Where("project_id = ?", pid).
		Count()
	return
}

func GetLearnSection(sectionId string, studentId string, projectId string) (l LearnSection, err error) {
	var b bool
	b, err = (&LearnSection{}).GetEngine().
		Where("section_id = ?", sectionId).
		Where("student_id = ?", studentId).
		Get(&l)
	if !b {
		l = LearnSection{
			StudentId:   studentId,
			SectionId:   sectionId,
			LearnMinute: 0,
			LearnSecond: 0,
		}
		err = (&l).Create(projectId)
	}
	return
}

func GetLastLearnSection(studentId string, projectId string) (l LastLearnSection, err error) {
	var b bool
	b, err = (&LastLearn{}).GetEngine().
		Where("student_id = ?", studentId).
		Where("project_id = ?", projectId).
		Join("LEFT OUTER", Section{}, "last_learn.section_id = section.id").
		Get(&l)
	l.Last = b
	return
}

func (l *LearnSection) Create(projectId string) (err error) {
	_, err = (&LearnSection{}).GetEngine().Insert(l)
	_, err = (&LastLearn{}).GetEngine().Insert(&LastLearn{
		StudentId: l.StudentId,
		ProjectId: projectId,
		SectionId: l.SectionId,
		ExitAt:    time.Now(),
	})
	return
}
func (l *LearnSection) Update(projectId string) (err error) {
	_, err = (&LearnSection{}).GetEngine().
		Where("student_id = ?", l.StudentId).
		Where("section_id = ?", l.SectionId).
		Update(l)
	_, err = (&LastLearn{}).GetEngine().
		Where("student_id = ?", l.StudentId).
		Where("project_id = ?", projectId).
		Update(&LastLearn{
			StudentId: l.StudentId,
			ProjectId: projectId,
			SectionId: l.SectionId,
			ExitAt:    time.Now(),
		})
	return
}