package controllers

import (
	"OpenPBL/models"
	"OpenPBL/util"
)

type ChaptersResponse struct {
	Code       int              `json:"code"`
	Msg        string           `json:"msg"`
	Chapters   []models.Outline `json:"chapters"`
	ShowMinute bool             `json:"showMinute"`
}

// GetProjectChapters
// @Title GetProjectChapters
// @Description
// @Param projectId path string true "project id"
// @Success 200 {object} ChapterResponse
// @router /:projectId/chapters [get]
func (p *ProjectController) GetProjectChapters() {
	user := p.GetSessionUser()
	uid := ""
	show := false
	if util.IsStudent(user) {
		uid = user.Name
		show = true
	}
	if util.IsTeacher(user) {
		uid = p.GetString("studentId")
		if uid != "" {
			show = true
		}
	}
	pid := p.GetString(":projectId")
	outline, err := models.GetChaptersByPid(pid, uid)
	if err != nil {
		p.Data["json"] = ChaptersResponse{
			Code:     400,
			Msg:      err.Error(),
			Chapters: make([]models.Outline, 0),
		}
	} else {
		p.Data["json"] = ChaptersResponse{
			Code:     200,
			Chapters: outline,
			ShowMinute: show,
		}
	}
	p.ServeJSON()
}
// CreateProjectChapter
// @Title CreateProjectChapter
// @Description
// @Param projectId path string true "The id of the project"
// @Success 200 {object} Response
// @router /:projectId/chapter [post]
func (p *ProjectController) CreateProjectChapter() {
	pid := p.GetString(":projectId")
	num, err := p.GetInt("chapterNumber")
	chapter := &models.Chapter{
		Id:               util.NewId(),
		ProjectId:        pid,
		ChapterName:      p.GetString("chapterName"),
		ChapterNumber:    num,
	}
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	}
	err = chapter.Create()
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
			Msg:  "添加成功",
			Data: chapter.Id,
		}
	}
	p.ServeJSON()
}

// UpdateProjectChapter
// @Title UpdateProjectChapter
// @Description
// @Param projectId path string true "The id of the project"
// @Param chapterId path string true "The id of the chapter"
// @Success 200 {object} Response
// @router /:projectId/chapter/:chapterId [post]
func (p *ProjectController) UpdateProjectChapter() {
	cid := p.GetString(":chapterId")
	pid := p.GetString(":projectId")
	num, _ := p.GetInt("chapterNumber")
	chapter := &models.Chapter{
		Id:               cid,
		ProjectId:        pid,
		ChapterName:      p.GetString("chapterName"),
		ChapterNumber:    num,
	}
	err := chapter.Update()
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  "更新失败",
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

// DeleteProjectChapter
// @Title DeleteProjectChapter
// @Description
// @Param projectId path string true "The id of the project"
// @Param chapterId path string true "The id of the chapter"
// @Success 200 {object} Response
// @router /:projectId/chapter/:chapterId/delete [post]
func (p *ProjectController) DeleteProjectChapter() {
	cid := p.GetString(":chapterId")
	pid := p.GetString(":projectId")
	num, _ := p.GetInt("chapterNumber")
	chapter := &models.Chapter{
		Id:               cid,
		ProjectId:        pid,
		ChapterName:      p.GetString("chapterName"),
		ChapterNumber:    num,
	}
	err := chapter.Delete()
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

// ExchangeProjectChapter
// @Title ExchangeProjectChapter
// @Description
// @Param projectId path string true "The id of the project"
// @Success 200 {object} Response
// @router /:projectId/chapters/exchange [post]
func (p *ProjectController) ExchangeProjectChapter() {
	cid1 := p.GetString("chapterId1")
	cid2 := p.GetString("chapterId2")

	err := models.ExchangeChapters(cid1, cid2)
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
			Data: true,
		}
	}
	p.ServeJSON()
}
