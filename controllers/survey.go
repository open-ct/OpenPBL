package controllers

import (
	"OpenPBL/models"
	"OpenPBL/util"
)


// GetSurveyDetailByTaskId
// @Title GetSurveyDetailByTaskId
// @Description
// @Param projectId path string true "The id of the project"
// @Param taskId path string true "The id of the task"
// @Success 200 {object}
// @router /:projectId/task/:taskId/survey [get]
func (p *ProjectController) GetSurveyDetailByTaskId() {
	tid := p.GetString(":taskId")
	survey, qns, err := models.GetSurveyByTaskId(tid)
	if err != nil {
		p.Data["json"] = map[string]interface{}{"survey": nil, "questions": qns}
	} else {
		p.Data["json"] = map[string]interface{}{"survey": survey, "questions": qns}
	}
	p.ServeJSON()
}

// CreateSurvey unused
// @Title
// @Description
// @Param body body models.Survey true ""
// @Success 200 {object} Response
// @Failure 400
// @router /:projectId/task/:taskId/survey [post]
func (p *ProjectController) CreateSurvey() {
	tid := p.GetString(":taskId")
	survey := &models.Survey{
		Id:              util.NewId(),
		TaskId:          tid,
		SurveyTitle:     p.GetString("surveyTitle"),
		SurveyIntroduce: p.GetString("surveyIntroduce"),
	}
	err := survey.Create()
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
			Msg:  "创建成功",
			Data: survey.Id,
		}
	}
	p.ServeJSON()
}

// UpdateSurvey  unused
// @Title
// @Description
// @Param body body models.Survey true ""
// @Success 200 {object} Response
// @Failure 400
// @router /:projectId/task/:taskId/survey/:sid [post]
func (p *ProjectController) UpdateSurvey() {
	sid := p.GetString(":sid")
	tid := p.GetString(":tid")
	survey := &models.Survey{
		Id:              sid,
		TaskId:          tid,
		SurveyTitle:     p.GetString("surveyTitle"),
		SurveyIntroduce: p.GetString("surveyIntroduce"),
	}
	err := survey.Create()
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
			Msg:  "更新成功",
		}
	}
	p.ServeJSON()
}

// CreateQuestion
// @Title CreateQuestion
// @Description
// @Param projectId path string true "The id of the project"
// @Param taskId path string true "The id of the task"
// @Param surveyId path string true "The id of the survey"
// @Success 200 {object} Response
// @router /:projectId/task/:taskId/survey/:surveyId/question [post]
func (p *ProjectController) CreateQuestion() {
	sid := p.GetString(":surveyId")
	o, _ := p.GetInt("questionOrder")
	question := &models.Question{
		Id:              util.NewId(),
		SurveyId:        sid,
		QuestionOrder:   o,
		QuestionType:    p.GetString("questionType"),
		QuestionTitle:   p.GetString("questionTitle"),
		QuestionOptions: p.GetString("questionOptions"),
	}
	err := question.Create()
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
			Msg:  "创建成功",
			Data: question.Id,
		}
	}
	p.ServeJSON()
}
// UpdateQuestion
// @Title UpdateQuestion
// @Description
// @Param projectId path string true "The id of the project"
// @Param taskId path string true "The id of the task"
// @Param surveyId path string true "The id of the survey"
// @Param questionId path string true "The id of the question"
// @Success 200 {object} Response
// @router /:projectId/task/:taskId/survey/:surveyId/question/:questionId [post]
func (p *ProjectController) UpdateQuestion() {
	qid := p.GetString(":questionId")
	sid := p.GetString(":surveyId")
	o, err := p.GetInt("questionOrder")
	question := &models.Question{
		Id:              qid,
		SurveyId:        sid,
		QuestionOrder:   o,
		QuestionType:    p.GetString("questionType"),
		QuestionTitle:   p.GetString("questionTitle"),
		QuestionOptions: p.GetString("questionOptions"),
	}
	if err != nil {
		p.Data["json"] = map[string]string{"error": err.Error()}
	}
	err = question.Update()
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
			Msg:  "更新成功",
		}
	}
	p.ServeJSON()
}

// ExchangeQuestion
// @Title ExchangeQuestion
// @Description
// @Param projectId path string true "The id of the project"
// @Param taskId path string true "The id of the task"
// @Param surveyId path string true "The id of the survey"
// @Param questionId1 body string true "The id of the question1"
// @Param questionId2 body string true "The id of the question2"
// @Success 200 {object} Response
// @router /:projectId/task/:taskId/survey/:surveyId/questions/exchange [post]
func (p *ProjectController) ExchangeQuestion() {
	id1 := p.GetString("questionId1")
	id2 := p.GetString("questionId2")
	err := models.ExchangeQuestion(id1, id2)
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg: err.Error(),
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
			Data: true,
		}
	}
	p.ServeJSON()
}


// DeleteQuestion
// @Title
// @Description
// @Param projectId path string true "The id of the project"
// @Param taskId path string true "The id of the task"
// @Param surveyId path string true "The id of the survey"
// @Param questionId path string true "The id of the question"
// @Success 200 {object} Response
// @router /:projectId/task/:taskId/survey/:surveyId/question/:questionId/delete [post]
func (p *ProjectController) DeleteQuestion() {
	qid := p.GetString(":questionId")
	question := &models.Question{
		Id:              qid,
	}
	err := question.Delete()
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
			Msg:  "删除成功",
		}
	}
	p.ServeJSON()
}