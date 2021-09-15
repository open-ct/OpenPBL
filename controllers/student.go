package controllers

import (
	"OpenPBL/models"
	"OpenPBL/util"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/casdoor/casdoor-go-sdk/auth"
	"time"
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
	if !util.IsStudent(user) {
		resp = Response{
			Code: 403,
			Msg:  "非法的用户",
		}
		u.Data["json"] = resp
		u.ServeJSON()
		return
	}
	uid := util.GetUserId(user)

	l := &models.LearnProject{
		Avatar:    user.Avatar,
		Name:      user.Name,
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

	project, err := models.GetProjectById(pid)
	if err == nil {
		content := fmt.Sprintf("成功加入课程 \"%v\" ", project.ProjectTitle)
		models.CreateMessage(&models.Message{
			ReceiverId:   uid,
			MessageType:  "info",
			MessageTitle: "加入课程",
			Content:      content,
			ReadMessage:  false,
			CreateAt:     time.Time{},
		})
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
	if !util.IsStudent(user) {
		resp = Response{
			Code: 403,
			Msg:  "非法的用户",
		}
		u.Data["json"] = resp
		u.ServeJSON()
		return
	}
	uid := util.GetUserId(user)

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

	project, err := models.GetProjectById(pid)
	if err == nil {
		content := fmt.Sprintf("成功退出课程 \"%v\" ", project.ProjectTitle)
		models.CreateMessage(&models.Message{
			ReceiverId:   uid,
			MessageType:  "info",
			MessageTitle: "退出课程",
			Content:      content,
			ReadMessage:  false,
			CreateAt:     time.Time{},
		})
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

// GetLearnSection
// @Title
// @Description
// @Param body body models.LearnSection true ""
// @Success 200 {object} Response
// @Failure 400
// @router /project/:projectId/section/:sectionId [get]
func (u *StudentController) GetLearnSection() {
	var resp Response
	user := u.GetSessionUser()
	if !util.IsStudent(user) {
		resp = Response{
			Code: 403,
			Msg:  "非法的用户",
		}
		u.Data["json"] = resp
		u.ServeJSON()
		return
	}
	uid := util.GetUserId(user)
	sid, err := u.GetInt64(":sectionId")
	projectId, err := u.GetInt64(":projectId")
	l, _ := models.GetLearnSection(sid, uid, projectId)
	if err != nil {
		u.Data["json"] = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	} else {
		u.Data["json"] = Response{
			Code: 200,
			Data: l,
		}
	}
	u.ServeJSON()
}

// UpdateLearnSection
// @Title
// @Description
// @Param body body models.LearnSection true ""
// @Success 200 {object} Response
// @Failure 400
// @router /project/:projectId/section/:sectionId [post]
func (u *StudentController) UpdateLearnSection() {
	var resp Response
	user := u.GetSessionUser()
	if !util.IsStudent(user) {
		resp = Response{
			Code: 403,
			Msg:  "非法的用户",
		}
		u.Data["json"] = resp
		u.ServeJSON()
		return
	}
	uid := util.GetUserId(user)
	sid, err := u.GetInt64(":sectionId")
	m, err := u.GetInt("learnMinute")
	s, err := u.GetInt("learnSecond")
	l := models.LearnSection{
		StudentId:     uid,
		SectionId:     sid,
		LearnMinute:   m,
		LearnSecond:   s,
	}
	pid, err := u.GetInt64(":projectId")
	err = l.Update(pid)
	if err != nil {
		u.Data["json"] = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	} else {
		u.Data["json"] = Response{
			Code: 200,
		}
	}
	u.ServeJSON()
}

// GetLastLearnSection
// @Title
// @Description
// @Success 200 {object} Response
// @Failure 400
// @router /last-learn/project/:projectId [get]
func (u *StudentController) GetLastLearnSection() {
	var resp Response
	user := u.GetSessionUser()
	if !util.IsStudent(user) {
		resp = Response{
			Code: 200,
			Msg:  "",
		}
		u.Data["json"] = resp
		u.ServeJSON()
		return
	}
	uid := util.GetUserId(user)
	projectId := u.GetString(":projectId")
	l, err := models.GetLastLearnSection(uid, projectId)
	if err != nil {
		u.Data["json"] = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	} else {
		u.Data["json"] = Response{
			Code: 200,
			Data: l,
		}
	}
	u.ServeJSON()
}