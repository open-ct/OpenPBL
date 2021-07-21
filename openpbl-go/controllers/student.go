package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/casdoor/casdoor-go-sdk/auth"
	"openpbl-go/models"
	"openpbl-go/util"
)

// StudentController
// Operations about Student
type StudentController struct {
	beego.Controller
}

func (pl *StudentController) GetSessionUser() *auth.Claims {
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


// GetStudent
// @Title GetStudent
// @Description
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Project
// @Failure 403 :id is empty
// @router /:sid [get]
func (u *StudentController) GetStudent() {
	var (
		sid   string
		err   error
		stu   models.Student
	)
	sid = u.GetString(":sid")
	if sid != "" {
		stu, err = models.GetStudentById(sid)
		if err != nil {
			u.Data["json"] = map[string]string{"error": err.Error()}
		} else {
			u.Data["json"] = map[string]models.Student{"student": stu}
		}
	} else {

	}
	u.ServeJSON()
}

// LearnProject
// @Title
// @Description
// @Param body body models.LearnProject true ""
// @Success 200 {object} models.Project
// @Failure 403 :id is empty
// @router /learn/:pid [post]
func (u *StudentController) LearnProject() {

	learning, err := u.GetBool("learning")
	sid, err := u.GetInt64("studentId")
	pid, err := u.GetInt64(":pid")

	l := &models.LearnProject{
		StudentId: sid,
		ProjectId: pid,
		Learning:  learning,
	}
	err = json.Unmarshal(u.Ctx.Input.RequestBody, &l)
	if err != nil {
		u.Data["json"] = map[string]string{"error": err.Error()}
	}
	err = l.Create()
	if err != nil {
		u.Data["json"] = map[string]string{"error": err.Error()}
	}
	u.Data["json"] = map[string]bool{"result": true}
	u.ServeJSON()
}

// FinishedProject
// @Title
// @Description
// @Param body body models.LearnProject true ""
// @Success 200 {object} models.Project
// @Failure 403 :id is empty
// @router /finished [post]
func (u *StudentController) FinishedProject() {
	var l models.LearnProject
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &l)

	if err != nil {
		u.Data["json"] = map[string]string{"error": err.Error()}
	}
	err = l.Update()

	if err != nil {
		u.Data["json"] = map[string]string{"error": err.Error()}
	}
	u.Data["json"] = map[string]bool{"result": true}
	u.ServeJSON()
}
