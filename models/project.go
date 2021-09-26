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
	"errors"
	"time"
	"xorm.io/xorm"
)


type Project struct {
	Id                 string    `json:"id" xorm:"not null pk"`
	Image              string    `json:"image"`
	ProjectTitle       string    `json:"projectTitle"`
	ProjectIntroduce   string    `json:"projectIntroduce" xorm:"text"`
	ProjectGoal        string    `json:"projectGoal" xorm:"text"`

	TeacherId          string    `json:"teacherId" xorm:"index"`

	Subjects           string    `json:"subjects" xorm:"default ''"`
	Skills             string    `json:"skills" xorm:"default ''"`

	CreateAt           time.Time `json:"createAt" xorm:"created"`
	UpdateAt           time.Time `json:"UpdateAt"`

	Closed             bool      `json:"closed" xorm:"default false index"`
	ClosedAt           time.Time `json:"closedAt"`

	Published          bool      `json:"published" xorm:"default false index"`
	PublishedAt        time.Time `json:"publishedAt"`

	TimedEnd           bool      `json:"timedEnd"`
	EndTime            time.Time `json:"endTime"`

	ReadNum            int64     `json:"readNum" xorm:"default 0"`
	JoinNum            int64     `json:"joinNum" xorm:"default 0"`

	LearnMinuteWeight  int       `json:"learnMinuteWeight" xorm:"default 100"`
}

type Favourite struct {
	ProjectId          string    `json:"projectId" xorm:"not null pk"`
	UserId             string    `json:"userId" xorm:"not null pk"`
	CreateAt           time.Time `json:"createAt" xorm:"created"`
}

type ProjectDetail struct {
	Project             `xorm:"extends"`
	Learning  bool      `json:"learning"`
	Created   bool      `json:"created"`
	Favourite bool      `json:"favourite"`
}

type ProjectSkill struct {
	Skill          string     `json:"skill" xorm:"not null pk"`
	ProjectId      string     `json:"projectId" xorm:"not null pk"`
}

type ProjectSubject struct {
	Subject        string     `json:"subject" xorm:"not null pk"`
	ProjectId      string     `json:"projectId" xorm:"not null pk"`
}

func (p *Project) GetEngine() *xorm.Session {
	return adapter.Engine.Table(p)
}
func (p *ProjectSkill) GetEngine() *xorm.Session {
	return adapter.Engine.Table(p)
}
func (p *ProjectSubject) GetEngine() *xorm.Session {
	return adapter.Engine.Table(p)
}
func (f *Favourite) GetEngine() *xorm.Session {
	return adapter.Engine.Table(f)
}

func GetProjectById(pid string) (project Project, err error) {
	var b bool
	b, err = (&Project{}).GetEngine().
		ID(pid).
		Get(&project)
	if !b {
		err = errors.New("404")
	}
	return
}

func GetProjectByPidForTeacher(pid string, uid string) (pd ProjectDetail, err error) {
	var p Project
	c, err := (&Project{}).GetEngine().
		ID(pid).
		Get(&p)
	favourite, _ := (&Favourite{}).GetEngine().
		Exist(&Favourite{
			ProjectId: pid,
			UserId:    uid,
		})

	created := uid == p.TeacherId
	pd = ProjectDetail{
		Project:  p,
		Learning: false,
		Created: created,
		Favourite: favourite,
	}
	if !c {
		err = errors.New("404")
	}
	return
}

func GetProjectByPidForStudent(pid string, uid string) (pd ProjectDetail, err error) {
	c, err := (&Project{}).GetEngine().
		Where("project.id = ?", pid).
		Join("LEFT OUTER", LearnProject{}, "project.id = learn_project.project_id and student_id = ?", uid).
		Get(&pd)
	favourite, _ := (&Favourite{}).GetEngine().
		Exist(&Favourite{
			ProjectId: pid,
			UserId: uid,
		})
	pd.Favourite = favourite
	if !c {
		err = errors.New("404")
	}
	return
}


func (p *Project) Create() (err error) {
	_, err = p.GetEngine().Insert(p)
	return
}

func (p *ProjectSkill) Create() (err error) {
	_, err = p.GetEngine().Insert(p)
	return
}
func (p *ProjectSubject) Create() (err error) {
	_, err = p.GetEngine().Insert(p)
	return
}

func (p *Project) Update() (err error) {
	_, err = p.GetEngine().ID(p.Id).Update(p)
	return
}

func UpdateWeight(p Project, t[]Task) (err error) {
	session := adapter.Engine.NewSession()
	defer session.Close()
	session.Begin()
	_, err = session.Table(&Project{}).ID(p.Id).Update(&p)
	for i:=0; i< len(t); i++ {
		(&t[i]).Update()
	}
	session.Commit()
	return
}

func (p *Project) UpdateInfo(subjects []*ProjectSubject, skills []*ProjectSkill) (err error) {
	_, err = p.GetEngine().ID(p.Id).MustCols("timed_end").Update(p)
	n1 := len(subjects)
	for i:=0; i<n1; i++ {
		_ = subjects[i].Create()
	}
	n2 := len(skills)
	for i:=0; i<n2; i++ {
		_ = skills[i].Create()
	}
	return
}

func (p *Project) Delete() (err error) {
	pid := p.Id
	_, err = p.GetEngine().ID(p.Id).Delete(p)
	err = DeleteProjectChapters(pid)
	return
}

func UpdatePublished(p Project) (err error) {
	_, err = (&Project{}).GetEngine().
		Where("id = ?", p.Id).
		Cols("published", "published_at").
		Update(p)
	return
}

func UpdateClosed(p Project) (err error) {
	_, err = (&Project{}).GetEngine().
		Where("id = ?", p.Id).
		Cols("closed", "closed_at").
		Update(p)
	return
}

func AddFavourite(uid string, pid string) (err error) {
	_, err = (&Favourite{}).GetEngine().Insert(Favourite{
		ProjectId: pid,
		UserId:    uid,
		CreateAt:  time.Time{},
	})
	return
}

func RemoveFavourite(uid string, pid string) (err error) {
	_, err = (&Favourite{}).GetEngine().Delete(Favourite{
		ProjectId: pid,
		UserId:    uid,
	})
	return
}

func GetSubjects() (subjects []string, err error) {
	err = (&ProjectSubject{}).GetEngine().
		Cols("subject").
		Distinct("subject").
		Find(&subjects)
	return
}
func GetSkills() (skills []string, err error) {
	err = (&ProjectSkill{}).GetEngine().
		Cols("skill").
		Distinct("skill").
		Find(&skills)
	return
}

func ViewProject(pid string) (err error) {
	_, err = adapter.Engine.
		Exec("update project set read_num = read_num + 1 where id = ?", pid)
	return
}

func CloneProject(uid string, pid string) (err error) {
	var project Project
	_, err = (&Project{}).GetEngine().ID(pid).Get(&project)
	project.TeacherId = uid
	project.Id = util.NewId()
	project.Closed = false
	project.Published = false
	project.CreateAt = time.Now()
	project.ReadNum = 0
	project.JoinNum = 0
	project.ProjectTitle = project.ProjectTitle + "-副本"
	_, err = (&Project{}).GetEngine().Insert(&project)
	newPid := project.Id
	err = CloneProjectChapters(pid, newPid)
	return
}

func IsEditableProject(pid string) (e bool) {
	var p Project
	var b bool
	b, _ = (&Project{}).GetEngine().ID(pid).Get(&p)
	if b {
		e = p.Closed == false
	} else {
		e = false
	}
	return
}