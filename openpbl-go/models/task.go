package models

import "xorm.io/xorm"

type Task struct {
	Id            int64     `json:"id" xorm:"not null pk autoincr"`
	SectionId     int64     `json:"sectionId" xorm:"not null index"`

	TaskOrder     int       `json:"taskOrder"`

	TaskTitle     string    `json:"taskTitle"`
	TaskIntroduce string    `json:"taskIntroduce" xorm:"text"`

	TaskType      string    `json:"taskType" xorm:"index"`
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

func GetSectionTasks(sid string) (t []Task, err error) {
	err = (&Task{}).GetEngine().
		Where("section_id = ?", sid).
		Asc("task_order").
		Find(&t)
	return
}

func ExchangeTasks(cid1 string, cid2 string) (err error) {
	_, err = adapter.Engine.
		Exec("update task t1 join task t2 on (t1.id = ? and t2.id = ?) " +
			"set t1.task_order = t2.task_order, t2.task_order = t1.task_order", cid1, cid2)
	return
}