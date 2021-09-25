package controllers

import (
	"OpenPBL/models"
	"github.com/astaxie/beego"
)


type HomeController struct {
	beego.Controller
}

type HomeProjectResponse struct {
	Code        int                  `json:"code"`
	Projects    []models.Project     `json:"projects"`
}

// GetLatestProjects
// @Title GetLatestProjects
// @Description Get latest projects
// @Param num params int true "Number of items"
// @Success 200 {object} HomeProjectResponse
// @router /projects/latest [get]
func (h *HomeController) GetLatestProjects() {
	num, err := h.GetInt("num")
	if err != nil {
		num = 3
	}
	projects, err := models.GetLatestProjects(num)
	if err != nil {
		h.Data["json"] = HomeProjectResponse{
			Code:    400,
		}
	} else {
		h.Data["json"] = HomeProjectResponse{
			Code:    200,
			Projects: projects,
		}
	}
	h.ServeJSON()
}
