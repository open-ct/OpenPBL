package controllers

import (
	"openpbl-go/models"
	"strconv"
)

// GetSection
// @Title
// @Description
// @Param body body models.File true ""
// @Success 200 {object}
// @Failure 403 body is empty
// @router /chapter/section/:sid [get]
func (p *ProjectController) GetSection() {
	sid := p.GetString(":sid")
	if sid != "" {
		section, err := models.GetSectionById(sid)
		if err != nil {
			p.Data["json"] = map[string]string{"error": err.Error()}
		}
		p.Data["json"] = map[string]models.Section{"section": section}
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

// GetSubmitFiles
// @Title
// @Description
// @Param pid path string true ""
// @Param sid path string true ""
// @Success 200 {object} []models.SubmitFile
// @Failure 403 body is empty
// @router /submit-files/:pid/:sid [get]
func (p *ProjectController) GetSubmitFiles() {
	sid := p.GetString(":sid")
	pid := p.GetString(":pid")
	if sid != "" && pid != "" {
		files, err := models.GetSubmitFiles(sid, pid)
		if err != nil {
			p.Data["json"] = map[string]string{"error": err.Error()}
		}
		p.Data["json"] = map[string][]models.SubmitFile{"files": files}
	}
	p.ServeJSON()
}

// CreateSubmitFile
// @Title
// @Description
// @Param body body models.SubmitFile true ""
// @Success 200 {object}
// @Failure 403 body is empty
// @router /submit-files [post]
func (p *ProjectController) CreateSubmitFile() {
	pid, err := p.GetInt64("projectId")
	sid, err := p.GetInt64("studentId")
	f := &models.SubmitFile{
		ProjectId:       pid,
		StudentId:       sid,
		SubmitIntroduce: p.GetString("submitIntroduce"),
		FilePath:        p.GetString("filePath"),
		FileName:        p.GetString("fileName"),
	}
	err = f.Create()
	if err != nil {
		p.Data["json"] = map[string]string{"error": err.Error()}
	}
	p.Data["json"] = map[string]string{"id": strconv.FormatInt(f.Id, 10)}
	p.ServeJSON()
}


