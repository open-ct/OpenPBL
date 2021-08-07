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
// @Param type path string true "received sent"
// @Param from params int false ""
// @Param size params int false ""
// @Success 200 {object} MessagesResponse
// @Failure 400
// @router /:type [get]
func (m *MessageController) GetUserMessages() {
	user := m.GetSessionUser()
	var resp ProjectResponse
	if user == nil {
		resp = ProjectResponse{
			Code:    401,
			Msg:     "请先登录",
		}
		m.Data["json"] = resp
		m.ServeJSON()
		return
	}
	from, err := m.GetInt("from")
	if err != nil {
		from = 0
	}
	size, err := m.GetInt("size")
	if err != nil {
		size = 10
	}
	t := m.GetString(":type")
	uid := user.Username

	var messages []models.Message
	var rows     int64
	if t == "received" {
		messages, rows, err = models.GetReceivedMessages(uid, "desc", from, size)
	} else if t == "sent" {
		messages, rows, err = models.GetSentMessages(uid, "desc", from, size)
	}
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