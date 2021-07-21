package controllers

import (
	"github.com/astaxie/beego"
	"github.com/casdoor/casdoor-go-sdk/auth"
	"openpbl-go/models"
	"openpbl-go/util"
	"time"
)

// TeacherController
// Operations about Student
type TeacherController struct {
	beego.Controller
}

func (u *TeacherController) GetSessionUser() *auth.Claims {
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

// PublishProject
// @Title
// @Description create a teacher
// @Param pid path string true ""
// @Success 200 {int}
// @Failure 403 body is empty
// @router /publish [post]
func (u *TeacherController) PublishProject() {
	pid, err := u.GetInt64("pid")
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
	if user.Tag != "teacher" {
		resp = Response{
			Code: 403,
			Msg:  "非法的用户",
		}
		u.Data["json"] = resp
		u.ServeJSON()
		return
	}

	p := models.Project{
		Id:               pid,
		PublishedAt:      time.Now(),
		Published:        true,
	}
	err = models.UpdatePublished(p)
	if err != nil {
		resp = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	} else {
		resp = Response{
			Code: 200,
			Msg:  "发布成功",
		}
	}
	u.Data["json"] = resp
	u.ServeJSON()
}