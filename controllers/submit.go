package controllers

import (
	"OpenPBL/models"
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
	if user == nil {
		resp = Response{
			Code: 401,
			Msg:  "请先登录",
		}
		p.Data["json"] = resp
		p.ServeJSON()
		return
	}
	if user.Tag != "student" {
		resp = Response{
			Code: 403,
			Msg:  "非法用户",
		}
		p.Data["json"] = resp
		p.ServeJSON()
		return
	}
	uid := user.Username
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
		FilePath:        p.GetString("filePath"),
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

// UpdateSubmit
// @Title
// @Description
// @Param body body models.Submit true ""
// @Success 200 {object}
// @Failure 403 body is empty
// @router /:projectId/task/:taskId/submit/:submitId [post]
func (p *ProjectController) UpdateSubmit() {
	var resp Response
	user := p.GetSessionUser()
	if user == nil {
		resp = Response{
			Code: 401,
			Msg:  "请先登录",
		}
		p.Data["json"] = resp
		p.ServeJSON()
		return
	}
	if user.Tag != "student" {
		resp = Response{
			Code: 403,
			Msg:  "非法用户",
		}
		p.Data["json"] = resp
		p.ServeJSON()
		return
	}
	uid := user.Username
	tid, err := p.GetInt64(":taskId")
	sid, err := p.GetInt64(":submitId")
	pid, err := p.GetInt64(":projectId")
	f := &models.Submit{
		Id:              sid,
		ProjectId:       pid,
		StudentId:       uid,
		TaskId:          tid,
		SubmitType:      p.GetString("submitType"),
		SubmitTitle:     p.GetString("submitTitle"),
		SubmitIntroduce: p.GetString("submitIntroduce"),
		SubmitContent:   p.GetString("submitContent"),
		FilePath:        p.GetString("filePath"),
		CreateAt:        time.Now(),
	}
	var c =make([]models.Choice, 0)
	if f.SubmitType == "survey" {
		err = json.Unmarshal([]byte(p.GetString("choices")), &c)
	}
	err = f.Update(c)
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  "提交失败",
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
			Msg:  "提交成功",
		}
	}
	p.ServeJSON()
}

