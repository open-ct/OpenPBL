package controllers

import (
	"OpenPBL/models"
	"strconv"
)

type TaskResponse struct {
	Response
	Tasks       []models.TaskDetail `json:"tasks"`
	Learning    bool                `json:"learning"`
}

// GetSectionTasks
// @Title
// @Description get all the tasks of a section
// @Param sid path string true ""
// @Success 200 {object}
// @Failure 400
// @router /tasks/:sid/:pid [get]
func (p *ProjectController) GetSectionTasks() {
	var resp TaskResponse
	sid := p.GetString(":sid")
	var learning bool
	user := p.GetSessionUser()
	if user == nil {
		resp = TaskResponse{
			Response: Response{
				Code: 401,
				Msg:  "请先登录",
			},
		}
		p.Data["json"] = resp
		p.ServeJSON()
		return
	}
	if user.Tag != "student" {
		learning = false
	}
	uid := user.Username
	pid, err := p.GetInt64(":pid")
	learning = models.IsLearningProject(pid, uid)
	tasks, err := models.GetSectionTasks(sid, uid, learning)
	if err != nil {
		p.Data["json"] = TaskResponse{
			Response: Response{
				Code: 400,
				Msg:  err.Error(),
			},
			Tasks:    nil,
			Learning: false,
		}
	} else {
		p.Data["json"] = TaskResponse{
			Response: Response{
				Code: 200,
			},
			Tasks:    tasks,
			Learning: learning,
		}
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
	task := &models.Task{
		SectionId:     sid,
		TaskOrder:     o,
		TaskTitle:     p.GetString("taskTitle"),
		TaskIntroduce: p.GetString("taskIntroduce"),
		TaskType:      p.GetString("taskType"),
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
	task := &models.Task{
		Id:            tid,
		SectionId:     sid,
		TaskOrder:     o,
		TaskTitle:     p.GetString("taskTitle"),
		TaskIntroduce: p.GetString("taskIntroduce"),
		TaskType:      p.GetString("taskType"),
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
	tid, err := p.GetInt64(":tid")
	task := &models.Task{
		Id:               tid,
	}
	err = task.Delete()
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

