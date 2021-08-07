package controllers

import (
	"OpenPBL/models"
	"OpenPBL/util"
	"github.com/astaxie/beego"
	"github.com/casdoor/casdoor-go-sdk/auth"
	"strconv"
)

// ProjectListController
// Operations about Projects
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
// @Title
// @Description
// @Param projectType path string true "editing published closed // learning finished"
// @Param from query string false "from"
// @Param size query string false "size"
// @Param orderBy query string false "orderBy"
// @Param orderType query string false "orderType"
// @Param subject query string false ""
// @Param skill query string false ""
// @Param text query string false ""
// @Success 200 {object} ProjectList
// @Failure 401
// @router /:projectType [get]
func (pl *ProjectListController) GetUserProjectList() {
	orderBy := pl.GetString("orderBy")
	if orderBy == "" {
		orderBy = "create_at"
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
	var data ProjectList

	if user == nil {
		data = ProjectList{
			Code:     401,
			Msg:      "请先登录",
		}
		pl.Data["json"] = data
		pl.ServeJSON()
		return
	}
	uid := user.Username

	var projects []models.ProjectDetail
	var count    int64

	if user.Tag == "student" {
		if t == "learning" {
			projects, count, err = models.GetMyProjectListBySid(uid, from, size, subject, skill, text, orderBy, orderType, true)
		} else if t == "finished" {
			projects, count, err = models.GetMyProjectListBySid(uid, from, size, subject, skill, text, orderBy, orderType, false)
		} else if t == "public" {
			projects, count, err = models.GetPublicProjectListForStudent(uid, from, size, subject, skill, text, orderBy, orderType)
		}
	} else if user.Tag == "teacher" {
		if t == "editing" {
			projects, count, err = models.GetMyProjectListByTid(uid, from, size, subject, skill, text, orderBy, orderType, false, false)
		} else if t == "published" {
			projects, count, err = models.GetMyProjectListByTid(uid, from, size, subject, skill, text, orderBy, orderType, true, false)
		} else if t == "finished" {
			projects, count, err = models.GetMyProjectListByTid(uid, from, size, subject, skill, text, orderBy, orderType, true, true)
		} else if t == "public" {
			projects, count, err = models.GetPublicProjectListForTeacher(uid, from, size, subject, skill, text, orderBy, orderType)
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