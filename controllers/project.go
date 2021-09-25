package controllers

import (
	"OpenPBL/models"
	"OpenPBL/util"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/casdoor/casdoor-go-sdk/auth"
	"strings"
	"time"
)


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
// @Title GetProjectDetail
// @Description Get project detail information
// @Param projectId path string true "project id"
// @Success 200 {object} ProjectResponse
// @router /:projectId [get]
func (p *ProjectController) GetProjectDetail() {
	pid := p.GetString(":projectId")
	user := p.GetSessionUser()
	uid := util.GetUserId(user)
	var project models.ProjectDetail
	var err error
	if util.IsStudent(user) {
		project, err = models.GetProjectByPidForStudent(pid, uid)
	} else if util.IsTeacher(user) {
		project, err = models.GetProjectByPidForTeacher(pid, uid)
	}
	if err != nil {
		p.Data["json"] = ProjectResponse{
			Code:    400,
			Msg:     err.Error(),
		}
	} else {
		p.Data["json"] = ProjectResponse{
			Code:    200,
			Project: project,
		}
	}
	p.ServeJSON()
}

// CreateProject
// @Title CreateProject
// @Description Create a new project
// @Success 200 {object} Response
// @router / [post]
func (p *ProjectController) CreateProject() {
	user := p.GetSessionUser()
	var resp Response
	if !util.IsTeacher(user) {
		resp = Response{
			Code: 403,
			Msg:  "非法用户",
		}
		p.Data["json"] = resp
		p.ServeJSON()
		return
	}
	uid := util.GetUserId(user)
	project := &models.Project{
		Id:                util.NewId(),
		TeacherId:         uid,
		LearnMinuteWeight: 100,
	}
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
// @Title UpdateProject
// @Description create project
// @Param projectId path string true "project id"
// @Success 200 {object} Response
// @router /:projectId [post]
func (p *ProjectController) UpdateProject() {
	user := p.GetSessionUser()
	var resp Response
	if !util.IsTeacher(user) {
		resp = Response{
			Code: 403,
			Msg:  "非法用户",
		}
		p.Data["json"] = resp
		p.ServeJSON()
		return
	}
	uid := util.GetUserId(user)
	pid := p.GetString(":projectId")

	var timeLayoutStr = "2006-01-02 15:04:05 -0700 MST"
	endTime, err := time.Parse(timeLayoutStr, p.GetString("endTime"))
	b := true
	if err != nil {
		b = false
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
		EndTime:          endTime,
		TimedEnd:         b,
	}
	projectSubjects, projectSkills, err := getProjectSubjectsAndSkills(pid, project.Subjects, project.Skills)
	err = project.UpdateInfo(projectSubjects, projectSkills)
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

// UpdateProjectWeight
// @Title UpdateProjectWeight
// @Description Update project weight
// @Param projectId path string true "project id"
// @Success 200 {object} Response
// @router /:projectId/weight [post]
func (p *ProjectController) UpdateProjectWeight() {
	user := p.GetSessionUser()
	var resp Response
	if !util.IsTeacher(user) {
		resp = Response{
			Code: 403,
			Msg:  "非法用户",
		}
		p.Data["json"] = resp
		p.ServeJSON()
		return
	}
	pid := p.GetString(":projectId")
	learnMinuteWeight, err := p.GetInt("learnMinuteWeight")

	if err != nil {
		resp = Response{
			Code: 400,
			Msg: err.Error(),
		}
		p.Data["json"] = resp
		p.ServeJSON()
		return
	}
	project := models.Project{
		Id:                pid,
		LearnMinuteWeight: learnMinuteWeight,
	}
	tasks := make([]models.Task, 0)
	err = json.Unmarshal([]byte(p.GetString("tasks")), &tasks)
	err = models.UpdateWeight(project, tasks)
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
// @Title PublishProject
// @Description
// @Param projectId path string true "The id of the project"
// @Success 200 {Response}
// @router /:projectId/publish [post]
func (p *ProjectController) PublishProject() {
	pid := p.GetString(":projectId")
	var resp Response
	user := p.GetSessionUser()
	if !util.IsTeacher(user) {
		resp = Response{
			Code: 403,
			Msg:  "非法的用户",
		}
		p.Data["json"] = resp
		p.ServeJSON()
		return
	}

	project := models.Project{
		Id:               pid,
		PublishedAt:      time.Now(),
		Published:        true,
	}
	err := models.UpdatePublished(project)
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
	p.Data["json"] = resp
	p.ServeJSON()
}

// CloseProject
// @Title
// @Description
// @Param pid path int true ""
// @Success 200 {object} Response
// @router /:projectId/close [post]
func (p *ProjectController) CloseProject() {
	pid := p.GetString(":projectId")
	var resp Response
	user := p.GetSessionUser()
	if !util.IsTeacher(user) {
		resp = Response{
			Code: 403,
			Msg:  "非法的用户",
		}
		p.Data["json"] = resp
		p.ServeJSON()
		return
	}

	project := models.Project{
		Id:            pid,
		ClosedAt:      time.Now(),
		Closed:        true,
	}
	err := models.UpdateClosed(project)
	if err != nil {
		resp = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	} else {
		resp = Response{
			Code: 200,
			Msg:  "结束成功",
		}
	}
	p.Data["json"] = resp
	p.ServeJSON()
}

// DeleteProject
// @Title
// @Description
// @Param pid path int true ""
// @Success 200 {Response}
// @Failure 400
// @Failure 401
// @Failure 403
// @router /:projectId/delete [post]
func (p *ProjectController) DeleteProject() {
	pid := p.GetString(":projectId")
	var resp Response
	user := p.GetSessionUser()
	if !util.IsTeacher(user) {
		resp = Response{
			Code: 403,
			Msg:  "非法的用户",
		}
		p.Data["json"] = resp
		p.ServeJSON()
		return
	}

	project := models.Project{
		Id:            pid,
	}
	err := project.Delete()
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
	p.Data["json"] = resp
	p.ServeJSON()
}

// RemoveStudent
// @Title
// @Description
// @Param pid path string true ""
// @Success 200 {object} Response
// @Failure 401
// @router /:projectId/remove/:studentId [post]
func (p *ProjectController) RemoveStudent() {
	pid := p.GetString(":projectId")
	sid := p.GetString(":studentId")
	var resp Response
	l := &models.LearnProject{
		StudentId: sid,
		ProjectId: pid,
	}
	err := l.Delete()
	if err != nil {
		resp = Response{
			Code: 400,
			Msg:  err.Error(),
		}
		p.Data["json"] = resp
	} else {
		resp = Response{
			Code: 200,
			Msg:  "移除成功",
		}
	}
	p.Data["json"] = resp
	p.ServeJSON()
}


func getProjectSubjectsAndSkills(pid string, subjects string, skills string) (subjectList []*models.ProjectSubject, skillList []*models.ProjectSkill, err error) {
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

// AddFavouriteProject
// @Title
// @Description
// @Param projectId path string true ""
// @Success 200 {object} Response
// @Failure 401
// @router /:projectId/favourite/add [post]
func (p *ProjectController) AddFavouriteProject() {
	pid := p.GetString(":projectId")
	uid := util.GetUserId(p.GetSessionUser())
	err := models.AddFavourite(uid, pid)
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
			Msg:  "收藏成功",
		}
	}
	p.ServeJSON()
}
// RemoveFavouriteProject
// @Title
// @Description
// @Param projectId path string true "project id"
// @Success 200 {object} Response
// @router /:projectId/favourite/remove [post]
func (p *ProjectController) RemoveFavouriteProject() {
	pid := p.GetString(":projectId")
	uid := util.GetUserId(p.GetSessionUser())
	err := models.RemoveFavourite(uid, pid)
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
			Msg:  "移除成功",
		}
	}
	p.ServeJSON()
}

type SubjectsAndSkillsResponse struct {
	Code       int        `json:"code"`
	Subjects   []string   `json:"subjects"`
	Skills     []string   `json:"skills"`
}

// GetProjectSubjectsAndSkills
// @Title
// @Description
// @Success 200 {object} SubjectsAndSkillsResponse
// @router /:projectId/subjects-skills [get]
func (p *ProjectController) GetProjectSubjectsAndSkills() {
	subjects, err := models.GetSubjects()
	skills, err := models.GetSkills()
	if err != nil {
		p.Data["json"] = SubjectsAndSkillsResponse{
			Code:     400,
		}
	} else {
		p.Data["json"] = SubjectsAndSkillsResponse{
			Code:    200,
			Subjects: subjects,
			Skills:   skills,
		}
	}
	p.ServeJSON()
}

// ViewProject
// @Title
// @Description View a project
// @Param id path string true "project id"
// @Success 200 {object} Response
// @router /:projectId/view [post]
func (p *ProjectController) ViewProject() {
	pid := p.GetString(":projectId")
	err := models.ViewProject(pid)
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
		}
	}
	p.ServeJSON()
}

// CloneProject
// @Title CloneProject
// @Description Clone project
// @Param projectId path string true "The id of the project"
// @Success 200 {object} Response
// @router /:projectId/clone [post]
func (p *ProjectController) CloneProject() {
	pid := p.GetString(":projectId")
	uid := util.GetUserId(p.GetSessionUser())
	err := models.CloneProject(uid, pid)
	if err != nil {
		p.Data["json"] = Response{
			Code: 400,
			Msg:  err.Error(),
		}
	} else {
		p.Data["json"] = Response{
			Code: 200,
			Msg: "复制成功，请到未发布项目中查看",
		}
	}
	p.ServeJSON()
}