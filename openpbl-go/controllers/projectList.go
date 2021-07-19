package controllers

import (
	"github.com/astaxie/beego"
	"openpbl-go/models"
	"strconv"
)

// ProjectListController
// Operations about Projects
type ProjectListController struct {
	beego.Controller
}

type ProjectList struct {
	Count    int64            `json:"count"`
	Projects []models.Project `json:"projects"`
}

// GetProjectListByTid
// @Title
// @Description get project list by teacher id
// @Param type path string true "editing published ..."
// @Param tid path string true "teacher id"
// @Param from query string false "from"
// @Param size query string false "size"
// @Param subject query string false ""
// @Param skill query string false ""
// @Param text query string false ""
// @Success 200 {object} ProjectList
// @Failure 403 :id is empty
// @router /teacher/:type/:tid [get]
func (pl *ProjectListController) GetProjectListByTid() {
	t := pl.GetString(":type")
	tid := pl.GetString(":tid")
	from, err := strconv.Atoi(pl.GetString("from"))
	if err != nil {
		from = 0
	}
	size, err := strconv.Atoi(pl.GetString("size"))
	if err != nil {
		size = 10
	}
	skill := pl.GetString("skill")
	subject := pl.GetString("subject")
	text := pl.GetString("text")

	var projects []models.Project
	var count    int64
	if t == "editing" {
		projects, count, err = models.GetMyProjectListByTid(tid, from, size, subject, skill, text, "create_at", "desc", false, false)
	} else if t == "published" {
		projects, count, err = models.GetMyProjectListByTid(tid, from, size, subject, skill, text, "create_at", "desc", true, false)
	} else {
		projects, count, err = models.GetMyProjectListByTid(tid, from, size, subject, skill, text, "create_at", "desc", true, true)
	}
	if err != nil {
		pl.Data["json"] = map[string]string{"err": err.Error()}
	}

	data := &ProjectList{Projects: projects, Count: count}
	pl.Data["json"] = data

	pl.ServeJSON()
}


type StudentProjectList struct {
	Count    int64                   `json:"count"`
	Projects []models.StudentProject `json:"projects"`
}

// GetProjectListBySid
// @Title Get
// @Description get project list by student id
// @Param type path string true "learning ..."
// @Param sid path string true ""
// @Param from query string false "from"
// @Param size query string false "size"
// @Param subject query string false ""
// @Param skill query string false ""
// @Param text query string false ""
// @Success 200 {object} StudentProjectList
// @Failure 403 :id is empty
// @router /student/:type/:sid [get]
func (pl *ProjectListController) GetProjectListBySid() {
	t := pl.GetString(":type")
	sid := pl.GetString(":sid")
	from, err := strconv.Atoi(pl.GetString("from"))
	if err != nil {
		from = 0
	}
	size, err := strconv.Atoi(pl.GetString("size"))
	if err != nil {
		size = 10
	}
	skill := pl.GetString("skill")
	subject := pl.GetString("subject")
	text := pl.GetString("text")
	b := false
	if t == "learning" {
		b = true
	}
	var projects []models.StudentProject
	var count    int64
	projects, count, err = models.GetMyProjectListBySid(sid, from, size, subject, skill, text, "create_at", "desc", b)

	if err != nil {
		pl.Data["json"] = map[string]string{"error": err.Error()}
	}

	data := &StudentProjectList{Projects: projects, Count: count}
	pl.Data["json"] = data

	pl.ServeJSON()
}

// GetTeacherPublicProjects
// @Title Get
// @Description
// @Param sid path string true ""
// @Param orderBy path string true "create_at join_num"
// @Param from query string false "from"
// @Param size query string false "size"
// @Param subject query string false ""
// @Param skill query string false ""
// @Param text query string false ""
// @Success 200 {object} ProjectList
// @Failure 403
// @router /public/teacher/:tid [get]
func (pl *ProjectListController) GetTeacherPublicProjects() {
	orderBy := pl.GetString("orderBy")
	tid := pl.GetString(":tid")
	from, err := strconv.Atoi(pl.GetString("from"))
	if err != nil {
		from = 0
	}
	size, err := strconv.Atoi(pl.GetString("size"))
	if err != nil {
		size = 10
	}
	skill := pl.GetString("skill")
	subject := pl.GetString("subject")
	text := pl.GetString("text")
	var projects []models.Project
	var count    int64
	projects, count, err = models.GetPublicProjectListForTeacher(tid, from, size, subject, skill, text, orderBy, "desc")

	if err != nil {
		pl.Data["json"] = map[string]string{"error": err.Error()}
	}

	data := &ProjectList{Projects: projects, Count: count}
	pl.Data["json"] = data

	pl.ServeJSON()
}

// GetStudentPublicProjects
// @Title
// @Description
// @Param sid path string true "student id"
// @Param orderBy query string true "create_at join_num"
// @Param from query string false "from"
// @Param size query string false "size"
// @Param skill query string true ""
// @Param subject path string true "数学,英语"
// @Param text path string false ""
// @Success 200 {object} StudentProjectList
// @Failure 403
// @router /public/student/:sid [get]
func (pl *ProjectListController) GetStudentPublicProjects() {
	sid := pl.GetString(":sid")
	skill := pl.GetString("skill")
	subject := pl.GetString("subject")
	text := pl.GetString("text")
	orderBy := pl.GetString("orderBy")
	from, err := strconv.Atoi(pl.GetString("from"))
	if err != nil {
		from = 0
	}
	size, err := strconv.Atoi(pl.GetString("size"))
	if err != nil {
		size = 10
	}
	var projects []models.StudentProject
	var count    int64
	projects, count, err = models.GetPublicProjectListForStudent(sid, from, size, subject, skill, text, orderBy, "desc")

	if err != nil {
		pl.Data["json"] = map[string]string{"error": err.Error()}
	}

	data := &StudentProjectList{Projects: projects, Count: count}
	pl.Data["json"] = data

	pl.ServeJSON()
}
