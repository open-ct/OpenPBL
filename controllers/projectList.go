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
	"OpenPBL/util"
	"github.com/astaxie/beego"
	"github.com/casdoor/casdoor-go-sdk/auth"
	"strconv"
)


type ProjectListController struct {
	beego.Controller
}

type ProjectList struct {
	Code     int                    `json:"code"`
	Msg      string                 `json:"msg"`
	Count    int64                  `json:"count"`
	Projects []models.ProjectDetail `json:"projects"`
}

func (pl *ProjectListController) GetSessionUser() *auth.Claims {
	s := pl.GetSession("user")
	if s == nil {
		return nil
	}
	claims := &auth.Claims{}
	err := util.JsonToStruct(s.(string), claims)
	if err != nil {
		panic(err)
	}
	return claims
}

// GetUserProjectList
// @Title GetUserProjectList
// @Description Get project list
// @Param projectType path string true "For teacher: editing/published/finished, for student: learning/finished"
// @Param from query string false "Page number"
// @Param size query string false "Page size"
// @Param orderBy query string false "Order by , default UpdateAt"
// @Param orderType query string false "Order type desc/asc, default desc"
// @Param subject query string false "Project subjects"
// @Param skill query string false "Project skills"
// @Param text query string false "Search key-words"
// @Success 200 {object} ProjectList
// @router /:projectType [get]
func (pl *ProjectListController) GetUserProjectList() {
	orderBy := pl.GetString("orderBy")
	if orderBy == "" {
		orderBy = "update_at"
	}
	orderType := pl.GetString("orderType")
	if orderType == "" {
		orderType = "desc"
	}
	t := pl.GetString(":projectType")
	from, err := strconv.Atoi(pl.GetString("from"))
	if err != nil {
		from = 0
	}
	size, err := strconv.Atoi(pl.GetString("size"))
	if err != nil {
		size = 10
	}
	skill := pl.GetString("skill")
	subject := pl.GetString("subject")
	text := pl.GetString("text")

	user := pl.GetSessionUser()
	uid := util.GetUserId(user)
	var data ProjectList

	var projects []models.ProjectDetail
	var count    int64

	if util.IsStudent(user) {
		if t == "learning" {
			projects, count, err = models.GetMyProjectListBySid(uid, from, size, subject, skill, text, orderBy, orderType, false)
		} else if t == "finished" {
			projects, count, err = models.GetMyProjectListBySid(uid, from, size, subject, skill, text, orderBy, orderType, true)
		} else if t == "public" {
			projects, count, err = models.GetPublicProjectListForStudent(uid, from, size, subject, skill, text, orderBy, orderType, false)
		} else if t == "favourite" {
			projects, count, err = models.GetPublicProjectListForStudent(uid, from, size, subject, skill, text, orderBy, orderType, true)
		}
	} else if util.IsTeacher(user) {
		if t == "editing" {
			projects, count, err = models.GetMyProjectListByTid(uid, from, size, subject, skill, text, orderBy, orderType, false, false)
		} else if t == "published" {
			projects, count, err = models.GetMyProjectListByTid(uid, from, size, subject, skill, text, orderBy, orderType, true, false)
		} else if t == "finished" {
			projects, count, err = models.GetMyProjectListByTid(uid, from, size, subject, skill, text, orderBy, orderType, true, true)
		} else if t == "public" {
			projects, count, err = models.GetPublicProjectListForTeacher(uid, from, size, subject, skill, text, orderBy, orderType, false)
		} else if t == "favourite" {
			projects, count, err = models.GetPublicProjectListForTeacher(uid, from, size, subject, skill, text, orderBy, orderType, true)
		}
	}
	data = ProjectList{
		Code:     200,
		Count:    count,
		Projects: projects,
	}
	pl.Data["json"] = data
	pl.ServeJSON()
}