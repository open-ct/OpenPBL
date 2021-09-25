package controllers

import (
	"OpenPBL/models"
	"OpenPBL/util"
	"encoding/json"
)


type SectionResponse struct {
	Response
	Section   models.SectionDetail  `json:"section"`
}

// GetSectionDetail
// @Title GetSectionDetail
// @Description
// @Param projectId path string true "The id of the project"
// @Param sectionId path string true "The id of the section"
// @Success 200 {object} SectionResponse
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
// @Title GetSectionFiles
// @Description
// @Param sectionId path string true "The id of the section"
// @Param projectId path string true "The id of the project"
// @Success 200 {object} Response
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
// @Title GetChapterSections
// @Description
// @Param projectId path string true "The id of the project"
// @Param chapterId path string true "The id of the chapter"
// @Success 200 {object} []models.Section
// @router /:projectId/chapter/:chapterId/sections [get]
func (p *ProjectController) GetChapterSections() {
	cid := p.GetString(":chapterId")
	sections, err := models.GetSectionsByCid(cid)
	if err != nil {
		p.Data["json"] = map[string][]models.Section{"sections": nil}
	} else {
		p.Data["json"] = map[string][]models.Section{"sections": sections}
	}
	p.ServeJSON()
}

// CreateChapterSection
// @Title
// @Description
// @Param projectId path string true "The id of the project"
// @Param chapterId path string true "The id of the chapter"
// @Success 200 {object}
// @Failure 403 body is empty
// @router /:projectId/chapter/:chapterId/section [post]
func (p *ProjectController) CreateChapterSection() {
	sectionNumber, err := p.GetInt("sectionNumber")
	chapterNumber, err := p.GetInt("chapterNumber")
	section := &models.Section{
		Id:               util.NewId(),
		ChapterId:        p.GetString(":chapterId"),
		SectionName:      p.GetString("sectionName"),
		SectionNumber:    sectionNumber,
		ChapterNumber:    chapterNumber,
		SectionMinute:    10,
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
			Data: section.Id,
		}
	}
	p.ServeJSON()
}

// UpdateChapterSection
// @Title UpdateChapterSection
// @Description
// @Param projectId path string true "The id of the project"
// @Param chapterId path string true "The id of the chapter"
// @Param sectionId path string true "The id of the section"
// @Success 200 {object}
// @Failure 401
// @router /:projectId/chapter/:chapterId/section/:sectionId [post]
func (p *ProjectController) UpdateChapterSection() {
	sectionNumber, _ := p.GetInt("sectionNumber")
	chapterNumber, _ := p.GetInt("chapterNumber")
	section := &models.Section{
		Id:               p.GetString(":sectionId"),
		ChapterId:        p.GetString(":chapterId"),
		SectionName:      p.GetString("sectionName"),
		SectionNumber:    sectionNumber,
		ChapterNumber:    chapterNumber,
	}
	err := section.Update()
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
// @Title DeleteChapterSection
// @Description
// @Param projectId path string true "The id of the project"
// @Param chapterId path string true "The id of the chapter"
// @Param sectionId path string true "The id of the section"
// @Success 200 {object} Response
// @router /:projectId/chapter/:chapterId/section/:sectionId/delete [post]
func (p *ProjectController) DeleteChapterSection() {
	sectionNumber, err := p.GetInt("sectionNumber")
	chapterNumber, err := p.GetInt("chapterNumber")
	section := &models.Section{
		Id:               p.GetString(":sectionId"),
		ChapterId:        p.GetString(":chapterId"),
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
// @Title ExchangeChapterSection
// @Description
// @Param projectId path string true "The id of the project"
// @Param chapterId path string true "The id of the chapter"
// @Success 200 {object} Response
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
// @Title UpdateSectionsMinute
// @Description
// @Param projectId path string true "The id of the project"
// @Param sections body []string true "The sections of the project"
// @Success 200 {object} Response
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
// @Title UploadSectionFile
// @Description
// @Param sectionId path string true "The id of the section"
// @Param name body string true "The name of the file"
// @Param filePath body string true "The path of the file"
// @Param url body string true "The url of the file"
// @Success 200 {object} Response
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
	r := &models.SectionFile{
		Id:         util.NewId(),
		SectionId:  p.GetString(":sectionId"),
		Name:       p.GetString("name"),
		FilePath:   p.GetString("filePath"),
		Url:        p.GetString("url"),
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
			Data: r,
		}
	}
	p.ServeJSON()
}

// UpdateSectionFile
// @Title UpdateSectionFile
// @Description
// @Param fileId body string true "The id of the file"
// @Param sectionId body string true "The section id"
// @Param name body string true "The name of the file"
// @Param filePath body string true "The path of the file"
// @Param url body string true "The url of the file"
// @Success 200 {object} models.TeacherProject
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
	r := &models.SectionFile{
		Id:         p.GetString(":fileId"),
		SectionId:  p.GetString(":sectionId"),
		Name:       p.GetString("name"),
		FilePath:   p.GetString("filePath"),
		Url:        p.GetString("url"),
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
			Data: r,
		}
	}
	p.ServeJSON()
}

// DeleteSectionFile
// @Title DeleteSectionFile
// @Description
// @Param projectId path string true "The id of the project"
// @Param :sectionId path string true "The id of the section"
// @Param :fileId path string true "The id of the file"
// @Success 200 {object} models.TeacherProject
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