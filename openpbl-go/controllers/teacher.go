package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"openpbl-go/models"
	"strconv"
	"time"
)

// TeacherController
// Operations about Student
type TeacherController struct {
	beego.Controller
}

// CreateTeacher
// @Title Create Teacher
// @Description create a teacher
// @Param body body models.Teacher true	"body for teacher content"
// @Success 200 {int} models.Teacher.Id
// @Failure 403 body is empty
// @router / [post]
func (u *TeacherController) CreateTeacher() {
	var (
		err     error
		teacher models.Teacher
		uid     int64
	)
	err = json.Unmarshal(u.Ctx.Input.RequestBody, &teacher)
	if err != nil {
		u.Data["json"] = map[string]string{"error": err.Error()}
	}
	err = teacher.Create()
	if err != nil {
		u.Data["json"] = map[string]string{"error": err.Error()}
	}
	uid = teacher.Id
	u.Data["json"] = map[string]string{"id": strconv.FormatInt(uid, 10)}
	u.ServeJSON()
}

// CheckTeacherEmail
// @Title CheckTeacherEmail
// @Description
// @Param email	path string	true ""
// @Success 200 {exist: true}
// @Failure 403 :email is empty
// @router /teacher/checkout/:email [get]
func (u *TeacherController) CheckTeacherEmail() {
	var (
		err   error
		exist bool
		email string
	)
	email = u.GetString(":email")
	if email != "" {
		exist, err = models.ExistTeacherEmail(email)
		if err != nil {
			u.Data["json"] = map[string]string{"error": err.Error()}
		}
		u.Data["json"] = map[string]bool{"exist": exist}
	} else {

	}
	u.ServeJSON()
}

// PublishProject
// @Title Create Teacher
// @Description create a teacher
// @Param pid path string true ""
// @Success 200 {int}
// @Failure 403 body is empty
// @router /publish [post]
func (u *TeacherController) PublishProject() {
	pid, err := u.GetInt64("pid")
	p := models.Project{
		Id:               pid,
		PublishedAt:      time.Now(),
		Published:        true,
	}

	err = models.UpdatePublished(p)
	if err != nil {
		u.Data["json"] = map[string]string{"error": err.Error()}
	}
	u.Data["json"] = map[string]bool{"result": true}
	u.ServeJSON()
}