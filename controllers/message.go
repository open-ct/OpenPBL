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
)


type MessageController struct {
	beego.Controller
}

func (m *MessageController) GetSessionUser() *auth.Claims {
	s := m.GetSession("user")
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


type MessagesResponse struct {
	Messages    []models.Message    `json:"messages"`
	Count       int64               `json:"count"`
	Code        int                 `json:"code"`
}

// GetUserMessages
// @Title GetUserMessages
// @Description Get user messages
// @Param readType params string true "read unread all"
// @Param messageType params string true "info warn error all"
// @Param from params int false "page number"
// @Param size params int false "page size"
// @Param orderType params string false "desc asc"
// @Success 200 {object} MessagesResponse
// @router / [get]
func (m *MessageController) GetUserMessages() {
	user := m.GetSessionUser()
	from, err := m.GetInt("from")
	if err != nil {
		from = 0
	}
	size, err := m.GetInt("size")
	if err != nil {
		size = 10
	}
	orderType := m.GetString("orderType")
	r := m.GetString("readType")
	t := m.GetString("messageType")
	uid := util.GetUserId(user)

	var messages []models.Message
	var rows     int64
	messages, rows, err = models.GetMessages(uid, orderType, t, r, from, size)

	if err != nil {
		m.Data["json"] = MessagesResponse{
			Code:    400,
		}
	} else {
		m.Data["json"] = MessagesResponse{
			Code:    200,
			Messages: messages,
			Count:    rows,
		}
	}
	m.ServeJSON()
}

// ReadUserMessage
// @Title ReadUserMessage
// @Description Set one message to read state
// @Param messageId path string true "message id"
// @Success 200 {object} Response
// @router /:messageId/read [post]
func (m *MessageController) ReadUserMessage() {
	user := m.GetSessionUser()
	uid := util.GetUserId(user)
	mid := m.GetString(":messageId")
	err := models.ReadMessage(mid, uid)
	if err != nil {
		m.Data["json"] = Response{
			Code:    400,
			Msg:     err.Error(),
		}
	} else {
		m.Data["json"] = Response{
			Code:    200,
		}
	}
	m.ServeJSON()
}

// DeleteUserMessage
// @Title DeleteUserMessage
// @Description Delete a message
// @Param messageId path string true "message id"
// @Success 200 {object} Response
// @router /:messageId/delete [post]
func (m *MessageController) DeleteUserMessage() {
	user := m.GetSessionUser()
	uid := util.GetUserId(user)
	mid := m.GetString(":messageId")
	msg := models.Message{
		Id:           mid,
		ReceiverId:   uid,
	}
	err := msg.Delete()
	if err != nil {
		m.Data["json"] = Response{
			Code:    400,
			Msg:     err.Error(),
		}
	} else {
		m.Data["json"] = Response{
			Code:    200,
		}
	}
	m.ServeJSON()
}

// ReadAllUserMessage
// @Title ReadAllUserMessage
// @Description Set all messages to read state
// @Success 200 {object} Response
// @router /read-all [post]
func (m *MessageController) ReadAllUserMessage() {
	user := m.GetSessionUser()
	uid := util.GetUserId(user)
	err := models.ReadAllMessage(uid)
	if err != nil {
		m.Data["json"] = Response{
			Code:    400,
			Msg:     err.Error(),
		}
	} else {
		m.Data["json"] = Response{
			Code:    200,
		}
	}
	m.ServeJSON()
}
