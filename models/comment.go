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

type Comment struct {
	Id           string    `json:"id" xorm:"not null pk"`
	ProjectId    string    `json:"projectId" xorm:"not null index"`
	UserId       string    `json:"userId" xorm:"not null index"`
	UserAvatar   string    `json:"userAvatar" xorm:"not null"`
	UserName     string    `json:"userName" xorm:"not null"`
	IsTeacher    bool      `json:"isTeacher" xorm:"not null index default false"`

	Content      string    `json:"content" xorm:"text"`

	CreateAt     time.Time `json:"createAt" xorm:"created"`
}

func (c *Comment) GetEngine() *xorm.Session {
	return adapter.Engine.Table(c)
}

func (c *Comment) Create() (err error) {
	_, err = c.GetEngine().Insert(c)
	return
}

func (c *Comment) Delete() (err error) {
	_, err = c.GetEngine().Delete(c)
	return
}

func GetProjectComments(pid string, from int, size int, text string, teacher bool) (comments []Comment, rows int64, err error) {
	s := (&Comment{}).GetEngine().Where("project_id = ?", pid)
	s2 := (&Comment{}).GetEngine().Where("project_id = ?", pid)
	if teacher {
		s = s.Where("is_teacher = true")
		s2 = s2.Where("is_teacher = true")
	}
	if text != "" {
		s = s.Where("content like %?%", text)
		s2 = s2.Where("content like %?%", text)
	}
	s = s.Desc("create_at").Limit(size, from)
	rows, err = s2.Count(&Comment{})
	err = s.Find(&comments)
	return
}