package models

import (
	"xorm.io/xorm"
)

type Section struct {
	Id                 int64   `json:"id" xorm:"not null pk autoincr"`
	ChapterId          int64   `json:"chapterId" xorm:"index"`
	SectionName        string  `json:"sectionName"`
	SectionNumber      int     `json:"sectionNumber" xorm:"index"`
}

type SectionDetail struct {
	Section     `xorm:"extends"`
	Resource    `json:"resource" xorm:"extends"`
}

func (p *Section) GetEngine() *xorm.Session {
	return adapter.Engine.Table(p)
}

func (p *Section) Create() (err error) {
	session := adapter.Engine.NewSession()
	defer session.Close()
	session.Begin()
	_, err = session.Insert(p)
	_, err = session.Insert(Resource{
		SectionId:         p.Id,
	})
	session.Commit()
	return
}
func (p *Section) Update() (err error) {
	_, err = p.GetEngine().ID(p.Id).Update(p)
	return
}
func (p *Section) Delete() (err error) {
	_, err = p.GetEngine().ID(p.Id).Delete(p)
	return
}
func ExchangeSections(id1 string, id2 string) (err error) {
	_, err = adapter.Engine.
		Exec("update section c1 join section c2 on (c1.id = ? and c2.id = ?) " +
			"set c1.section_number = c2.section_number, c2.section_number = c1.section_number", id1, id2)
	return
}

func GetSectionsByCid(cid string) (s []Section, err error) {
	err = (&Section{}).GetEngine().
		Where("chapter_id = ?", cid).
		Asc("section_number").
		Find(&s)
	return
}

func GetSectionById(sid string) (s Section, err error) {
	_, err = (&Section{}).GetEngine().
		Where("id = ?", sid).
		Get(&s)
	return
}

func GetSectionDetailById(sid string) (s SectionDetail, err error) {
	_, err = adapter.Engine.
		Table("section").
		Where("section.id = ?", sid).
		Join("INNER", "resource", "resource.section_id = section.id").
		Get(&s)
	return
}