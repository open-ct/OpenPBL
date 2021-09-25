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
	pid, err := p.GetInt64(":projectId")
	user := p.GetSessionUser()
	uid := util.GetUserId(user)
	comment := &models.Comment{
		ProjectId: pid,
		UserId:    uid,
		UserAvatar: user.Avatar,
		UserName:  user.DisplayName,
		IsTeacher: util.IsTeacher(user),
		Content:   p.GetString("content"),
		CreateAt:  time.Time{},
	}
	err = comment.Create()
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
	cid, err := p.GetInt64(":commentId")
	err = (&models.Comment{Id: cid}).Delete()
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
