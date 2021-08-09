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
		FilePath:        p.GetString("filePath"),
		CreateAt:        time.Now(),
		Score:           score,
		Scored:          scored,
	}
	var c = make([]models.Choice, 0)
	if user.Tag == "student" && f.SubmitType == "survey" {
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
