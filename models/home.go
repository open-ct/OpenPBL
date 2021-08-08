package models

func GetLatestProjects(num int) (projects []Project, err error) {
	err = (&Project{}).GetEngine().
		Desc("create_at").
		Limit(num, 0).
		Find(&projects)
	return
}
