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
	uuid "github.com/satori/go.uuid"
	"xorm.io/xorm"
)

type Resource struct {
	Id                string   `json:"id" xorm:"not null pk"`
	SectionId         string   `json:"sectionId" xorm:"not null index unique"`
	Content           string  `json:"content" xorm:"longtext"`
}

func (p *Resource) GetEngine() *xorm.Session {
	return adapter.Engine.Table(p)
}
func (p *Resource) Create() (err error) {
	_, err = p.GetEngine().Insert(p)
	return
}
func (p *Resource) Update() (err error) {
	_, err = p.GetEngine().
		ID(p.Id).
		Update(p)
	return
}
func GetResourceById(id string) (r Resource, err error) {
	_, err = (&Resource{}).GetEngine().
		ID(id).
		Get(&r)
	return
}

func DeleteSectionResource(sid string) (err error) {
	_, err = (&Resource{}).GetEngine().Where("section_id = ?", sid).Delete(&Resource{})
	return
}

func CloneSectionResource(sid string, newSid string) (err error) {
	var resource Resource
	_, err = (&Resource{}).GetEngine().Where("section_id = ?", sid).Get(&resource)
	resource.SectionId = newSid
	resource.Id = uuid.NewV4().String()
	_, err = (&Resource{}).GetEngine().Insert(&resource)
	return
}