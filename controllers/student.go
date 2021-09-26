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
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/casdoor/casdoor-go-sdk/auth"
	"time"
)


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
// @Title LearnProject
// @Description
// @Param projectId path string true "The id of the project"
// @Success 200 {object} Response
// @router /learn/:projectId [post]
func (u *StudentController) LearnProject() {
	pid := u.GetString(":projectId")
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
	err := l.Create()
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
		models.CreateMessage(&models.Message{
			Id:           util.NewId(),
			ReceiverId:   uid,
			MessageType:  "info",
			MessageTitle: "加入课程",
			Content:      fmt.Sprintf("成功加入课程 \"%v\" ", project.ProjectTitle),
			ReadMessage:  false,
			CreateAt:     time.Time{},
		})
		models.CreateMessage(&models.Message{
			Id:           util.NewId(),
			ReceiverId:   project.TeacherId,
			MessageType:  "info",
			MessageTitle: "加入课程",
			Content:      fmt.Sprintf("学生 %v 加入课程 \"%v\" ", user.DisplayName, project.ProjectTitle),
			ReadMessage:  false,
			CreateAt:     time.Time{},
		})
	}
	u.Data["json"] = resp
	u.ServeJSON()
}

// ExitProject
// @Title ExitProject
// @Description
// @Param projectId path string true "The id of the project"
// @Success 200 {object} Response
// @router /exit/:projectId [post]
func (u *StudentController) ExitProject() {
	pid := u.GetString(":projectId")
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
	err := l.Delete()
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
		models.CreateMessage(&models.Message{
			Id:           util.NewId(),
			ReceiverId:   uid,
			MessageType:  "info",
			MessageTitle: "退出课程",
			Content:      fmt.Sprintf("成功退出课程 \"%v\" ", project.ProjectTitle),
			ReadMessage:  false,
			CreateAt:     time.Time{},
		})
		models.CreateMessage(&models.Message{
			Id:           util.NewId(),
			ReceiverId:   project.TeacherId,
			MessageType:  "info",
			MessageTitle: "退出课程",
			Content:      fmt.Sprintf("学生 %v 退出课程 \"%v\" ", user.DisplayName, project.ProjectTitle),
			ReadMessage:  false,
			CreateAt:     time.Time{},
		})
	}

	u.Data["json"] = resp
	u.ServeJSON()
}

// FinishedProject
// @Title FinishedProject
// @Description
// @Success 200 {object} models.Project
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
// @Title GetLearnSection
// @Description
// @Param projectId path string true "The id of the project"
// @Param sectionId path string true "The id of the section"
// @Success 200 {object} Response
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
	sid := u.GetString(":sectionId")
	projectId := u.GetString(":projectId")
	l, err := models.GetLearnSection(sid, uid, projectId)
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
// @Title UpdateLearnSection
// @Description
// @Param sectionId path string true "The id of the section"
// @Param learnMinute path string true "The learning minute"
// @Param learnSecond path string true "The learning second"
// @Success 200 {object} Response
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
	sid := u.GetString(":sectionId")
	m, err := u.GetInt("learnMinute")
	s, err := u.GetInt("learnSecond")
	l := models.LearnSection{
		StudentId:     uid,
		SectionId:     sid,
		LearnMinute:   m,
		LearnSecond:   s,
	}
	pid := u.GetString(":projectId")
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
// @Title GetLastLearnSection
// @Description
// @Param projectId path string true "The id of the project"
// @Success 200 {object} Response
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