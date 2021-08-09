package models

func GetLatestProjects(num int) (projects []Project, err error) {
	err = (&Project{}).GetEngine().
		Where("closed = false").
		Where("published = true").
		Desc("create_at").
		Limit(num, 0).
		Find(&projects)
	return
}
