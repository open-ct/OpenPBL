package controllers

import (
	"OpenPBL/models"
)


// GetResource
// @Title
// @Description
// @Param id path string true ""
// @Success 200 {object} models.Resource
// @router /:projectId/resource/:resourceId [get]
func (p *ProjectController) GetResource() {
	var err error
	rid := p.GetString(":id")
	resource, err := models.GetResourceById(rid)
	if err != nil {
		p.Data["json"] = map[string]models.Resource{"resource": {}}
	} else {
		p.Data["json"] = map[string]models.Resource{"resource": resource}
	}
	p.ServeJSON()
}

// CreateResource
// @Title
// @Description
// @Param body body models.Resource true ""
// @Success 200 {object} models.TeacherProject
// @Failure 400
// @router /:projectId/resource [post]
func (p *ProjectController) CreateResource() {
	s, err := p.GetInt64("sectionId")
	resource := &models.Resource{
		SectionId:         s,
		Content:           p.GetString("Content"),
	}
	err = resource.Create()
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
			Msg:  "创建成功",
			Data: resource.Id,
		}
	}
	p.ServeJSON()
}

// UpdateResource
// @Title
// @Description
// @Param body body models.Resource true ""
// @Success 200 {object} models.TeacherProject
// @Failure 400
// @router /:projectId/resource/:resourceId [post]
func (p *ProjectController) UpdateResource() {
	s, err := p.GetInt64("sectionId")
	id, err := p.GetInt64(":resourceId")
	resource := &models.Resource{
		Id:                id,
		SectionId:         s,
		Content:           p.GetString("content"),
	}
	err = resource.Update()
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
			Msg:  "更新成功",
		}
	}
	p.ServeJSON()
}

// UpdateResourceContent
// @Title
// @Description
// @Param id path string true ""
// @Param content body string true ""
// @Success 200 {object} models.TeacherProject
// @Failure 400
// @router /:projectId/resource/:resourceId/content [post]
func (p *ProjectController) UpdateResourceContent() {
	id, err := p.GetInt64(":resourceId")
	resource := &models.Resource{
		Id:      id,
		Content: p.GetString("content"),
	}
	err = resource.Update()
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
			Msg:  "保存成功",
		}
	}
	p.ServeJSON()
}
