// Copyright 2021 The OpenPBL Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controllers

import (
	"OpenPBL/models"
	"OpenPBL/util"
	"encoding/json"
	"time"
)

// CreateSubmit
// @Title CreateSubmit
// @Description
// @Param taskId path string true "The id of the task"
// @Param projectId path string true "The id of the project"
// @Param submitType body string true "The type of submit (survey, file, comment)"
// @Success 200 {object} Response
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
	tid := p.GetString(":taskId")
	pid := p.GetString(":projectId")

	submit := &models.Submit{
		Id:              util.NewId(),
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
		_ = json.Unmarshal([]byte(p.GetString("choices")), &c)
	}
	err := submit.Create(c)
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
			Msg:  "提交成功",
			Data: submit.Id,
		}
	}
	p.ServeJSON()
}

// UploadSubmitFile
// @Title UploadSubmitFile
// @Description
// @Param projectId path string true "The id of the project"
// @Param taskId path string true "The id of the task"
// @Param submitId path string true "The id of the submit"
// @Param filePath body string true "The file path"
// @Success 200 {object} Response
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
	url := p.GetString("url")
	fileName := p.GetString("name")
	filePath := p.GetString("filePath")
	sid := p.GetString(":submitId")
	if sid == "0" {
		submit := &models.Submit{
			Id:              util.NewId(),
			ProjectId:       p.GetString(":projectId"),
			StudentId:       util.GetUserId(user),
			TaskId:          p.GetString(":taskId"),
			SubmitType:      "file",
			CreateAt:        time.Now(),
			Scored:          false,
			Score:           0,
		}
		_ = submit.Create(make([]models.Choice, 0))
		sid = submit.Id
	}
	r := &models.SubmitFile{
		Id:         util.NewId(),
		SubmitId:   sid,
		Name:       fileName,
		FilePath:   filePath,
		Url:        url,
	}
	err := r.Create()
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
// @Title UpdateSubmitFile
// @Description
// @Param projectId path string true "The id of the project"
// @Param taskId path string true "The id of the task"
// @Param submitId path string true "The id of the submit"
// @Param fileId path string true "The id of the file"
// @Param filePath body string true "The path of the file"
// @Success 200 {object} Response
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
	filePath := p.GetString("filePath")
	r := &models.SubmitFile{
		Id:         p.GetString(":fileId"),
		SubmitId:   p.GetString(":submitId"),
		Name:       fileName,
		FilePath:   filePath,
		Url:        url,
	}
	err := r.Update()
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
// @Title DeleteSubmitFile
// @Description
// @Param projectId path string true "The id of the project"
// @Param taskId path string true "The id of the task"
// @Param submitId path string true "The id of the submit"
// @Param fileId path string true "The id of the file"
// @Param filePath body string true "The path of the file"
// @Success 200 {object} Response
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
// @Title UpdateSubmit
// @Description
// @Param projectId path string true "The id of the project"
// @Param taskId path string true "The id of the task"
// @Param submitId path string true "The id of the submit"
// @Success 200 {object} Response
// @router /:projectId/task/:taskId/submit/:submitId [post]
func (p *ProjectController) UpdateSubmit() {
	user := p.GetSessionUser()
	var uid string
	if util.IsStudent(user) {
		uid = util.GetUserId(user)
	}
	tid := p.GetString(":taskId")
	sid := p.GetString(":submitId")
	pid := p.GetString(":projectId")
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
// @Title GetSubmitFiles
// @Description
// @Param projectId path string true "The id of the project"
// @Param taskId path string true "The id of the task"
// @Param submitId path string true "The id of the submit"
// @Success 200 {object}
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