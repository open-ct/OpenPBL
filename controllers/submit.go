package controllers

import (
	"OpenPBL/models"
	"OpenPBL/util"
	"encoding/json"
	"strconv"
	"time"
)

// CreateSubmit
// @Title
// @Description
// @Param body body models.Submit true ""
// @Success 200 {object}
// @Failure 403 body is empty
// @router /:projectId/task/:taskId/submit [post]
func (p *ProjectController) CreateSubmit() {
	var resp Response
	user := p.GetSessionUser()
	if !util.IsStudent(user) {
		resp = Response{
			Code: 403,
			Msg:  "非法用户",
		}
		p.Data["json"] = resp
		p.ServeJSON()
		return
	}
	uid := util.GetUserId(user)
	tid, err := p.GetInt64(":taskId")
	pid, err := p.GetInt64(":projectId")

	submit := &models.Submit{
		ProjectId:       pid,
		StudentId:       uid,
		TaskId:          tid,
		SubmitType:      p.GetString("submitType"),
		SubmitTitle:     p.GetString("submitTitle"),
		SubmitIntroduce: p.GetString("submitIntroduce"),
		SubmitContent:   p.GetString("submitContent"),
		CreateAt:        time.Now(),
		Scored:          true,
		Score:           100,
	}
	if submit.SubmitType == "file" {
		submit.Scored = false
		submit.Score = 0
	}
	var c = make([]models.Choice, 0)
	if submit.SubmitType == "survey" {
		err = json.Unmarshal([]byte(p.GetString("choices")), &c)
	}
	err = submit.Create(c)
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
			Msg:  "提交成功",
			Data: strconv.FormatInt(submit.Id, 10),
		}
	}
	p.ServeJSON()
}

// UploadSubmitFile
// @Title
// @Description
// @Param filePath body string true ""
// @Success 200 {object} models.TeacherProject
// @Failure 400
// @router /:projectId/task/:taskId/submit/:submitId/file [post]
func (p *ProjectController) UploadSubmitFile() {
	user := p.GetSessionUser()
	if !util.IsStudent(user) {
		p.Data["json"] = Response{
			Code: 403,
			Msg:  "非法用户",
		}
		p.ServeJSON()
		return
	}
	uid := util.GetUserId(user)
	tid, err := p.GetInt64(":taskId")
	pid, err := p.GetInt64(":projectId")
	url := p.GetString("url")
	fileName := p.GetString("name")
	filePath := p.GetString("filePath")
	sid, err := p.GetInt64(":submitId")
	if sid == 0 {
		submit := &models.Submit{
			ProjectId:       pid,
			StudentId:       uid,
			TaskId:          tid,
			SubmitType:      "file",
			CreateAt:        time.Now(),
			Scored:          false,
			Score:           0,
		}
		err = submit.Create(make([]models.Choice, 0))
		sid = submit.Id
	}
	r := &models.SubmitFile{
		SubmitId:   sid,
		Name:       fileName,
		FilePath:   filePath,
		Url:        url,
	}
	err = r.Create()
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
			Msg:  "上传成功",
		}
	}
	p.ServeJSON()
}

// UpdateSubmitFile
// @Title
// @Description
// @Param filePath body string true ""
// @Success 200 {object} models.TeacherProject
// @Failure 400
// @router /:projectId/task/:taskId/submit/:submitId/file/:fileId/update [post]
func (p *ProjectController) UpdateSubmitFile() {
	user := p.GetSessionUser()
	if !util.IsStudent(user) {
		p.Data["json"] = Response{
			Code: 403,
			Msg:  "非法用户",
		}
		p.ServeJSON()
		return
	}
	url := p.GetString("url")
	fileName := p.GetString("name")
	sid, err := p.GetInt64(":submitId")
	fid, err := p.GetInt64(":fileId")
	filePath := p.GetString("filePath")
	r := &models.SubmitFile{
		Id:         fid,
		SubmitId:   sid,
		Name:       fileName,
		FilePath:   filePath,
		Url:        url,
	}
	err = r.Update()
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
			Msg:  "上传成功",
		}
	}
	p.ServeJSON()
}

// DeleteSubmitFile
// @Title
// @Description
// @Param filePath body string true ""
// @Success 200 {object} models.TeacherProject
// @Failure 400
// @router /:projectId/task/:taskId/submit/:submitId/file/:fileId/delete [post]
func (p *ProjectController) DeleteSubmitFile() {
	user := p.GetSessionUser()
	if !util.IsStudent(user) {
		p.Data["json"] = Response{
			Code: 403,
			Msg:  "非法用户",
		}
		p.ServeJSON()
		return
	}
	fid := p.GetString(":fileId")
	err := models.DeleteSubmitFile(fid)
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

// UpdateSubmit
// @Title
// @Description
// @Param body body models.Submit true ""
// @Success 200 {object}
// @Failure 403 body is empty
// @router /:projectId/task/:taskId/submit/:submitId [post]
func (p *ProjectController) UpdateSubmit() {
	user := p.GetSessionUser()
	var uid string
	if util.IsStudent(user) {
		uid = util.GetUserId(user)
	}
	tid, err := p.GetInt64(":taskId")
	sid, err := p.GetInt64(":submitId")
	pid, err := p.GetInt64(":projectId")
	score, err := p.GetInt("score")
	scored, err := p.GetBool("scored")
	f := &models.Submit{
		Id:              sid,
		ProjectId:       pid,
		StudentId:       uid,
		TaskId:          tid,
		SubmitType:      p.GetString("submitType"),
		SubmitTitle:     p.GetString("submitTitle"),
		SubmitIntroduce: p.GetString("submitIntroduce"),
		SubmitContent:   p.GetString("submitContent"),
		CreateAt:        time.Now(),
		Score:           score,
		Scored:          scored,
	}
	var c = make([]models.Choice, 0)
	if util.IsStudent(user) && f.SubmitType == "survey" {
		err = json.Unmarshal([]byte(p.GetString("choices")), &c)
	}
	err = f.Update(c)
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
			Msg:  "提交成功",
		}
	}
	p.ServeJSON()
}


// GetSubmitFiles
// @Title
// @Description
// @Param body body models.Submit true ""
// @Success 200 {object}
// @Failure 403 body is empty
// @router /:projectId/task/:taskId/submit/:submitId/files [get]
func (p *ProjectController) GetSubmitFiles() {
	sid := p.GetString(":submitId")
	files, err := models.GetSubmitFiles(sid)
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
			Msg:  "提交成功",
			Data: files,
		}
	}
	p.ServeJSON()
}