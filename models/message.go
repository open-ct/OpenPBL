package models

import (
	"time"
	"xorm.io/xorm"
)

type Message struct {
	Id            int64    	`json:"id" xorm:"not null pk autoincr"`
	ReceiverId    string   	`json:"receiverId" xorm:"index"`

	MessageType   string    `json:"messageType" xorm:"not null index"`  // warn error info

	MessageTitle  string    `json:"messageTitle" xorm:"text"`
	Content       string    `json:"content" xorm:"longtext"`

	Read          bool      `json:"read" xorm:"default false index"`

	CreateAt      time.Time `json:"createAt" xorm:"created"`
}


func (m *Message) GetEngine() *xorm.Session {
	return adapter.Engine.Table(m)
}

func (m *Message) Create() (err error) {
	_, err = m.GetEngine().Insert(m)
	return
}
func (m *Message) Delete() (err error) {
	_, err = m.GetEngine().ID(m.Id).Delete(m)
	return
}

func ReadMessage(messageId int64, uid string) (err error) {
	_, err = (&Message{}).GetEngine().
		ID(messageId).
		Update(&Message{
			Id:   messageId,
			ReceiverId: uid,
			Read: true,
		})
	return
}
func ReadAllMessage(uid string) (err error) {
	_, err = (&Message{}).GetEngine().
		Update(&Message{
			ReceiverId: uid,
			Read: true,
		})
	return
}

func GetMessages(uid string, orderType string, messageType string, read string, from int, size int) (m []Message, rows int64, err error) {
	s := (&Message{}).GetEngine().
		Where("receiver_id = ?", uid)
	s2 := (&Message{}).GetEngine().
		Where("receiver_id = ?", uid)
	if read == "read" {
		s = s.Where("read = true")
		s2 = s2.Where("read = true")
	} else if read == "unread" {
		s = s.Where("read = false")
		s2 = s2.Where("read = false")
	}
	if messageType == "error" || messageType == "info" || messageType == "warn" {
		s = s.Where("message_type = ?", messageType)
		s2 = s2.Where("message_type = ?", messageType)
	}

	rows, err = s2.Count()

	if orderType == "asc" {
		s = s.Asc("create_at")
	} else {
		s = s.Desc("create_at")
	}

	err = s.Limit(size, from).Find(&m)
	return
}