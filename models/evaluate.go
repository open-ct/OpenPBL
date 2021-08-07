package models

type Evaluate struct {
	Chapter        `xorm:"extends"`
	SectionOutline `xorm:"extends"`
}


func GetProjectEvaluate(pid string) (evaluate []Evaluate, err error) {
	var c []Chapter
	err = (&Chapter{}).GetEngine().
		Where("project_id = ?", pid).
		Asc("chapter_number").
		Find(&c)
	evaluate = make([]Evaluate, len(c))
	var outline SectionOutline
	for i:=0; i< len(c); i++ {
		evaluate[i].Chapter = c[i]
		err = (&Section{}).GetEngine().
			Where("chapter_id = ?", c[i].Id).
			Find(&outline)
		evaluate[i].SectionOutline = outline
	}
	return
}