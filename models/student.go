package models

import (
	"xorm.io/xorm"
)

type LearnProject struct {
	StudentId  string   `json:"studentId" xorm:"not null index pk"`
	ProjectId  int64    `json:"projectId" xorm:"not null index pk"`
	Learning   bool     `json:"learning" xorm:"index default 0"`
}

type StudentInfo struct {

}

func (l *LearnProject) GetEngine() *xorm.Session {
	return adapter.Engine.Table(l)
}
func (sp *ProjectDetail) GetEngine() *xorm.Session {
	return adapter.Engine.Table(sp)
}

func (l *LearnProject) Create() (err error) {
	_, err = (&LearnProject{}).GetEngine().Insert(l)
	return
}

func (l *LearnProject) Update() (err error) {
	_, err = (&LearnProject{}).GetEngine().
		Where("student_id = ?", l.StudentId).
		Where("project_id = ?", l.ProjectId).
		Update(&l)
	return
}



func GetProjectStudents(pid string, from int, size int) (s []StudentInfo, err error) {
	err = adapter.Engine.
		SQL("select * from student " +
			"inner join " +
			"( select * from learn_project where project_id = ?) as l " +
			"on student.id = l.student_id limit ?, ?", pid, from, size).
		Find(&s)
	return
}
func CountProjectStudents(pid string, from int, size int) (rows int64, err error) {
	_, err = adapter.Engine.
		SQL("select count(*) from student " +
			"inner join " +
			"( select * from learn_project where project_id = ?) as l " +
			"on student.id = l.student_id", pid).
		Get(&rows)
	return
}

