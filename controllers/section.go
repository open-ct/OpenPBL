package controllers

import (
	"OpenPBL/models"
	"OpenPBL/util"
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

// GetSectionFiles
// @Title
// @Description
// @Param sectionId path string true ""
// @Param projectId path string true ""
// @Success 200 {object}
// @Failure 403 body is empty
// @router /:projectId/section/:sectionId/files [get]
func (p *ProjectController) GetSectionFiles() {
	sid := p.GetString(":sectionId")
	files, err := models.GetSectionFiles(sid)
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg: err.Error(),
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
			Data: files,
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
		SectionMinute:    1,
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

// DeleteChapterSection
// @Title
// @Description
// @Param sid path string true ""
// @Success 200 {object}
// @Failure 401
// @router /:projectId/chapter/:chapterId/section/:sectionId/delete [post]
func (p *ProjectController) DeleteChapterSection() {
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
	err = section.Delete()
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
			Msg: err.Error(),
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

// UploadSectionFile
// @Title
// @Description
// @Param filePath body string true ""
// @Success 200 {object} models.TeacherProject
// @Failure 400
// @router /:projectId/section/:sectionId/file [post]
func (p *ProjectController) UploadSectionFile() {
	user := p.GetSessionUser()
	if !util.IsTeacher(user) {
		p.Data["json"] = Response{
			Code: 403,
			Msg:  "非法用户",
		}
		p.ServeJSON()
		return
	}
	sid, err := p.GetInt64(":sectionId")
	name := p.GetString("name")
	url := p.GetString("url")
	filePath := p.GetString("filePath")
	r := &models.SectionFile{
		SectionId:  sid,
		Name:       name,
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
			Data: r,
		}
	}
	p.ServeJSON()
}

// UpdateSectionFile
// @Title
// @Description
// @Param filePath body string true ""
// @Success 200 {object} models.TeacherProject
// @Failure 400
// @router /:projectId/section/:sectionId/file/:fileId/update [post]
func (p *ProjectController) UpdateSectionFile() {
	user := p.GetSessionUser()
	if !util.IsTeacher(user) {
		p.Data["json"] = Response{
			Code: 403,
			Msg:  "非法用户",
		}
		p.ServeJSON()
		return
	}
	sid, err := p.GetInt64(":sectionId")
	fileId, err := p.GetInt64(":fileId")
	name := p.GetString("name")
	url := p.GetString("url")
	filePath := p.GetString("filePath")
	r := &models.SectionFile{
		Id:         fileId,
		SectionId:  sid,
		Name:       name,
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
			Data: r,
		}
	}
	p.ServeJSON()
}

// DeleteSectionFile
// @Title
// @Description
// @Param filePath body string true ""
// @Success 200 {object} models.TeacherProject
// @Failure 400
// @router /:projectId/section/:sectionId/file/:fileId/delete [post]
func (p *ProjectController) DeleteSectionFile() {
	user := p.GetSessionUser()
	if !util.IsTeacher(user) {
		p.Data["json"] = Response{
			Code: 403,
			Msg:  "非法用户",
		}
		p.ServeJSON()
		return
	}
	fid := p.GetString(":fileId")
	err := models.DeleteSectionFile(fid)
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