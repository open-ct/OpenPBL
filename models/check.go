package models

import (
	"github.com/astaxie/beego/toolbox"
	"time"
)

func StartTask() {
	spec := "0 0 0 * * *"
	t := toolbox.NewTask("check_project", spec, checkProject)
	err := t.Run()
	if err != nil {
		panic(err)
	}
	toolbox.AddTask("check_project", t)
	toolbox.StartTask()
}

func checkProject() error {
	project := make([]Project, 0)
	_ = adapter.Engine.Table("project").
		Where("published = true").
		Where("closed = false").
		Where("timed_end = true").
		Where("end_time < ?", time.Now().Format("2006-01-02 15:04:05")).
		Find(&project)
	for i:=0; i<len(project); i++ {
		p := project[i]
		p.Closed = true
		p.ClosedAt = time.Now()
		_ = UpdateClosed(p)
	}

	type LP struct {
		ProjectTitle  string
		StudentId     string
		EndTime       time.Time
	}
	p := make([]LP, 0)
	_ = adapter.Engine.Table("project").
		Join("INNER", "learn_project", "project.id = learn_project.project_id").
		Where("published = true").
		Where("closed = false").
		Where("timed_end = true").
		Where("end_time < ?", time.Now().AddDate(0, 0, 3).Format("2006-01-02 15:04:05")).
		Find(&p)
	for i:=0; i<len(p); i++ {
		CreateMessage(&Message{
			ReceiverId:   p[i].StudentId,
			MessageType:  "warn",
			MessageTitle: "您参加的课程 “" + p[i].ProjectTitle + "” 将于 " +
				p[i].EndTime.Format("2006-01-02 15:04:05") + " 结束，请注意提交作业",
			Content:      "",
			ReadMessage:  false,
			CreateAt:     time.Now(),
		})
	}

	return nil
}
