package controllers

import (
	"OpenPBL/models"
	"OpenPBL/util"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/casdoor/casdoor-go-sdk/auth"
)

// StudentController
// Operations about Student
type StudentController struct {
	beego.Controller
}

func (u *StudentController) GetSessionUser() *auth.Claims {
	s := u.GetSession("user")
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

// LearnProject
// @Title
// @Description
// @Param pid path string true ""
// @Success 200 {object} Response
// @Failure 401
// @Failure 403
// @router /learn/:projectId [post]
func (u *StudentController) LearnProject() {
	pid, err := u.GetInt64(":projectId")
	var resp Response
	user := u.GetSessionUser()
	if user == nil {
		resp = Response{
			Code: 401,
			Msg:  "请先登录",
		}
		u.Data["json"] = resp
		u.ServeJSON()
		return
	}
	if user.Tag != "student" {
		resp = Response{
			Code: 403,
			Msg:  "非法的用户",
		}
		u.Data["json"] = resp
		u.ServeJSON()
		return
	}
	uid := user.Username

	l := &models.LearnProject{
		StudentId: uid,
		ProjectId: pid,
		Learning:  true,
	}
	err = l.Create()
	if err != nil {
		resp = Response{
			Code: 400,
			Msg:  err.Error(),
		}
		u.Data["json"] = resp
	}
	resp = Response{
		Code: 200,
		Msg:  "加入成功",
	}
	u.Data["json"] = resp
	u.ServeJSON()
}

// ExitProject
// @Title
// @Description
// @Param pid path string true ""
// @Success 200 {object} Response
// @Failure 401
// @Failure 403
// @router /exit/:projectId [post]
func (u *StudentController) ExitProject() {
	pid, err := u.GetInt64(":projectId")
	var resp Response
	user := u.GetSessionUser()
	if user == nil {
		resp = Response{
			Code: 401,
			Msg:  "请先登录",
		}
		u.Data["json"] = resp
		u.ServeJSON()
		return
	}
	if user.Tag != "student" {
		resp = Response{
			Code: 403,
			Msg:  "非法的用户",
		}
		u.Data["json"] = resp
		u.ServeJSON()
		return
	}
	uid := user.Username

	l := &models.LearnProject{
		StudentId: uid,
		ProjectId: pid,
	}
	err = l.Delete()
	if err != nil {
		resp = Response{
			Code: 400,
			Msg:  err.Error(),
		}
		u.Data["json"] = resp
	}
	resp = Response{
		Code: 200,
		Msg:  "退出成功",
	}
	u.Data["json"] = resp
	u.ServeJSON()
}

// FinishedProject
// @Title
// @Description
// @Param body body models.LearnProject true ""
// @Success 200 {object} models.Project
// @Failure 403
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

// GetProjectStudents
// todo need refactor
// @Title
// @Description
// @Param from query int true ""
// @Param size query int true ""
// @Success 200 {object} []models.StudentInfo
// @Failure 403 body is empty
// @router /students/:pid [get]
func (p *ProjectController) GetProjectStudents() {
	pid := p.GetString(":pid")
	from, err := p.GetInt("from")
	if err != nil {
		from = 0
	}
	size, err := p.GetInt("size")
	if err != nil {
		size = 10
	}
	if pid != "" {
		students, err := models.GetProjectStudents(pid, from, size)
		rows, err := models.CountProjectStudents(pid, from, size)
		if err != nil {
			p.Data["json"] = map[string]string{"error": err.Error()}
		}
		p.Data["json"] = StudentList{
			Count: rows,
			Students: students,
		}
	}
	p.ServeJSON()
}
