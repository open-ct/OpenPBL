package controllers

import (
	"openpbl-go/models"
	"strconv"
)

// CreateSubmit
// @Title
// @Description
// @Param body body models.Submit true ""
// @Success 200 {object}
// @Failure 403 body is empty
// @router /submit [post]
func (p *ProjectController) CreateSubmit() {
	tid, err := p.GetInt64("taskId")
	sid, err := p.GetInt64("studentId")
	f := &models.Submit{
		TaskId:          tid,
		StudentId:       sid,
		SubmitIntroduce: p.GetString("submitIntroduce"),
		FilePath:        p.GetString("filePath"),
	}
	err = f.Create()
	if err != nil {
		p.Data["json"] = map[string]string{"error": err.Error()}
	}
	p.Data["json"] = map[string]string{"id": strconv.FormatInt(f.Id, 10)}
	p.ServeJSON()
}


