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
// @router /task/:tid/submit [post]
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
	tid, err := p.GetInt64(":tid")

	f := &models.Submit{
		StudentId:       uid,
		TaskId:          tid,
		SubmitType:      p.GetString("submitType"),
		SubmitTitle:     p.GetString("submitTitle"),
		SubmitIntroduce: p.GetString("submitIntroduce"),
		SubmitContent:   p.GetString("submitContent"),
		FilePath:        p.GetString("filePath"),
	}
	var c []models.Choice
	err = json.Unmarshal([]byte(p.GetString("choices")), &c)
	err = f.Create(c)
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
			Msg:  "提交成功",
			Data: strconv.FormatInt(f.Id, 10),
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
// @router /task/:tid/submit/:sid [post]
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
	tid, err := p.GetInt64(":tid")
	sid, err := p.GetInt64(":sid")
	f := &models.Submit{
		Id:              sid,
		StudentId:       uid,
		TaskId:          tid,
		SubmitType:      p.GetString("submitType"),
		SubmitTitle:     p.GetString("submitTitle"),
		SubmitIntroduce: p.GetString("submitIntroduce"),
		SubmitContent:   p.GetString("submitContent"),
		FilePath:        p.GetString("filePath"),
		CreateAt:        time.Now(),
	}
	var c []models.Choice
	err = json.Unmarshal([]byte(p.GetString("choices")), &c)
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

