package models

import "xorm.io/xorm"

type Resource struct {
	Id                int64   `json:"id" xorm:"not null pk autoincr"`
	ResourceType      string  `json:"resourceType"`
	ResourceTitle     string  `json:"resourceTitle"`
	ResourceIntroduce string  `json:"resourceIntroduce" xorm:"text"`
	ResourceOrder     int     `json:"resourceOrder"`

	Content          string  `json:"content" xorm:"long text"`
	FilePath         string  `json:"filePath" xorm:"index"`

	SectionId        int64   `json:"sectionId" xorm:"not null index"`
}
func (p *Resource) GetEngine() *xorm.Session {
	return adapter.Engine.Table(p)
}
func (p *Resource) Create() (err error) {
	_, err = p.GetEngine().Insert(p)
	return
}
func (p *Resource) Update() (err error) {
	_, err = p.GetEngine().ID(p.Id).Update(p)
	return
}
func GetResourceById(id string) (r Resource, err error) {
	_, err = (&Resource{}).GetEngine().ID(id).Get(&r)
	return
}
