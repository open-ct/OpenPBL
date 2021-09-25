package controllers

import (
	"OpenPBL/models"
	"OpenPBL/util"
)


// GetResource
// @Title GetResource
// @Description
// @Param projectId path string true "The id of the project"
// @Param resourceId path string true "The id of the resource"
// @Success 200 {object} models.Resource
// @router /:projectId/resource/:resourceId [get]
func (p *ProjectController) GetResource() {
	var err error
	rid := p.GetString(":resourceId")
	resource, err := models.GetResourceById(rid)
	if err != nil {
		p.Data["json"] = map[string]models.Resource{"resource": {}}
	} else {
		p.Data["json"] = map[string]models.Resource{"resource": resource}
	}
	p.ServeJSON()
}

// CreateResource
// @Title CreateResource
// @Description
// @Param projectId path string true "The id of the project"
// @Param sectionId body string true "The id of the section"
// @Param content body string true "The content of the resource"
// @Success 200 {object} Response
// @router /:projectId/resource [post]
func (p *ProjectController) CreateResource() {
	s := p.GetString("sectionId")
	resource := &models.Resource{
		Id:                util.NewId(),
		SectionId:         s,
		Content:           p.GetString("content"),
	}
	err := resource.Create()
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
// @Title UpdateResource
// @Description
// @Param resourceId path string true "The id of the resource"
// @Param sectionId body string true "The id of the section"
// @Success 200 {object} Response
// @router /:projectId/resource/:resourceId [post]
func (p *ProjectController) UpdateResource() {
	id := p.GetString(":resourceId")
	resource := &models.Resource{
		Id:                id,
		SectionId:         p.GetString("sectionId"),
		Content:           p.GetString("content"),
	}
	err := resource.Update()
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
// @Title UpdateResourceContent
// @Description
// @Param projectId path string true "The id of the project"
// @Param resourceId path string true "The id of the resource"
// @Param content body string true ""
// @Success 200 {object} models.TeacherProject
// @router /:projectId/resource/:resourceId/content [post]
func (p *ProjectController) UpdateResourceContent() {
	resource := &models.Resource{
		Id:      p.GetString(":resourceId"),
		Content: p.GetString("content"),
	}
	err := resource.Update()
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
