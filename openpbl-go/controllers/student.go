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
// @router /learn/:pid [post]
func (u *StudentController) LearnProject() {
	pid, err := u.GetInt64(":pid")
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
