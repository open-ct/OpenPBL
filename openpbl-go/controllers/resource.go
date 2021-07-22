package controllers

import (
	"github.com/astaxie/beego"
	"openpbl-go/models"
)

// ResourceController
// Operations about Projects
type ResourceController struct {
	beego.Controller
}

// GetResource
// @Title
// @Description
// @Param id path string true ""
// @Success 200 {object} models.Resource
// @router /:id [get]
func (p *ResourceController) GetResource() {
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
// @router / [post]
func (p *ResourceController) CreateResource() {
	s, err := p.GetInt64("SectionId")
	h, err := p.GetBool("hasHomeWork")
	resource := &models.Resource{
		SectionId:         s,
		FileTitle:         p.GetString("fileTitle"),
		FileIntroduce:     p.GetString("fileIntroduce"),
		FilePath:          p.GetString("FilePath"),
		Content:           p.GetString("Content"),
		HasHomeWork:       h,
		HomeWorkTitle:     p.GetString("HomeWorkTitle"),
		HomeWorkIntroduce: p.GetString("HomeWorkIntroduce"),
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
// @router /:id [post]
func (p *ResourceController) UpdateResource() {
	h, err := p.GetBool("hasHomeWork")
	s, err := p.GetInt64("sectionId")
	id, err := p.GetInt64(":id")
	resource := &models.Resource{
		Id:                id,
		SectionId:         s,
		FileTitle:         p.GetString("fileTitle"),
		FileIntroduce:     p.GetString("fileIntroduce"),
		FilePath:          p.GetString("filePath"),
		Content:           p.GetString("content"),
		HasHomeWork:       h,
		HomeWorkTitle:     p.GetString("homeWorkTitle"),
		HomeWorkIntroduce: p.GetString("homeWorkIntroduce"),
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
// @router /content/:id [post]
func (p *ResourceController) UpdateResourceContent() {
	id, err := p.GetInt64(":id")
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