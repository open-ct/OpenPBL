package controllers

import (
	"openpbl-go/models"
	"strconv"
)

// GetSectionTasks
// @Title
// @Description
// @Param sid path string true ""
// @Success 200 {object}
// @Failure 400
// @router /chapter/section/tasks/:sid [get]
func (p *ProjectController) GetSectionTasks() {
	sid := p.GetString(":sid")
	tasks, err := models.GetSectionTasks(sid)
	if err != nil {
		p.Data["json"] = map[string][]models.Task{"tasks": nil}
	} else {
		p.Data["json"] = map[string][]models.Task{"tasks": tasks}
	}
	p.ServeJSON()
}

// CreateTask
// @Title
// @Description
// @Param body body models.Task true ""
// @Success 200 {object} Response
// @Failure 400
// @router /task [post]
func (p *ProjectController) CreateTask() {
	sid, err := p.GetInt64("sectionId")
	o, err := p.GetInt("taskOrder")
	sur, err := p.GetInt64("surveyId")
	task := &models.Task{
		SectionId:     sid,
		TaskOrder:     o,
		TaskTitle:     p.GetString("taskTitle"),
		TaskIntroduce: p.GetString("taskIntroduce"),
		TaskType:      p.GetString("taskType"),
		SurveyId:      sur,
	}
	if err != nil {
		p.Data["json"] = map[string]string{"error": err.Error()}
	}
	err = task.Create()
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  "创建失败",
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
			Msg:  "创建成功",
			Data: strconv.FormatInt(task.Id, 10),
		}
	}
	p.ServeJSON()
}

// UpdateTask
// @Title
// @Description
// @Param body body models.Task true ""
// @Success 200 {object} Response
// @Failure 401
// @router /task/:tid [post]
func (p *ProjectController) UpdateTask() {
	tid, err := p.GetInt64(":tid")
	sid, err := p.GetInt64("sectionId")
	o, err := p.GetInt("taskOrder")
	sur, err := p.GetInt64("surveyId")
	task := &models.Task{
		Id:            tid,
		SectionId:     sid,
		TaskOrder:     o,
		TaskTitle:     p.GetString("taskTitle"),
		TaskIntroduce: p.GetString("taskIntroduce"),
		TaskType:      p.GetString("taskType"),
		SurveyId:      sur,
	}
	err = task.Update()
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  "更新失败",
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
			Msg: "更新成功",
			Data: true,
		}
	}
	p.ServeJSON()
}

// DeleteTask
// @Title
// @Description
// @Param cid path string true ""
// @Success 200 {object} Response
// @Failure 401
// @router /task/delete/:tid [post]
func (p *ProjectController) DeleteTask() {
	cid, err := p.GetInt64(":tid")
	chapter := &models.Chapter{
		Id:               cid,
	}
	err = chapter.Delete()
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  "删除失败",
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
			Msg: "删除成功",
			Data: true,
		}
	}
	p.ServeJSON()
}

// ExchangeTask
// @Title
// @Description
// @Param cid path string true ""
// @Success 200 {object} Response
// @Failure 401
// @router /task/exchange/:tid1/:tid2 [post]
func (p *ProjectController) ExchangeTask() {
	tid1 := p.GetString(":tid1")
	tid2 := p.GetString(":tid2")
	err := models.ExchangeTasks(tid1, tid2)
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

