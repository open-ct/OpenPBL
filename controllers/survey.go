package controllers

import (
	"OpenPBL/models"
	"strconv"
)

// GetSurveyDetailByTaskId
// @Title
// @Description
// @Param sid path string true ""
// @Success 200 {object}
// @Failure 400
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
	tid, err := p.GetInt64(":taskId")
	survey := &models.Survey{
		TaskId:          tid,
		SurveyTitle:     p.GetString("surveyTitle"),
		SurveyIntroduce: p.GetString("surveyIntroduce"),
	}
	if err != nil {
		p.Data["json"] = map[string]string{"error": err.Error()}
	}
	err = survey.Create()
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  "创建失败",
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
			Msg:  "创建成功",
			Data: strconv.FormatInt(survey.Id, 10),
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
	sid, err := p.GetInt64(":sid")
	tid, err := p.GetInt64(":tid")
	survey := &models.Survey{
		Id:              sid,
		TaskId:          tid,
		SurveyTitle:     p.GetString("surveyTitle"),
		SurveyIntroduce: p.GetString("surveyIntroduce"),
	}
	if err != nil {
		p.Data["json"] = map[string]string{"error": err.Error()}
	}
	err = survey.Create()
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  "更新失败",
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
// @Title
// @Description
// @Param body body models.Question true ""
// @Success 200 {object} Response
// @Failure 400
// @router /:projectId/task/:taskId/survey/:surveyId/question [post]
func (p *ProjectController) CreateQuestion() {
	sid, err := p.GetInt64(":surveyId")
	o, err := p.GetInt("questionOrder")
	question := &models.Question{
		SurveyId:        sid,
		QuestionOrder:   o,
		QuestionType:    p.GetString("questionType"),
		QuestionTitle:   p.GetString("questionTitle"),
		QuestionOptions: p.GetString("questionOptions"),
	}
	if err != nil {
		p.Data["json"] = map[string]string{"error": err.Error()}
	}
	err = question.Create()
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  "创建失败",
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
			Msg:  "创建成功",
			Data: strconv.FormatInt(question.Id, 10),
		}
	}
	p.ServeJSON()
}
// UpdateQuestion
// @Title
// @Description
// @Param body body models.Survey true ""
// @Success 200 {object} Response
// @Failure 400
// @router /:projectId/task/:taskId/survey/:surveyId/question/:questionId [post]
func (p *ProjectController) UpdateQuestion() {
	qid, err := p.GetInt64(":questionId")
	sid, err := p.GetInt64(":surveyId")
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
			Msg:  "更新失败",
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
// @Title
// @Description
// @Param body body string true ""
// @Success 200 {object} Response
// @Failure 400
// @router /:projectId/task/:taskId/survey/:surveyId/questions/exchange [post]
func (p *ProjectController) ExchangeQuestion() {
	id1 := p.GetString("questionId1")
	id2 := p.GetString("questionId2")

	err := models.ExchangeQuestion(id1, id2)
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
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
// @Param qid path string true ""
// @Success 200 {object} Response
// @Failure 400
// @router /:projectId/task/:taskId/survey/:surveyId/question/:questionId/delete [post]
func (p *ProjectController) DeleteQuestion() {
	qid, err := p.GetInt64(":questionId")
	question := &models.Question{
		Id:              qid,
	}
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  "参数非法",
		}
		p.ServeJSON()
		return
	}
	err = question.Delete()
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  "删除失败",
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
			Msg:  "删除成功",
		}
	}
	p.ServeJSON()
}