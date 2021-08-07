package models

import (
	"time"
	"xorm.io/xorm"
)

type Message struct {
	Id           int64     `json:"id" xorm:"not null pk"`
	ReceiverId   string    `json:"receiverId" xorm:"not null index"`

	SenderId     string    `json:"senderId" xorm:"index"`
	SenderName   string    `json:"senderName"`

	ProjectId    int64     `json:"projectId" xorm:"index"`
	ProjectTitle string    `json:"projectTitle"`

	MessageType string    `json:"messageType" xorm:"not null index"`  // remind message
	Content   string    `json:"content" xorm:"text"`

	Read      bool      `json:"read" xorm:"default false index"`

	CreateAt  time.Time `json:"CreateAt" xorm:"created"`
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

func ReadMessage(messageId int64) (err error) {
	_, err = (&Message{}).GetEngine().
		ID(messageId).
		Update(&Message{
			Id:   messageId,
			Read: true,
		})
	return
}

func GetReceivedMessages(uid string, orderType string, from int, size int) (m []Message, rows int64, err error) {
	err = (&Message{}).GetEngine().
		Where("receiver_id = ?", uid).
		Desc("create_at").
		Limit(size, from).
		Find(&m)
	rows, err = (&Message{}).GetEngine().
		Count(&Message{ReceiverId:   uid})
	return
}

func GetSentMessages(uid string, orderType string, from int, size int) (m []Message, rows int64, err error) {
	err = (&Message{}).GetEngine().
		Where("sender_id = ?", uid).
		Desc("create_at").
		Limit(size, from).
		Find(&m)
	rows, err = (&Message{}).GetEngine().
		Count(&Message{SenderId:   uid})
	return
}