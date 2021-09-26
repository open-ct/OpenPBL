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
	"time"
)

type CommentResponse struct {
	Code        int              `json:"code"`
	Comments    []models.Comment `json:"comments"`
	Count       int64            `json:"count"`
}

// GetProjectComments
// @Title
// @Description
// @Success 200 {object} Response
// @router /:projectId/comments [get]
func (p *ProjectController) GetProjectComments() {
	pid := p.GetString(":projectId")
	from, err := p.GetInt("from")
	if err != nil {
		from = 0
	}
	size, err := p.GetInt("size")
	if err != nil {
		size = 10
	}
	isTeacher, err := p.GetBool("isTeacher")
	if err != nil {
		isTeacher = false
	}
	text := p.GetString("text")
	comments, rows, err := models.GetProjectComments(pid, from, size, text, isTeacher)
	if err != nil {
		p.Data["json"] = CommentResponse{
			Code: 400,
		}
	} else {
		p.Data["json"] = CommentResponse{
			Code: 200,
			Comments: comments,
			Count: rows,
		}
	}
	p.ServeJSON()
}


// CreateProjectComment
// @Title
// @Description
// @Success 200 {object} Response
// @router /:projectId/comment [post]
func (p *ProjectController) CreateProjectComment() {
	pid := p.GetString(":projectId")
	user := p.GetSessionUser()
	uid := util.GetUserId(user)
	comment := &models.Comment{
		Id:        util.NewId(),
		ProjectId: pid,
		UserId:    uid,
		UserAvatar: user.Avatar,
		UserName:  user.DisplayName,
		IsTeacher: util.IsTeacher(user),
		Content:   p.GetString("content"),
		CreateAt:  time.Time{},
	}
	err := comment.Create()
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
			Msg: "发布成功",
		}
	}
	p.ServeJSON()
}

// DeleteProjectComment
// @Title
// @Description
// @Success 200 {object} Response
// @router /:projectId/comment/:commentId/delete [post]
func (p *ProjectController) DeleteProjectComment() {
	cid := p.GetString(":commentId")
	err := (&models.Comment{Id: cid}).Delete()
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
			Msg: "删除成功",
		}
	}
	p.ServeJSON()
}
