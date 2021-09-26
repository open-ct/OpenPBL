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

package models

import (
	"time"
	"xorm.io/xorm"
)

type Message struct {
	Id            string    `json:"id" xorm:"not null pk"`
	ReceiverId    string   	`json:"receiverId" xorm:"index"`

	MessageType   string    `json:"messageType" xorm:"not null index"`  // warn error info

	MessageTitle  string    `json:"messageTitle" xorm:"text"`
	Content       string    `json:"content" xorm:"longtext"`

	ReadMessage   bool      `json:"readMessage" xorm:"default false index"`

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

func ReadMessage(messageId string, uid string) (err error) {
	_, err = (&Message{}).GetEngine().
		ID(messageId).
		MustCols("read_message").
		Update(&Message{
			Id:   messageId,
			ReceiverId: uid,
			ReadMessage: true,
		})
	return
}
func ReadAllMessage(uid string) (err error) {
	_, err = (&Message{}).GetEngine().
		Where("receiver_id = ?", uid).
		MustCols("read_message").
		Update(&Message{
			ReceiverId: uid,
			ReadMessage: true,
		})
	return
}

func GetMessages(uid string, orderType string, messageType string, readType string, from int, size int) (m []Message, rows int64, err error) {
	s := (&Message{}).GetEngine().
		Where("receiver_id = ?", uid)
	s2 := (&Message{}).GetEngine().
		Where("receiver_id = ?", uid)
	if readType == "read" {
		s = s.Where("read_message = true")
		s2 = s2.Where("read_message = true")
	} else if readType == "unread" {
		s = s.Where("read_message = false")
		s2 = s2.Where("read_message = false")
	}
	rows, err = s2.Count(&Message{})
	if orderType == "asc" {
		s = s.Asc("create_at")
	} else {
		s = s.Desc("create_at")
	}
	err = s.Desc("create_at").Limit(size, from).Find(&m)
	return
}

func CreateMessage(msg *Message) bool {
	err := msg.Create()
	if err != nil {
		return false
	}
	return true
}