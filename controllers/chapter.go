package controllers

import (
	"OpenPBL/models"
	"strconv"
)

// GetProjectChapters
// @Title
// @Description
// @Param pid path string true "project id"
// @Success 200 {object} []models.Outline
// @Failure 403 body is empty
// @router /:id/chapters [get]
func (p *ProjectController) GetProjectChapters() {
	pid := p.GetString(":id")
	outline, err := models.GetChaptersByPid(pid)
	if err != nil {
		p.Data["json"] = map[string][]models.Outline{"chapters": nil}
	} else {
		p.Data["json"] = map[string][]models.Outline{"chapters": outline}
	}
	p.ServeJSON()
}
// CreateProjectChapter
// @Title
// @Description
// @Param body body models.Chapter true ""
// @Success 200 {object} Response
// @Failure 401
// @router /:id/chapter [post]
func (p *ProjectController) CreateProjectChapter() {
	pid, err := p.GetInt64(":id")
	num, err := p.GetInt("chapterNumber")
	chapter := &models.Chapter{
		ProjectId:        pid,
		ChapterName:      p.GetString("chapterName"),
		ChapterNumber:    num,
	}
	if err != nil {
		p.Data["json"] = map[string]string{"error": err.Error()}
	}
	err = chapter.Create()
	if err != nil {
		p.Data["json"] = map[string]string{"error": err.Error()}
	}
	p.Data["json"] = map[string]string{"id": strconv.FormatInt(chapter.Id, 10)}
	p.ServeJSON()
}

// UpdateProjectChapter
// @Title
// @Description
// @Param body body models.Chapter true ""
// @Success 200 {object} Response
// @Failure 401
// @router /:projectId/chapter/:chapterId [post]
func (p *ProjectController) UpdateProjectChapter() {
	cid, err := p.GetInt64(":chapterId")
	pid, err := p.GetInt64(":projectId")
	num, err := p.GetInt("chapterNumber")
	chapter := &models.Chapter{
		Id:               cid,
		ProjectId:        pid,
		ChapterName:      p.GetString("chapterName"),
		ChapterNumber:    num,
	}
	err = chapter.Update()
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
// @Title
// @Description
// @Param cid path string true ""
// @Success 200 {object} Response
// @Failure 401
// @router /:projectId/chapter/:chapterId/delete [post]
func (p *ProjectController) DeleteProjectChapter() {
	cid, err := p.GetInt64(":chapterId")
	chapter := &models.Chapter{
		Id:               cid,
	}
	err = chapter.Delete()
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  "删除失败",
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
// @Title
// @Description
// @Param cid path string true ""
// @Success 200 {object} Response
// @Failure 401
// @router /:projectId/chapters/exchange [post]
func (p *ProjectController) ExchangeProjectChapter() {
	cid1 := p.GetString(":chapterId1")
	cid2 := p.GetString(":chapterId2")

	err := models.ExchangeChapters(cid1, cid2)
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
			Data: true,
		}
	}
	p.ServeJSON()
}
