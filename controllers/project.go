package controllers

import (
	"OpenPBL/models"
	"OpenPBL/util"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/casdoor/casdoor-go-sdk/auth"
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
	uid := user.Username
	var err error
	var project models.ProjectDetail
	if user.Tag == "student" {
		project, err = models.GetProjectByPidForStudent(pid, uid)
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
// @router /:id/publish [post]
func (u *ProjectController) PublishProject() {
	pid, err := u.GetInt64(":id")
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

// CloseProject
// @Title
// @Description
// @Param pid path int true ""
// @Success 200 {object} Response
// @Failure 400
// @Failure 401
// @Failure 403
// @router /:id/close [post]
func (u *ProjectController) CloseProject() {
	pid, err := u.GetInt64(":id")
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
		Id:            pid,
		ClosedAt:      time.Now(),
		Closed:        true,
	}
	err = models.UpdateClosed(p)
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

// DeleteProject
// @Title
// @Description
// @Param pid path int true ""
// @Success 200 {Response}
// @Failure 400
// @Failure 401
// @Failure 403
// @router /:id/delete [post]
func (u *ProjectController) DeleteProject() {
	pid, err := u.GetInt64(":id")
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
		Id:            pid,
	}
	err = p.Delete()
	if err != nil {
		resp = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	} else {
		resp = Response{
			Code: 200,
			Msg:  "删除成功",
		}
	}
	u.Data["json"] = resp
	u.ServeJSON()
}

// RemoveStudent
// @Title
// @Description
// @Param pid path string true ""
// @Success 200 {object} Response
// @Failure 401
// @router /:projectId/remove/:studentId [post]
func (u *ProjectController) RemoveStudent() {
	pid, err := u.GetInt64(":projectId")
	sid := u.GetString(":studentId")
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
	l := &models.LearnProject{
		StudentId: sid,
		ProjectId: pid,
	}
	err = l.Delete()
	if err != nil {
		resp = Response{
			Code: 400,
			Msg:  err.Error(),
		}
		u.Data["json"] = resp
	} else {
		resp = Response{
			Code: 200,
			Msg:  "移除成功",
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

type StudentsResponse struct {
	Code        int                   `json:"code"`
	Students    []models.LearnProject `json:"students"`
	Count       int64                 `json:"count"`
}

// GetProjectStudents
// @Title
// @Description
// @Param from query int true ""
// @Param size query int true ""
// @Success 200 {object} StudentsResponse
// @Failure 403
// @router /:projectId/students [get]
func (p *ProjectController) GetProjectStudents() {
	pid := p.GetString(":projectId")
	from, err := p.GetInt("from")
	if err != nil {
		from = 0
	}
	size, err := p.GetInt("size")
	if err != nil {
		size = 10
	}
	students, rows, err := models.GetProjectStudents(pid, from, size)
	if err != nil {
		p.Data["json"] = StudentsResponse{
			Code:     400,
			Students: nil,
		}
	} else {
		p.Data["json"] = StudentsResponse{
			Code:     200,
			Students: students,
			Count:    rows,
		}
	}
	p.ServeJSON()
}
