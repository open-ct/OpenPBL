package models

import (
	"xorm.io/xorm"
)

type Task struct {
	Id            int64     `json:"id" xorm:"not null pk autoincr"`
	SectionId     int64     `json:"sectionId" xorm:"not null index"`
	ProjectId     int64     `json:"projectId" xorm:"not null index"`

	SectionNumber int       `json:"sectionNumber" xorm:"index"`
	ChapterNumber int       `json:"chapterNumber" xorm:"index"`

	TaskOrder     int       `json:"taskOrder"`

	TaskTitle     string    `json:"taskTitle"`
	TaskIntroduce string    `json:"taskIntroduce" xorm:"text"`

	TaskType      string    `json:"taskType" xorm:"index"`
	TaskWeight    int       `json:"taskWeight"`
}

type TaskEvaluate struct {
	Task             `xorm:"extends"`
	Submitted bool   `json:"submitted"`
	Submit    Submit `json:"submit"`
}

type TaskDetail struct {
	Task         `xorm:"extends"`
	SurveyDetail `xorm:"extends"`
	SubmitDetail `xorm:"extends"`
}


func (t *Task) GetEngine() *xorm.Session {
	return adapter.Engine.Table(t)
}

func (t *Task) Create() (err error) {
	session := adapter.Engine.NewSession()
	defer session.Close()
	session.Begin()
	_, err = session.Insert(t)
	if t.TaskType == "survey" {
		_, err = session.Insert(Survey{
			TaskId:          t.Id,
		})
	}
	session.Commit()
	return
}
func (t *Task) Update() (err error) {
	_, err = t.GetEngine().ID(t.Id).Update(t)
	return
}
func (t *Task) Delete() (err error) {
	_, err = t.GetEngine().ID(t.Id).Delete(t)
	return
}

func GetSectionTasks(sid string, uid string, learning bool) (t []TaskDetail, err error) {
	err = (&Task{}).GetEngine().
		Where("section_id = ?", sid).
		Asc("task_order").
		Find(&t)
	var b bool
	for i := 0; i < len(t); i ++ {
		var s Survey
		var qs []Question
		var c []Choice
		if t[i].TaskType == "survey" {
			var m Submit
			_, err = (&Survey{}).GetEngine().
				Where("task_id = ?", t[i].Id).
				Get(&s)
			err = (&Question{}).GetEngine().
				Where("survey_id = ?", s.Id).
				Asc("question_order").
				Find(&qs)
			t[i].Survey = s
			t[i].Questions = qs

			if learning {
				b, err = (&Submit{}).GetEngine().
					Where("task_id = ?", t[i].Id).
					Where("student_id = ?", uid).
					Get(&m)
				if b {
					t[i].Submitted = true
				}
				t[i].Submit = m

				err = (&Choice{}).GetEngine().
					Where("submit_id = ?", m.Id).
					Asc("choice_order").
					Find(&c)
				t[i].Choices = c
			}
		} else {
			var m Submit
			if learning {
				b, err = (&Submit{}).GetEngine().
					Where("task_id = ?", t[i].Id).
					Where("student_id = ?", uid).
					Get(&m)
				t[i].Submit = m
				if b {
					t[i].Submitted = true
				}
			}
		}
	}
	return
}

func ExchangeTasks(cid1 string, cid2 string) (err error) {
	_, err = adapter.Engine.
		Exec("update task t1 join task t2 on (t1.id = ? and t2.id = ?) " +
			"set t1.task_order = t2.task_order, t2.task_order = t1.task_order", cid1, cid2)
	return
}


func GetProjectTasksDetail(sid string, uid string, learning bool) (t []TaskDetail, err error) {
	err = (&Task{}).GetEngine().
		Where("project_id = ?", sid).
		Asc("chapter_number").
		Asc("section_number").
		Asc("task_order").
		Find(&t)
	var b bool
	for i := 0; i < len(t); i ++ {
		var s Survey
		var qs []Question
		var c []Choice

		if t[i].TaskType == "survey" {
			var m Submit
			b, err = (&Survey{}).GetEngine().
				Where("task_id = ?", t[i].Id).
				Get(&s)
			if b {
				err = (&Question{}).GetEngine().
					Where("survey_id = ?", s.Id).
					Asc("question_order").
					Find(&qs)
				t[i].Questions = qs
			}
			t[i].Survey = s

			if learning {
				b, err = (&Submit{}).GetEngine().
					Where("task_id = ?", t[i].Id).
					Where("student_id = ?", uid).
					Get(&m)
				if b {
					t[i].Submitted = true
					err = (&Choice{}).GetEngine().
						Where("submit_id = ?", m.Id).
						Asc("choice_order").
						Find(&c)
					t[i].Choices = c
				}
				t[i].Submit = m
			}
		} else {
			var m Submit
			if learning {
				b, err = (&Submit{}).GetEngine().
					Where("task_id = ?", t[i].Id).
					Where("student_id = ?", uid).
					Get(&m)
				t[i].Submit = m
				if b {
					t[i].Submitted = true
				}
			}
		}
	}
	return
}

func GetProjectTasks(pid string) (t []TaskEvaluate, err error) {
	err = (&Task{}).GetEngine().
		Where("project_id = ?", pid).
		Asc("chapter_number").
		Asc("section_number").
		Asc("task_order").
		Find(&t)
	return
}