package controllers

import (
	"OpenPBL/models"
	"OpenPBL/util"
	"github.com/astaxie/beego"
	"github.com/casdoor/casdoor-go-sdk/auth"
)

// MessageController
// Operations about Projects
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
// @Title
// @Description
// @Param readType params string true "read unread all"
// @Param messageType params string true "info warn error all"
// @Param from params int false ""
// @Param size params int false ""
// @Param orderType params string false "desc asc"
// @Success 200 {object} MessagesResponse
// @Failure 400
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
// @Title
// @Description
// @Param messageId path string true ""
// @Success 200 {object} Response
// @Failure 400
// @router /:messageId/read [post]
func (m *MessageController) ReadUserMessage() {
	user := m.GetSessionUser()
	uid := util.GetUserId(user)
	mid, err := m.GetInt64(":messageId")
	err = models.ReadMessage(mid, uid)
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
// @Title
// @Description
// @Param messageId path string true ""
// @Success 200 {object} Response
// @Failure 400
// @router /:messageId/delete [post]
func (m *MessageController) DeleteUserMessage() {
	user := m.GetSessionUser()
	uid := util.GetUserId(user)
	mid, err := m.GetInt64(":messageId")
	msg := models.Message{
		Id:           mid,
		ReceiverId:   uid,
	}
	err = msg.Delete()
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
// @Title
// @Description
// @Success 200 {object} Response
// @Failure 400
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




func CreateMessage(msg *models.Message) bool {
	err := msg.Create()
	if err != nil {
		return false
	}
	return true
}
