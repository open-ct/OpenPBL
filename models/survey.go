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
	QuestionCount   string    `json:"questionCount" xorm:"text"`
}
type SurveyDetail struct {
	Survey               `json:"survey" xorm:"extends"`
	Questions []Question `json:"questions" xorm:"extends"`
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

func DeleteSurvey(tid int64) (err error) {
	var survey Survey
	_, err = (&Survey{}).GetEngine().Where("task_id = ?", tid).Get(&survey)
	suid := survey.Id
	_, err = (&Survey{}).GetEngine().ID(suid).Delete(&Survey{})
	_, err = (&Question{}).GetEngine().Where("survey_id = ?", suid).Delete(&Question{})
	return
}

func CloneSurvey(tid int64, newTid int64) (err error) {
	var survey Survey
	_, err = (&Survey{}).GetEngine().Where("task_id = ?", tid).Get(&survey)
	suid := survey.Id
	survey.Id = 0
	survey.TaskId = newTid
	_, err = (&Survey{}).GetEngine().Insert(&survey)
	newSuid := survey.Id
	var questions []Question
	err = (&Question{}).GetEngine().Where("survey_id = ?", suid).Find(&questions)
	for i:=0; i< len(questions); i++ {
		q := questions[i]
		q.Id = 0
		q.SurveyId = newSuid
		q.QuestionCount = ""
		_, err = (&Question{}).GetEngine().Insert(&q)
	}
	return
}