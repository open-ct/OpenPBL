package models

import (
	"openpbl-go/models/db"
	"time"
	"xorm.io/xorm"
)

type Project struct {
	Id                 int64     `json:"id" xorm:"not null pk autoincr"`
	Image              string    `json:"image" xorm:"longtext"`
	ProjectTitle       string    `json:"projectTitle"`
	ProjectIntroduce   string    `json:"projectIntroduce"`
	ProjectGoal        string     `json:"projectGoal"`

	TeacherId          int64     `json:"teacherId" xorm:"index"`

	Subjects           string    `json:"subjects" xorm:"default ''"`
	Skills             string    `json:"skills" xorm:"default ''"`

	Closed             bool      `json:"closed" xorm:"default false index"`
	CreateAt           time.Time `json:"createAt" xorm:"created"`
	PublishedAt        time.Time `json:"publishedAt"`
	ClosedAt           time.Time `json:"closedAt"`
	ReadNum            int64     `json:"readNum" xorm:"default 0"`
	JoinNum            int64     `json:"joinNum" xorm:"default 0"`

	Published          bool      `json:"published" xorm:"default false index"`
}

type Outline struct {
	Chapter                `xorm:"extends"`
	Section    []Section   `json:"section" xorm:"extends"`
}

type Chapter struct {
	Id                 int64   `json:"id" xorm:"not null pk autoincr"`
	ProjectId          int64   `json:"projectId" xorm:"index"`
	ChapterName        string  `json:"chapterName"`
	ChapterNumber      int     `json:"chapterNumber" xorm:"index"`
}

type Section struct {
	Id                 int64   `json:"id" xorm:"not null pk autoincr"`
	ChapterId          int64   `json:"chapterId" xorm:"index"`
	SectionName        string  `json:"sectionName"`
	SectionNumber      int     `json:"sectionNumber" xorm:"index"`
}

type Resource struct {
	Id               int64   `json:"id" xorm:"not null pk autoincr"`
	FilePath         string  `json:"filePath" xorm:"index"`
	FileType         string  `json:"fileType" xorm:"index"`
	FileName         string  `json:"fileName"`
	SectionId        int64   `json:"sectionId" xorm:"index"`
	FileNumber      int      `json:"fileNumber" xorm:"index"`
}

type SubmitFile struct {
	Id              int64     `json:"id" xorm:"not null pk autoincr"`

	ProjectId       int64     `json:"projectId" xorm:"index"`
	StudentId       int64     `json:"studentId" xorm:"index"`

	SubmitIntroduce string    `json:"submitIntroduce"`

	FilePath        string    `json:"filePath"`
	FileName        string    `json:"fileName"`

	CreateAt        time.Time `json:"createAt" xorm:"created"`
}

type ProjectSkill struct {
	Skill          string     `json:"skill" xorm:"not null pk"`
	ProjectId      int64      `json:"projectId" xorm:"not null pk"`
}

type ProjectSubject struct {
	Subject        string     `json:"subject" xorm:"not null pk"`
	ProjectId      int64      `json:"projectId" xorm:"not null pk"`
}

func (p *Project) GetEngine() *xorm.Session {
	return db.GetEngine().Table(p)
}
func (p *Chapter) GetEngine() *xorm.Session {
	return db.GetEngine().Table(p)
}
func (p *Section) GetEngine() *xorm.Session {
	return db.GetEngine().Table(p)
}
func (p *SubmitFile) GetEngine() *xorm.Session {
	return db.GetEngine().Table(p)
}
func (p *ProjectSkill) GetEngine() *xorm.Session {
	return db.GetEngine().Table(p)
}
func (p *ProjectSubject) GetEngine() *xorm.Session {
	return db.GetEngine().Table(p)
}

func GetProjectByPidForTeacher(pid string) (pd ProjectDetail, err error) {
	_, err = (&Project{}).GetEngine().
		Join("INNER", "teacher", "teacher.id = project.teacher_id").
		Where("project.id = ?", pid).
		Get(&pd)
	return
}

func GetProjectByPidForStudent(pid string) (pd ProjectDetail, err error) {
	_, err = db.GetEngine().
		SQL("select * from project " +
			"inner join teacher t on project.teacher_id = t.id " +
			"left join learn_project lp on project.id = lp.project_id " +
			"where project.id = ?", pid).
		Get(&pd)
	return
}


func (p *Project) Create() (err error) {
	_, err = p.GetEngine().Insert(p)
	return
}

func (p *ProjectSkill) Create() (err error) {
	_, err = p.GetEngine().Insert(p)
	return
}
func (p *ProjectSubject) Create() (err error) {
	_, err = p.GetEngine().Insert(p)
	return
}

func (p *Project) Update(subjects []*ProjectSubject, skills []*ProjectSkill) (err error) {
	_, err = p.GetEngine().ID(p.Id).Update(p)
	n1 := len(subjects)
	for i:=0; i<n1; i++ {
		err = subjects[i].Create()
	}
	n2 := len(skills)
	for i:=0; i<n2; i++ {
		err = skills[i].Create()
	}
	return
}

func (p *Chapter) Create() (err error) {
	_, err = p.GetEngine().Insert(p)
	return
}
func (p *Chapter) Update() (err error) {
	_, err = p.GetEngine().Where("id = ", p.Id).Update(p)
	return
}

func (p *Section) Create() (err error) {
	_, err = p.GetEngine().Insert(p)
	return
}
func (p *Section) Update() (err error) {
	_, err = p.GetEngine().Where("id = ", p.Id).Update(p)
	return
}

func (p *SubmitFile) Create() (err error) {
	_, err = p.GetEngine().Insert(p)
	return
}

func GetSubmitFiles(sid string, pid string) (f[] SubmitFile, err error) {
	err = (&SubmitFile{}).GetEngine().
		Where("project_id = ? and student_id = ?", pid, sid).
		Find(&f)
	return
}

func GetOutlineByPid(pid string) (c []Outline, err error) {
	err = db.GetEngine().
		SQL("select * from chapter left join section s on chapter.id = s.chapter_id where chapter.project_id = 1").
		// Where("project_id = ?", pid).
		// Asc("chapter_number").
		Find(&c)
	return
}

func GetChaptersByPid(pid string) (c []Chapter, err error) {
	err = (&Chapter{}).GetEngine().
		Where("project_id = ?", pid).
		Asc("chapter_number").
		Find(&c)
	return
}

func GetSectionsByCid(cid string) (s []Section, err error) {
	err = (&Section{}).GetEngine().
		Where("chapter_id = ?", cid).
		Asc("section_number").
		Find(&s)
	return
}


func UpdatePublished(p Project) (err error) {
	_, err = (&Project{}).GetEngine().Where("id = ?", p.Id).Cols("published", "publish_at").Update(p)
	return
}

func GetSectionById(sid string) (s Section, err error) {
	_, err = (&Section{}).GetEngine().
		Where("id = ?", sid).
		Get(&s)
	return
}


func GetProjectStudents(pid string, from int, size int) (s []StudentInfo, err error) {
	err = db.GetEngine().
		SQL("select * from student " +
			"inner join " +
			"( select * from learn_project where project_id = ?) as l " +
			"on student.id = l.student_id limit ?, ?", pid, from, size).
		Find(&s)
	return
}
func CountProjectStudents(pid string, from int, size int) (rows int64, err error) {
	_, err = db.GetEngine().
		SQL("select count(*) from student " +
			"inner join " +
			"( select * from learn_project where project_id = ?) as l " +
			"on student.id = l.student_id", pid).
		Get(&rows)
	return
}