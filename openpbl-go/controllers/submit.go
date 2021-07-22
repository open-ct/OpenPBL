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
	rid, err := p.GetInt64("resourceId")
	sid, err := p.GetInt64("studentId")
	f := &models.Submit{
		ResourceId:      rid,
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


// GetSubmits
// @Title
// @Description
// @Param pid path string true ""
// @Param sid path string true ""
// @Success 200 {object} []models.SubmitFile
// @Failure 403 body is empty
// @router /submit/:pid/:sid [get]
func (p *ProjectController) GetSubmits() {
	sid := p.GetString(":sid")
	pid := p.GetString(":pid")
	if sid != "" && pid != "" {
		files, err := models.GetSubmitFiles(sid, pid)
		if err != nil {
			p.Data["json"] = map[string]string{"error": err.Error()}
		}
		p.Data["json"] = map[string][]models.Submit{"files": files}
	}
	p.ServeJSON()
}

