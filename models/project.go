package models

import (
	"errors"
	"time"
	"xorm.io/xorm"
)


type Project struct {
	Id                 int64     `json:"id" xorm:"not null pk autoincr"`
	Image              string    `json:"image" xorm:"longtext"`
	ProjectTitle       string    `json:"projectTitle"`
	ProjectIntroduce   string    `json:"projectIntroduce"`
	ProjectGoal        string    `json:"projectGoal"`

	TeacherId          string    `json:"teacherId" xorm:"index"`

	Subjects           string    `json:"subjects" xorm:"default ''"`
	Skills             string    `json:"skills" xorm:"default ''"`

	Closed             bool      `json:"closed" xorm:"default false index"`
	CreateAt           time.Time `json:"createAt" xorm:"created"`
	PublishedAt        time.Time `json:"publishedAt"`
	ClosedAt           time.Time `json:"closedAt"`
	ReadNum            int64     `json:"readNum" xorm:"default 0"`
	JoinNum            int64     `json:"joinNum" xorm:"default 0"`

	Published          bool      `json:"published" xorm:"default false index"`

	LearnMinute        int       `json:"learnMinute" xorm:"default 10"`
	LearnMinuteWeight  int       `json:"learnMinuteWeight" xorm:"default 100"`
}

type ProjectDetail struct {
	Project            `xorm:"extends"`
	Learning bool      `json:"learning"`
}

type ProjectSkill struct {
	Skill          string     `json:"skill" xorm:"not null pk"`
	ProjectId      int64      `json:"projectId" xorm:"not null pk"`
}

type ProjectSubject struct {
	Subject        string     `json:"subject" xorm:"not null pk"`
	ProjectId      int64      `json:"projectId" xorm:"not null pk"`
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

func GetProjectByPidForTeacher(pid string) (pd ProjectDetail, err error) {
	var p Project
	c, err := (&Project{}).GetEngine().
		ID(pid).
		Get(&p)
	pd = ProjectDetail{
		Project:  p,
		Learning: false,
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
	_, err = p.GetEngine().ID(p.Id).Update(p)
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
	_, err = p.GetEngine().ID(p.Id).Delete(p)
	return
}

func GetOutlineByPid(pid string) (c []Outline, err error) {
	err = adapter.Engine.
		SQL("select * from chapter left join section s on chapter.id = s.chapter_id where chapter.project_id = 1").
		// Where("project_id = ?", pid).
		// Asc("chapter_number").
		Find(&c)
	return
}

func UpdatePublished(p Project) (err error) {
	_, err = (&Project{}).GetEngine().
		Where("id = ?", p.Id).
		Cols("published", "publish_at").
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