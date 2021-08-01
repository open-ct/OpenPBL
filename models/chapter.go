package models

import "xorm.io/xorm"

type Chapter struct {
	Id                 int64   `json:"id" xorm:"not null pk autoincr"`
	ProjectId          int64   `json:"projectId" xorm:"index"`
	ChapterName        string  `json:"chapterName"`
	ChapterNumber      int     `json:"chapterNumber" xorm:"index"`
}

type Outline struct {
	Chapter                `xorm:"extends"`
	Sections    []Section  `json:"sections" xorm:"extends"`
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
	_, err = p.GetEngine().ID(p.Id).Delete(p)
	// TODO need modify delete all sections belong to this chapter
	return
}
func ExchangeChapters(cid1 string, cid2 string) (err error) {
	_, err = adapter.Engine.
		Exec("update chapter c1 join chapter c2 on (c1.id = ? and c2.id = ?) " +
			"set c1.chapter_number = c2.chapter_number, c2.chapter_number = c1.chapter_number", cid1, cid2)
	return
}

func GetChaptersByPid(pid string) (outline []Outline, err error) {
	var c []Chapter
	err = (&Chapter{}).GetEngine().
		Where("project_id = ?", pid).
		Asc("chapter_number").
		Find(&c)
	outline = make([]Outline, len(c))
	var sections []Section
	for i:=0; i< len(c); i++ {
		outline[i].Chapter = c[i]
		err = (&Section{}).GetEngine().
			Where("chapter_id = ?", c[i].Id).
			Find(&sections)
		outline[i].Sections = sections
	}
	return
}
