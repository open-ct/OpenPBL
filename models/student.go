package models

import (
	"time"
	"xorm.io/xorm"
)

type LearnProject struct {
	Avatar     string    `json:"avatar" xorm:"text"`
	Name       string    `json:"name"`
	StudentId  string    `json:"studentId" xorm:"not null index pk"`
	ProjectId  int64     `json:"projectId" xorm:"not null index pk"`
	Learning   bool      `json:"learning" xorm:"index default 0"`
	JoinTime   time.Time `json:"joinTime" xorm:"created"`
}

type LearnSection struct {
	StudentId     string    `json:"studentId" xorm:"not null index pk"`
	SectionId     int64     `json:"sectionId" xorm:"not null index pk"`

	SectionTime   time.Time `json:"sectionTime"`
	LearnTime     time.Time `json:"learnTime"`

	TaskNum       int       `json:"taskNum"`
	SubmittedTask int       `json:"submittedTask" xorm:"default 0"`
}

func (l *LearnProject) GetEngine() *xorm.Session {
	return adapter.Engine.Table(l)
}
func (l *LearnSection) GetEngine() *xorm.Session {
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

func IsLearningProject(pid int64, uid string) (e bool) {
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

func (l *LearnSection) Create() (err error) {
	_, err = (&LearnSection{}).GetEngine().Insert(l)
	return
}
func (l *LearnSection) Update() (err error) {
	_, err = (&LearnSection{}).GetEngine().
		Where("student_id = ?", l.StudentId).
		Where("section_id = ?", l.SectionId).
		Update(l)
	return
}