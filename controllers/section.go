package controllers

import (
	"OpenPBL/models"
	"strconv"
)

type SectionResponse struct {
	Response
	Section   models.SectionDetail  `json:"section"`
	Editable  bool                  `json:"editable"`
}

// GetSectionDetail
// @Title
// @Description
// @Param sid path string true ""
// @Param pid path string true ""
// @Success 200 {object}
// @Failure 403 body is empty
// @router /chapter/section/:sid/:pid [get]
func (p *ProjectController) GetSectionDetail() {
	var resp SectionResponse
	var editable bool
	user := p.GetSessionUser()
	if user == nil {
		resp = SectionResponse{
			Response: Response{
				Code: 401,
				Msg:  "请先登录",
			},
		}
		p.Data["json"] = resp
		p.ServeJSON()
		return
	}
	if user.Tag != "student" {
		editable = false
	}
	uid := user.Username
	pid, err := p.GetInt64(":pid")

	editable = models.IsLearningProject(pid, uid)
	sid := p.GetString(":sid")
	section, err := models.GetSectionDetailById(sid)
	if err != nil {
		p.Data["json"] = SectionResponse{
			Section:  models.SectionDetail{},
			Editable: false,
		}
	} else {
		p.Data["json"] = SectionResponse{
			Section:  section,
			Editable: editable,
		}
	}
	p.ServeJSON()
}

type StudentList struct {
	Count    int64                `json:"count"`
	Students []models.StudentInfo `json:"students"`
}

// GetProjectStudents
// todo need refactor
// @Title
// @Description
// @Param from query int true ""
// @Param size query int true ""
// @Success 200 {object} []models.StudentInfo
// @Failure 403 body is empty
// @router /students/:pid [get]
func (p *ProjectController) GetProjectStudents() {
	pid := p.GetString(":pid")
	from, err := p.GetInt("from")
	if err != nil {
		from = 0
	}
	size, err := p.GetInt("size")
	if err != nil {
		size = 10
	}
	if pid != "" {
		students, err := models.GetProjectStudents(pid, from, size)
		rows, err := models.CountProjectStudents(pid, from, size)
		if err != nil {
			p.Data["json"] = map[string]string{"error": err.Error()}
		}
		p.Data["json"] = StudentList{
			Count: rows,
			Students: students,
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
// @router /chapter/sections/:cid [get]
func (p *ProjectController) GetChapterSections() {
	cid := p.GetString(":cid")
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
// @router /chapter/section [post]
func (p *ProjectController) CreateChapterSection() {
	cid, err := p.GetInt64("chapterId")
	num, err := p.GetInt("sectionNumber")
	section := &models.Section{
		ChapterId:        cid,
		SectionName:      p.GetString("sectionName"),
		SectionNumber:    num,
	}
	if err != nil {
		p.Data["json"] = map[string]string{"error": err.Error()}
	}
	err = section.Create()
	if err != nil {
		p.Data["json"] = map[string]string{"error": err.Error()}
	}
	p.Data["json"] = map[string]string{"id": strconv.FormatInt(section.Id, 10)}
	p.ServeJSON()
}

// UpdateChapterSection
// @Title
// @Description
// @Param body body models.Section true ""
// @Success 200 {object}
// @Failure 401
// @router /chapter/section/:sid [post]
func (p *ProjectController) UpdateChapterSection() {
	sid, err := p.GetInt64(":sid")
	cid, err := p.GetInt64("chapterId")
	num, err := p.GetInt("sectionNumber")
	section := &models.Section{
		Id:               sid,
		ChapterId:        cid,
		SectionName:      p.GetString("sectionName"),
		SectionNumber:    num,
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
// @router /chapter/section/delete/:sid [post]
func (p *ProjectController) DeleteChapterSection() {
	sid, err := p.GetInt64(":sid")
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
// @router /chapter/section/exchange/:sid1/:sid2 [post]
func (p *ProjectController) ExchangeChapterSection() {
	sid1 := p.GetString(":sid1")
	sid2 := p.GetString(":sid2")
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



