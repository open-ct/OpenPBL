package controllers

import (
	"OpenPBL/models"
	"encoding/json"
	"strconv"
)

type SectionResponse struct {
	Response
	Section   models.SectionDetail  `json:"section"`
}

// GetSectionDetail
// @Title
// @Description
// @Param sid path string true ""
// @Param pid path string true ""
// @Success 200 {object}
// @Failure 403 body is empty
// @router /:projectId/section/:sectionId [get]
func (p *ProjectController) GetSectionDetail() {
	sid := p.GetString(":sectionId")
	section, err := models.GetSectionDetailById(sid)
	if err != nil {
		p.Data["json"] = SectionResponse{
			Section:  models.SectionDetail{},
			Response: Response{
				Code: 400,
				Msg: err.Error(),
			},
		}
	} else {
		p.Data["json"] = SectionResponse{
			Section:  section,
			Response: Response{
				Code: 200,
			},
		}
	}
	p.ServeJSON()
}

// GetChapterSections
// @Title
// @Description
// @Param cid path string true "chapter id"
// @Success 200 {object} []models.Section
// @Failure 403 body is empty
// @router /:projectId/chapter/:chapterId/sections [get]
func (p *ProjectController) GetChapterSections() {
	cid := p.GetString(":chapterId")
	if cid != "" {
		sections, err := models.GetSectionsByCid(cid)
		if err != nil {
			p.Data["json"] = map[string][]models.Section{"sections": nil}
		} else {
			p.Data["json"] = map[string][]models.Section{"sections": sections}
		}
	}
	p.ServeJSON()
}

// CreateChapterSection
// @Title
// @Description
// @Param body body models.Section true ""
// @Success 200 {object}
// @Failure 403 body is empty
// @router /:projectId/chapter/:chapterId/section [post]
func (p *ProjectController) CreateChapterSection() {
	cid, err := p.GetInt64(":chapterId")
	sectionNumber, err := p.GetInt("sectionNumber")
	chapterNumber, err := p.GetInt("chapterNumber")
	section := &models.Section{
		ChapterId:        cid,
		SectionName:      p.GetString("sectionName"),
		SectionNumber:    sectionNumber,
		ChapterNumber:    chapterNumber,
	}
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	}
	err = section.Create()
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
			Msg:  "创建成功",
			Data: strconv.FormatInt(section.Id, 10),
		}
	}
	p.ServeJSON()
}

// UpdateChapterSection
// @Title
// @Description
// @Param body body models.Section true ""
// @Success 200 {object}
// @Failure 401
// @router /:projectId/chapter/:chapterId/section/:sectionId [post]
func (p *ProjectController) UpdateChapterSection() {
	sid, err := p.GetInt64(":sectionId")
	cid, err := p.GetInt64(":chapterId")
	sectionNumber, err := p.GetInt("sectionNumber")
	chapterNumber, err := p.GetInt("chapterNumber")
	section := &models.Section{
		Id:               sid,
		ChapterId:        cid,
		SectionName:      p.GetString("sectionName"),
		SectionNumber:    sectionNumber,
		ChapterNumber:    chapterNumber,
	}
	err = section.Update()
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

// DeleteChapterSection
// @Title
// @Description
// @Param sid path string true ""
// @Success 200 {object}
// @Failure 401
// @router /:projectId/chapter/:chapterId/section/:sectionId/delete [post]
func (p *ProjectController) DeleteChapterSection() {
	sid, err := p.GetInt64(":sectionId")
	section := &models.Section{
		Id:               sid,
	}
	err = section.Delete()
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

// ExchangeChapterSection
// @Title
// @Description
// @Param sid path string true ""
// @Success 200 {object}
// @Failure 401
// @router /:projectId/chapter/:chapterId/sections/exchange [post]
func (p *ProjectController) ExchangeChapterSection() {
	sid1 := p.GetString("sectionId1")
	sid2 := p.GetString("sectionId2")
	err := models.ExchangeSections(sid1, sid2)
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

// UpdateSectionsMinute
// @Title
// @Description
// @Param body body []models.Section true ""
// @Success 200 {object}
// @Failure 401
// @router /:projectId/sections-minute [post]
func (p *ProjectController) UpdateSectionsMinute() {
	sections := make([]models.Section, 0)
	err := json.Unmarshal([]byte(p.GetString("sections")), &sections)
	err = models.UpdateSectionsMinute(sections)
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  err.Error(),
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

