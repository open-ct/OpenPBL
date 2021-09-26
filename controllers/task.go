package controllers

import (
	"OpenPBL/models"
	"OpenPBL/util"
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
// @Title GetSectionTasksDetail
// @Description Get section tasks with submit and survey
// @Param projectId path string true "The id of the project"
// @Param sectionId path string true "The id of the section"
// @Success 200 {object} TaskResponse
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
// @Title GetProjectTasks
// @Description get all the tasks of a section
// @Param projectId path string true "The id of the project"
// @Success 200 {object} Response
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
// @Title GetProjectTasksDetail
// @Description get all the tasks of a section
// @Param projectId path string true "The id of the project"
// @Success 200 {object} TaskResponse
// @router /:projectId/tasks-detail [get]
func (p *ProjectController) GetProjectTasksDetail() {
	var learning bool
	user := p.GetSessionUser()
	showSubmit := false
	teacherScore := false
	uid := util.GetUserId(user)
	editable := false
	showCount := false
	pid := p.GetString(":projectId")
	if util.IsTeacher(user) {
		uid = p.GetString("studentId")
		showSubmit = true
		teacherScore = true
		showCount = true
	}
	if util.IsStudent(user) {
		editable = models.IsEditableProject(pid)
		learning = models.IsLearningProject(pid, uid)
	} else {
		learning = false
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
// @Title CreateTask
// @Description
// @Param projectId path string true "The id of the project"
// @Success 200 {object} Response
// @router /:projectId/task [post]
func (p *ProjectController) CreateTask() {
	sid := p.GetString("sectionId")
	pid := p.GetString(":projectId")
	o, _ := p.GetInt("taskOrder")
	sn, _ := p.GetInt("sectionNumber")
	cn, _ := p.GetInt("chapterNumber")

	task := &models.Task{
		Id:            util.NewId(),
		SectionId:     sid,
		ProjectId:     pid,
		SectionNumber: sn,
		ChapterNumber: cn,
		TaskOrder:     o,
		TaskTitle:     p.GetString("taskTitle"),
		TaskIntroduce: p.GetString("taskIntroduce"),
		TaskType:      p.GetString("taskType"),
	}
	err := task.Create()
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
			Msg:  "创建成功",
			Data: task.Id,
		}
	}
	p.ServeJSON()
}

// UpdateTask
// @Title UpdateTask
// @Description
// @Param projectId path string true "The id of the project"
// @Success 200 {object} Response
// @router /:projectId/task/:taskId [post]
func (p *ProjectController) UpdateTask() {
	tid := p.GetString(":taskId")
	sid := p.GetString("sectionId")
	pid := p.GetString(":projectId")
	o, _ := p.GetInt("taskOrder")
	sn, _ := p.GetInt("sectionNumber")
	cn, _ := p.GetInt("chapterNumber")

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
	err := task.Update()
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
// @Title DeleteTask
// @Description
// @Param projectId path string true "The id of the project"
// @Param taskId path string true "The id of the task"
// @Success 200 {object} Response
// @router /:projectId/task/:taskId/delete [post]
func (p *ProjectController) DeleteTask() {
	tid := p.GetString(":taskId")
	task := &models.Task{
		Id:               tid,
	}
	err := task.Delete()
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
// @Title ExchangeTask
// @Description Exchange two tasks
// @Param projectId path string true "project id"
// @Param taskId1 body string true "task1 id"
// @Param taskId2 body string true "task2 id"
// @Success 200 {object} Response
// @router /:projectId/tasks/exchange [post]
func (p *ProjectController) ExchangeTask() {
	tid1 := p.GetString("taskId1")
	tid2 := p.GetString("taskId2")
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
