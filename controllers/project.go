package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/casdoor/casdoor-go-sdk/auth"
	"OpenPBL/models"
	"OpenPBL/util"
	"strings"
	"time"
)

// ProjectController
// Operations about Projects
type ProjectController struct {
	beego.Controller
}

func (p *ProjectController) GetSessionUser() *auth.Claims {
	s := p.GetSession("user")
	if s == nil {
		return nil
	}
	claims := &auth.Claims{}
	err := util.JsonToStruct(s.(string), claims)
	if err != nil {
		panic(err)
	}
	return claims
}

type ProjectResponse struct {
	Code    int                  `json:"code"`
	Msg     string               `json:"msg"`
	Project models.ProjectDetail `json:"project"`
}

// GetProjectDetail
// @Title
// @Description
// @Param id path string true "project id"
// @Success 200 {object} models.TeacherProject
// @Failure 400
// @router /:id [get]
func (p *ProjectController) GetProjectDetail() {
	pid := p.GetString(":id")
	user := p.GetSessionUser()
	var resp ProjectResponse
	if user == nil {
		resp = ProjectResponse{
			Code:    401,
			Msg:     "请先登录",
		}
		p.Data["json"] = resp
		p.ServeJSON()
		return
	}
	var err error
	var project models.ProjectDetail
	if user.Tag == "student" {
		project, err = models.GetProjectByPidForStudent(pid)
	} else if user.Tag == "teacher" {
		project, err = models.GetProjectByPidForTeacher(pid)
	}
	if err != nil {
		resp = ProjectResponse{
			Code:    400,
			Msg:     err.Error(),
		}
	} else {
		resp = ProjectResponse{
			Code:    200,
			Project: project,
		}
	}
	p.Data["json"] = resp
	p.ServeJSON()
}

// CreateProject
// @Title
// @Description create project
// @Success 200 {object} Response
// @Failure 401
// @Failure 400
// @Failure 403
// @router / [post]
func (p *ProjectController) CreateProject() {
	user := p.GetSessionUser()
	var resp Response
	if user == nil {
		resp = Response{
			Code: 401,
			Msg:  "请先登录",
		}
		p.Data["json"] = resp
		p.ServeJSON()
		return
	}
	if user.Tag != "teacher" {
		resp = Response{
			Code: 403,
			Msg:  "非法用户",
		}
		p.Data["json"] = resp
		p.ServeJSON()
		return
	}
	uid := user.Username
	project := &models.Project{
		TeacherId:        uid,
	}
	fmt.Println(project)
	err := project.Create()
	if err != nil {
		resp = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	} else {
		resp = Response{
			Code: 200,
			Msg:  "创建成功",
			Data: project.Id,
		}
	}
	p.Data["json"] = resp
	p.ServeJSON()
}

// UpdateProject
// @Title
// @Description create project
// @Param body body models.Project true	""
// @Success 200 {int} models.Project.Id
// @Failure 403 body is empty
// @router /:id [post]
func (p *ProjectController) UpdateProject() {
	user := p.GetSessionUser()
	var resp Response
	if user == nil {
		resp = Response{
			Code: 401,
			Msg:  "请先登录",
		}
		p.Data["json"] = resp
		p.ServeJSON()
		return
	}
	if user.Tag != "teacher" {
		resp = Response{
			Code: 403,
			Msg:  "非法用户",
		}
		p.Data["json"] = resp
		p.ServeJSON()
		return
	}
	uid := user.Username
	pid, err := p.GetInt64(":id")
	if err != nil {
		resp = Response{
			Code: 400,
			Msg: err.Error(),
		}
		p.Data["json"] = resp
		p.ServeJSON()
		return
	}
	project := &models.Project{
		Id:               pid,
		Image:            p.GetString("image"),
		ProjectTitle:     p.GetString("projectTitle"),
		ProjectIntroduce: p.GetString("projectIntroduce"),
		ProjectGoal:      p.GetString("projectGoal"),
		TeacherId:        uid,
		Subjects:         p.GetString("subjects"),
		Skills:           p.GetString("skills"),
	}
	projectSubjects, projectSkills, err := getProjectSubjectsAndSkills(pid, project.Subjects, project.Skills)
	err = project.Update(projectSubjects, projectSkills)
	if err != nil {
		resp = Response{
			Code: 400,
			Msg:  err.Error(),
			Data: true,
		}
	} else {
		resp = Response{
			Code: 200,
			Msg:  "更新成功",
		}
	}
	p.Data["json"] = resp
	p.ServeJSON()
}


// PublishProject
// @Title
// @Description
// @Param pid path int true ""
// @Success 200 {Response}
// @Failure 400
// @Failure 401
// @Failure 403
// @router /publish/:pid [post]
func (u *ProjectController) PublishProject() {
	pid, err := u.GetInt64(":pid")
	var resp Response
	user := u.GetSessionUser()
	if user == nil {
		resp = Response{
			Code: 401,
			Msg:  "请先登录",
		}
		u.Data["json"] = resp
		u.ServeJSON()
		return
	}
	if user.Tag != "teacher" {
		resp = Response{
			Code: 403,
			Msg:  "非法的用户",
		}
		u.Data["json"] = resp
		u.ServeJSON()
		return
	}

	p := models.Project{
		Id:               pid,
		PublishedAt:      time.Now(),
		Published:        true,
	}
	err = models.UpdatePublished(p)
	if err != nil {
		resp = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	} else {
		resp = Response{
			Code: 200,
			Msg:  "发布成功",
		}
	}
	u.Data["json"] = resp
	u.ServeJSON()
}

func getProjectSubjectsAndSkills(pid int64, subjects string, skills string) (subjectList []*models.ProjectSubject, skillList []*models.ProjectSkill, err error) {
	var (
		subjectL []string
		skillL   []string
	)

	if subjects == "" {
		subjectL = make([]string, 0)
	} else {
		subjectL = strings.Split(subjects, ",")
	}
	if skills == "" {
		skillL = make([]string, 0)
	} else {
		skillL = strings.Split(skills, ",")
	}
	n1 := len(subjectL)
	n2 := len(skillL)

	subjectList = make([]*models.ProjectSubject, n1)
	skillList = make([]*models.ProjectSkill, n2)
	for i:=0; i<n1; i++ {
		subjectList[i] = &models.ProjectSubject{
			Subject:   subjectL[i],
			ProjectId: pid,
		}
	}
	for i:=0; i<n2; i++ {
		skillList[i] = &models.ProjectSkill{
			Skill:     skillL[i],
			ProjectId: pid,
		}
	}
	return
}