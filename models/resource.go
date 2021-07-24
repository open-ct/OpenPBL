package models

import (
	"xorm.io/xorm"
)

type Resource struct {
	Id                int64   `json:"id" xorm:"not null pk autoincr"`
	SectionId         int64   `json:"sectionId" xorm:"not null index unique"`

	Content           string  `json:"content" xorm:"longtext"`

	FileTitle         string  `json:"fileTitle"`
	FileIntroduce     string  `json:"fileIntroduce"`
	FilePath          string  `json:"filePath"`
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