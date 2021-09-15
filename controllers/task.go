package controllers

import (
	"OpenPBL/models"
	"OpenPBL/util"
	"strconv"
)

type TaskResponse struct {
	Response
	Tasks        []models.TaskDetail `json:"tasks"`
	Learning     bool                `json:"learning"`
	Editable     bool                `json:"editable"`
	TeacherScore bool                `json:"teacherScore"`
	ShowCount    bool                `json:"showCount"`
}

// GetSectionTasksDetail
// @Title
// @Description get section tasks with submit and survey
// @Param sid path string true ""
// @Success 200 {object} TaskResponse
// @Failure 400
// @router /:projectId/section/:sectionId/tasks [get]
func (p *ProjectController) GetSectionTasksDetail() {
	sid := p.GetString(":sectionId")
	user := p.GetSessionUser()
	uid := util.GetUserId(user)
	pid := p.GetString(":projectId")
	showCount := false
	learning := false
	editable := false
	if util.IsTeacher(user) {
		showCount = true
	} else {
		editable = models.IsEditableProject(pid)
		learning = models.IsLearningProject(pid, uid)
	}
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
			Editable: editable,
			ShowCount: showCount,
		}
	}
	p.ServeJSON()
}

// GetProjectTasks
// @Title
// @Description get all the tasks of a section
// @Param sid path string true ""
// @Success 200 {object}
// @Failure 400
// @router /:projectId/tasks [get]
func (p *ProjectController) GetProjectTasks() {
	pid := p.GetString(":projectId")
	tasks, err := models.GetProjectTasks(pid)
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
			Data: tasks,
		}
	}
	p.ServeJSON()
}

// GetProjectTasksDetail
// @Title
// @Description get all the tasks of a section
// @Param sid path string true ""
// @Success 200 {object}
// @Failure 400
// @router /:projectId/tasks-detail [get]
func (p *ProjectController) GetProjectTasksDetail() {
	var learning bool
	user := p.GetSessionUser()
	showSubmit := false
	teacherScore := false
	uid := util.GetUserId(user)
	editable := true
	showCount := false
	pid := p.GetString(":projectId")
	if util.IsTeacher(user) {
		uid = p.GetString("studentId")
		showSubmit = true
		editable = false
		teacherScore = true
		showCount = true
	}
	if !util.IsStudent(user) {
		learning = false
	} else {
		learning = models.IsLearningProject(pid, uid)
	}
	if learning {
		showSubmit = true
	}

	tasks, err := models.GetProjectTasksDetail(pid, uid, showSubmit)
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
			Editable: editable,
			TeacherScore: teacherScore,
			ShowCount: showCount,
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
// @router /:projectId/task [post]
func (p *ProjectController) CreateTask() {
	sid, err := p.GetInt64("sectionId")
	pid, err := p.GetInt64(":projectId")
	o, err := p.GetInt("taskOrder")
	sn, err := p.GetInt("sectionNumber")
	cn, err := p.GetInt("chapterNumber")

	task := &models.Task{
		SectionId:     sid,
		ProjectId:     pid,
		SectionNumber: sn,
		ChapterNumber: cn,
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
			Msg:  err.Error(),
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
// @router /:projectId/task/:taskId [post]
func (p *ProjectController) UpdateTask() {
	tid, err := p.GetInt64(":taskId")
	sid, err := p.GetInt64("sectionId")
	pid, err := p.GetInt64(":projectId")
	o, err := p.GetInt("taskOrder")
	sn, err := p.GetInt("sectionNumber")
	cn, err := p.GetInt("chapterNumber")

	task := &models.Task{
		Id:            tid,
		SectionId:     sid,
		ProjectId:     pid,
		SectionNumber: sn,
		ChapterNumber: cn,
		TaskOrder:     o,
		TaskTitle:     p.GetString("taskTitle"),
		TaskIntroduce: p.GetString("taskIntroduce"),
		TaskType:      p.GetString("taskType"),
	}
	err = task.Update()
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  err.Error(),
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
// @router /:projectId/task/:taskId/delete [post]
func (p *ProjectController) DeleteTask() {
	tid, err := p.GetInt64(":taskId")
	task := &models.Task{
		Id:               tid,
	}
	err = task.Delete()
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  err.Error(),
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
// @router /:projectId/tasks/exchange [post]
func (p *ProjectController) ExchangeTask() {
	tid1 := p.GetString(":taskId1")
	tid2 := p.GetString(":taskId2")
	err := models.ExchangeTasks(tid1, tid2)
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
