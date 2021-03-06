package models

import (
	"xorm.io/xorm"
)

type Resource struct {
	Id                int64   `json:"id" xorm:"not null pk autoincr"`
	SectionId         int64   `json:"sectionId" xorm:"not null index unique"`
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

func DeleteSectionResource(sid int64) (err error) {
	_, err = (&Resource{}).GetEngine().Where("section_id = ?", sid).Delete(&Resource{})
	return
}

func CloneSectionResource(sid int64, newSid int64) (err error) {
	var resource Resource
	_, err = (&Resource{}).GetEngine().Where("section_id = ?", sid).Get(&resource)
	resource.SectionId = newSid
	resource.Id = 0
	_, err = (&Resource{}).GetEngine().Insert(&resource)
	return
}