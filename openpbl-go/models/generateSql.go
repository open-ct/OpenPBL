package models

import (
	"bytes"
	"strings"
)

func getSubjectExistSql(subject string) (sql string) {
	if subject == "" {
		return ""
	} else {
		s := strings.Split(subject, ",")
		sql := `
			and exists (
				select project_subject.project_id from project_subject where project_subject.project_id = project.id and (`
		var buf bytes.Buffer
		buf.WriteString(sql)
		n := len(s)
		for i:=0; i<n; i++ {
			buf.WriteString(" project_subject.subject = '")
			buf.WriteString(s[i])
			buf.WriteString("' ")
			if i < n-1 {
				buf.WriteString(" or ")
			}
		}
		buf.WriteString(" ) ) ")
		return buf.String()
	}
}
func getSkillExistSql(skill string) (sql string) {
	if skill == "" {
		return ""
	} else {
		s := strings.Split(skill, ",")
		sql := `
			and exists (
				select project_skill.project_id from project_skill where project_skill.project_id = project.id and (`
		var buf bytes.Buffer
		buf.WriteString(sql)
		n := len(s)
		for i:=0; i<n; i++ {
			buf.WriteString(" project_skill.skill = '")
			buf.WriteString(s[i])
			buf.WriteString("' ")
			if i < n-1 {
				buf.WriteString(" or ")
			}
		}
		buf.WriteString(" ) ) ")
		return buf.String()
	}
}
func getTextSql(text string) (sql string) {
	if text == "" {
		return ""
	} else {
		var buf bytes.Buffer
		buf.WriteString(" and project.project_title like '%")
		buf.WriteString(text)
		buf.WriteString("%'")
		return buf.String()
	}
}
