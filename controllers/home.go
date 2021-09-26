// Copyright 2021 The OpenPBL Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
