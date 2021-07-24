package models

import (
	"xorm.io/xorm"
)

type Survey struct {
	Id              int64     `json:"id" xorm:"not null pk autoincr"`
	TaskId          int64     `json:"taskId" xorm:"not null index"`

	SurveyTitle     string    `json:"surveyTitle"`
	SurveyIntroduce string    `json:"surveyIntroduce"`
}
type Question struct {
	Id              int64     `json:"id" xorm:"not null pk autoincr"`
	SurveyId        int64     `json:"surveyId" xorm:"not null index"`
	QuestionOrder   int       `json:"questionOrder" xorm:"not null index"`
	QuestionTitle   string    `json:"questionTitle"`
	QuestionType    string    `json:"questionType"`
	QuestionOptions string    `json:"questionOptions" xorm:"text"`
	QuestionCount   string    `json:"question" xorm:"text"`
}
type SurveyDetail struct {
	Survey               `xorm:"extends"`
	Questions []Question `json:"questions" xorm:"extends"`
}

type Answer struct {
	Id              int64     `json:"id" xorm:"not null pk autoincr"`
	SubmitId        int64     `json:"submitId" xorm:"not null index"`
	SurveyIntroduce string    `json:"surveyIntroduce"`
}
type Choice struct {
	Id            int64     `json:"id" xorm:"not null pk autoincr"`
	AnswerId      int64     `json:"answerId" xorm:"not null index"`
	ChoiceOrder   int       `json:"choiceOrder" xorm:"not null index"`
	Options       string    `json:"options"`
}

func (s *Survey) GetEngine() *xorm.Session {
	return adapter.Engine.Table(s)
}
func (s *Survey) Create() (err error) {
	_, err = s.GetEngine().Insert(s)
	return
}
func (s *Survey) Update() (err error) {
	_, err = s.GetEngine().ID(s.Id).Update(s)
	return
}
func (s *Survey) Delete() (err error) {
	_, err = s.GetEngine().ID(s.Id).Delete(s)
	return
}

func (q *Question) GetEngine() *xorm.Session {
	return adapter.Engine.Table(q)
}
func (q *Question) Create() (err error) {
	_, err = q.GetEngine().Insert(q)
	return
}
func (q *Question) Update() (err error) {
	_, err = q.GetEngine().ID(q.Id).Update(q)
	return
}
func (q *Question) Delete() (err error) {
	_, err = q.GetEngine().ID(q.Id).Delete(q)
	return
}

func GetSurveyByTaskId(tid string) (s Survey, qs []Question, err error) {
	_, err = (&Survey{}).GetEngine().
		Where("task_id = ?", tid).
		Get(&s)
	err = (&Question{}).GetEngine().
		Where("survey_id = ?", s.Id).
		Asc("question_order").
		Find(&qs)
	return
}

func ExchangeQuestion(id1 string, id2 string) (err error) {
	_, err = adapter.Engine.
		Exec("update question t1 join question t2 on (t1.id = ? and t2.id = ?) " +
			"set t1.question_order = t2.question_order, t2.question_order = t1.question_order", id1, id2)
	return
}