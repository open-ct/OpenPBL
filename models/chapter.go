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
	"OpenPBL/util"
	"fmt"
	"xorm.io/xorm"
)

type Chapter struct {
	Id                 string   `json:"id" xorm:"not null pk"`
	ProjectId          string   `json:"projectId" xorm:"index"`
	ChapterName        string  `json:"chapterName"`
	ChapterNumber      int     `json:"chapterNumber" xorm:"index"`
}

type Outline struct {
	Chapter                  `xorm:"extends"`
	Sections []SectionMinute `json:"sections" xorm:"extends"`
}

func (p *Chapter) GetEngine() *xorm.Session {
	return adapter.Engine.Table(p)
}

func (p *Chapter) Create() (err error) {
	_, err = p.GetEngine().Insert(p)
	return
}
func (p *Chapter) Update() (err error) {
	_, err = p.GetEngine().ID(p.Id).Update(p)
	return
}
func (p *Chapter) Delete() (err error) {
	session := adapter.Engine.NewSession()
	defer session.Close()
	session.Begin()
	_, err = session.
		Exec("update chapter set chapter_number = chapter_number - 1 " +
			"where project_id = ? and chapter_number > ?", p.ProjectId, p.ChapterNumber)
	fmt.Println(err)
	_, err = session.Table(&Section{}).Delete(Section{ChapterId: p.Id})
	_, err = session.Table(&Chapter{}).ID(p.Id).Delete(p)
	session.Commit()
	return
}
func ExchangeChapters(cid1 string, cid2 string) (err error) {
	_, err = adapter.Engine.
		Exec("update chapter c1 join chapter c2 on (c1.id = ? and c2.id = ?) " +
			"set c1.chapter_number = c2.chapter_number, c2.chapter_number = c1.chapter_number", cid1, cid2)
	return
}

func GetChaptersByPid(pid string, uid string) (outline []Outline, err error) {
	var c []Chapter
	err = (&Chapter{}).GetEngine().
		Where("project_id = ?", pid).
		Asc("chapter_number").
		Find(&c)
	outline = make([]Outline, len(c))
	for i:=0; i< len(c); i++ {
		sections := make([]SectionMinute, 0)
		outline[i].Chapter = c[i]
		if uid == "" {
			err = (&Section{}).GetEngine().
				Where("chapter_id = ?", c[i].Id).
				Asc("section_number").
				Find(&sections)
		} else {
			err = adapter.Engine.
				SQL("select * from (select * from section where chapter_id = ?) s LEFT JOIN learn_section ls on s.id = ls.section_id and ls.student_id = ?  order by s.section_number", c[i].Id, uid).
				Find(&sections)
		}
		outline[i].Sections = sections
	}
	return
}

func DeleteProjectChapters(pid string) (err error) {
	var chapters []Chapter
	err = (&Chapter{}).GetEngine().Where("project_id = ?", pid).Find(&chapters)
	for i:=0; i< len(chapters); i++ {
		c := chapters[i]
		cid := c.Id
		_, err = (&Chapter{}).GetEngine().ID(cid).Delete(&Chapter{})
		err = DeleteChapterSections(cid)
	}
	return
}

func CloneProjectChapters(pid string, newPid string) (err error) {
	var chapters []Chapter
	err = (&Chapter{}).GetEngine().Where("project_id = ?", pid).Find(&chapters)
	for i:=0; i< len(chapters); i++ {
		c := chapters[i]
		cid := c.Id
		c.Id = util.NewId()
		c.ProjectId = newPid
		_, err = (&Chapter{}).GetEngine().Insert(&c)
		newCid := c.Id
		err = CloneChapterSections(newPid, cid, newCid)
	}
	return
}