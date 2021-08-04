package models

import (
	"strconv"
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

	LearnMinute   int       `json:"learnMinute" xorm:"default 0"`
	LearnSecond   int       `json:"learnSecond" xorm:"default 0"`
}

type LastLearn struct {
	StudentId     string    `json:"studentId" xorm:"not null pk"`
	SectionId     int64     `json:"sectionId" xorm:"not null index"`
	ExitAt        time.Time `json:"exitAt" xorm:"updated"`
}

type LastLearnSection struct {
	LastLearn             `xorm:"extends"`
	Id            int64   `json:"id"`
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
	id, err := strconv.ParseInt(pid, 10, 64)
	e, err = (&LearnProject{}).GetEngine().Exist(&LearnProject{
		StudentId: uid,
		ProjectId: id,
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

func GetLearnSection(sectionId int64, studentId string) (l LearnSection, err error) {
	var b bool
	b, err = (&LearnSection{}).GetEngine().
		Where("section_id = ?", sectionId).
		Where("student_id = ?", studentId).
		Get(&l)
	if !b {
		l.SectionId = sectionId
		l.StudentId = studentId
		err = (&l).Create()
	}
	return
}

func GetLastLearnSection(studentId string) (l LastLearnSection, err error) {
	var b bool
	b, err = (&LastLearn{}).GetEngine().
		Where("student_id = ?", studentId).
		Join("LEFT OUTER", Section{}, "last_learn.section_id = section.id").
		Get(&l)
	l.Last = b
	return
}

func (l *LearnSection) Create() (err error) {
	_, err = (&LearnSection{}).GetEngine().Insert(l)
	_, err = (&LastLearn{}).GetEngine().Insert(&LastLearn{
		StudentId: l.StudentId,
		SectionId: l.SectionId,
		ExitAt:    time.Now(),
	})
	return
}
func (l *LearnSection) Update() (err error) {
	_, err = (&LearnSection{}).GetEngine().
		Where("student_id = ?", l.StudentId).
		Where("section_id = ?", l.SectionId).
		Update(l)
	_, err = (&LastLearn{}).GetEngine().
		Where("student_id = ?", l.StudentId).
		Update(&LastLearn{
			StudentId: l.StudentId,
			SectionId: l.SectionId,
			ExitAt:    time.Now(),
		})
	return
}