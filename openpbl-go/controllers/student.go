package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"openpbl-go/models"
	"strconv"
)

// StudentController
// Operations about Student
type StudentController struct {
	beego.Controller
}

// CreateStudent
// @Title Create Student
// @Description create a student
// @Param body body models.Student true	"body for student content"
// @Success 200 {int} models.Student.Id
// @Failure 403 body is empty
// @router / [post]
func (u *StudentController) CreateStudent() {
	var (
		err     error
		student models.Student
		uid     int64
	)
	err = json.Unmarshal(u.Ctx.Input.RequestBody, &student)
	if err != nil {
		u.Data["json"] = map[string]string{"error": err.Error()}
	}
	err = student.Create()
	if err != nil {
		u.Data["json"] = map[string]string{"error": err.Error()}
	}
	uid = student.Id
	u.Data["json"] = map[string]string{"id": strconv.FormatInt(uid, 10)}
	u.ServeJSON()
}

// CheckStudentEmail
// @Title CheckStudentEmail
// @Description
// @Param email	path string	true ""
// @Success 200 {exist: true}
// @Failure 403 :email is empty
// @router /checkout/:email [get]
func (u *StudentController) CheckStudentEmail() {
	var (
		err   error
		exist bool
		email string
	)
	email = u.GetString(":email")
	if email != "" {
		exist, err = models.ExistStudentEmail(email)
		if err != nil {
			u.Data["json"] = map[string]string{"error": err.Error()}
		}
		u.Data["json"] = map[string]bool{"exist": exist}
	} else {

	}
	u.ServeJSON()
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
// @router /learn [post]
func (u *StudentController) LearnProject() {
	learning, err := u.GetBool("learning")
	sid, err := u.GetInt64("studentId")
	pid, err := u.GetInt64("projectId")

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
