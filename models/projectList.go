package models

import (
	"fmt"
)

// ============= student ===================

func GetMyProjectListBySid(sid string, from int, size int,
	subject string, skill string, text string, orderBy string, orderType string, learning bool) (p []ProjectDetail, rows int64, err error) {
	const baseSql = `
		select %s from (
    		select * from project
        		inner join learn_project on (
            		learn_project.student_id = '%s' and
            		learn_project.learning = %v and
					project.closed = false and
            		project.id = learn_project.project_id
        		)
		) as project where true
    	%s %s %s
	`
	const pageSql = " order by %s %s limit %d, %d "
	e1 := getSubjectExistSql(subject)
	e2 := getSkillExistSql(skill)
	e3 := getTextSql(text)

	sql1 := fmt.Sprintf(baseSql, "*", sid, learning, e1, e2, e3) +
		fmt.Sprintf(pageSql, orderBy, orderType, from, size)
	sql2 := fmt.Sprintf(baseSql, "count(*)", sid, learning, e1, e2, e3)
	err = adapter.Engine.
		SQL(sql1).
		Find(&p)
	_, err = adapter.Engine.
		SQL(sql2).
		Get(&rows)
	return
}

func GetPublicProjectListForStudent(sid string, from int, size int,
	subject string, skill string, text string, orderBy string, orderType string, favourite bool) (p []ProjectDetail, rows int64, err error) {
	baseSql := `
		select %s from (
			select * from project where project.published = true and project.closed = false 
			%s %s %s 
		) as p1 left join learn_project on (
			p1.id = learn_project.project_id and learn_project.student_id = '%s'
		)
	`
	if favourite {
		baseSql = `
			select %s from (
				select * from project where exists (
					select favourite.project_id from favourite where favourite.project_id = project.id and favourite.user_id = '`+sid+`'
				)
				%s %s %s 
			) as p1 left join learn_project on (
				p1.id = learn_project.project_id and learn_project.student_id = '%s'
			)
		`
	}

	const pageSql = " order by p1.%s %s limit %d, %d "
	e1 := getSubjectExistSql(subject)
	e2 := getSkillExistSql(skill)
	e3 := getTextSql(text)

	sql1 := fmt.Sprintf(baseSql, "*", e1, e2, e3, sid) +
		fmt.Sprintf(pageSql, orderBy, orderType, from, size)
	sql2 := fmt.Sprintf(baseSql, "count(*)", e1, e2, e3, sid)

	err = adapter.Engine.
		SQL(sql1).
		Find(&p)
	_, err = adapter.Engine.
		SQL(sql2).
		Get(&rows)
	return
}

// ============= teacher ===================

func GetMyProjectListByTid(tid string, from int, size int,
	subject string, skill string, text string, orderBy string, orderType string, published bool, closed bool) (p []ProjectDetail, rows int64, err error) {
	const baseSql = `
		select %s from project where teacher_id = '%s'
    		and published = %v
    		and closed = %v
			%s %s %s 
	`
	const pageSql = " order by %s %s limit %d, %d "
	e1 := getSubjectExistSql(subject)
	e2 := getSkillExistSql(skill)
	e3 := getTextSql(text)

	sql1 := fmt.Sprintf(baseSql, "*", tid, published, closed, e1, e2, e3) +
		fmt.Sprintf(pageSql, orderBy, orderType, from, size)
	sql2 := fmt.Sprintf(baseSql, "count(*)", tid, published, closed, e1, e2, e3)

	err = adapter.Engine.
		SQL(sql1).
		Find(&p)
	_, err = adapter.Engine.
		SQL(sql2).
		Get(&rows)

	return
}

func GetPublicProjectListForTeacher(sid string, from int, size int,
	subject string, skill string, text string, orderBy string, orderType string, favourite bool) (p []ProjectDetail, rows int64, err error) {
	baseSql := "select %s from project where published = true and closed = false %s %s %s %s "
	pageSql := " order by %s %s limit %d, %d "

	if favourite {
		baseSql = `select %s from project where exists (
				select favourite.project_id from favourite where favourite.project_id = project.id and favourite.user_id = '`+sid+`'
			)
			%s %s %s 
		`
	}

	e1 := getSubjectExistSql(subject)
	e2 := getSkillExistSql(skill)
	e3 := getTextSql(text)

	sql1 := fmt.Sprintf(baseSql, "*", e1, e2, e3) +
		fmt.Sprintf(pageSql, orderBy, orderType, from, size)
	sql2 := fmt.Sprintf(baseSql, "count(*)", e1, e2, e3)

	err = adapter.Engine.
		SQL(sql1).
		Find(&p)
	_, err = adapter.Engine.
		SQL(sql2).
		Get(&rows)
	return
}
