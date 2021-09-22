// Package routers
// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"OpenPBL/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// auth
	beego.Router("/api/auth/login", &controllers.AuthController{}, "POST:Login")
	beego.Router("/api/auth/logout", &controllers.AuthController{}, "POST:Logout")
	beego.Router("/api/auth/account", &controllers.AuthController{}, "GET:GetAccount")

	// project
	beego.Router("/api/project", &controllers.ProjectController{}, "POST:CreateProject")
	beego.Router("/api/project/:id", &controllers.ProjectController{}, "GET:GetProjectDetail")
	beego.Router("/api/project/:id", &controllers.ProjectController{}, "POST:UpdateProject")
	beego.Router("/api/project/:id/weight", &controllers.ProjectController{}, "POST:UpdateProjectWeight")
	beego.Router("/api/project/:id/publish", &controllers.ProjectController{}, "POST:PublishProject")
	beego.Router("/api/project/:id/close", &controllers.ProjectController{}, "POST:CloseProject")
	beego.Router("/api/project/:id/delete", &controllers.ProjectController{}, "POST:DeleteProject")

	beego.Router("/api/project/:projectId/students", &controllers.ProjectController{}, "GET:GetProjectStudents")
	beego.Router("/api/project/:projectId/remove/:studentId", &controllers.ProjectController{}, "POST:RemoveStudent")
	beego.Router("/api/project/:projectId/favourite/add", &controllers.ProjectController{}, "POST:AddFavouriteProject")
	beego.Router("/api/project/:projectId/favourite/remove", &controllers.ProjectController{}, "POST:RemoveFavouriteProject")

	beego.Router("/api/project/:id/subjects-skills", &controllers.ProjectController{}, "GET:GetProjectSubjectsAndSkills")
	beego.Router("/api/project/:id/view", &controllers.ProjectController{}, "POST:ViewProject")
	beego.Router("/api/project/:id/clone", &controllers.ProjectController{}, "POST:CloneProject")

	// chapter
	beego.Router("/api/project/:id/chapters", &controllers.ProjectController{}, "GET:GetProjectChapters")
	beego.Router("/api/project/:id/chapter", &controllers.ProjectController{}, "POST:CreateProjectChapter")
	beego.Router("/api/project/:projectId/chapter/:chapterId", &controllers.ProjectController{}, "POST:UpdateProjectChapter")
	beego.Router("/api/project/:projectId/chapter/:chapterId/delete", &controllers.ProjectController{}, "POST:DeleteProjectChapter")
	beego.Router("/api/project/:projectId/chapters/exchange", &controllers.ProjectController{}, "POST:ExchangeProjectChapter")

	// section
	beego.Router("/api/project/:projectId/section/:sectionId", &controllers.ProjectController{}, "GET:GetSectionDetail")
	beego.Router("/api/project/:projectId/section/:sectionId/files", &controllers.ProjectController{}, "GET:GetSectionFiles")
	beego.Router("/api/project/:projectId/chapter/:chapterId/sections", &controllers.ProjectController{}, "GET:GetChapterSections")
	beego.Router("/api/project/:projectId/chapter/:chapterId/section", &controllers.ProjectController{}, "POST:CreateChapterSection")
	beego.Router("/api/project/:projectId/chapter/:chapterId/section/:sectionId", &controllers.ProjectController{}, "POST:UpdateChapterSection")
	beego.Router("/api/project/:projectId/chapter/:chapterId/section/:sectionId/delete", &controllers.ProjectController{}, "POST:DeleteChapterSection")
	beego.Router("/api/project/:projectId/chapter/:chapterId/sections/exchange", &controllers.ProjectController{}, "POST:ExchangeChapterSection")
	beego.Router("/api/project/:projectId/sections-minute", &controllers.ProjectController{}, "POST:UpdateSectionsMinute")
	beego.Router("/api/project/:projectId/section/:sectionId/file", &controllers.ProjectController{}, "POST:UploadSectionFile")
	beego.Router("/api/project/:projectId/section/:sectionId/file/:fileId/update", &controllers.ProjectController{}, "POST:UpdateSectionFile")
	beego.Router("/api/project/:projectId/section/:sectionId/file/:fileId/delete", &controllers.ProjectController{}, "POST:DeleteSectionFile")

	// survey
	beego.Router("/api/project/:projectId/task/:taskId/survey", &controllers.ProjectController{}, "GET:GetSurveyDetailByTaskId")
	beego.Router("/api/project/:projectId/task/:taskId/survey/:surveyId/question", &controllers.ProjectController{}, "POST:CreateQuestion")
	beego.Router("/api/project/:projectId/task/:taskId/survey/:surveyId/question/:questionId", &controllers.ProjectController{}, "POST:UpdateQuestion")
	beego.Router("/api/project/:projectId/task/:taskId/survey/:surveyId/questions/exchange", &controllers.ProjectController{}, "POST:ExchangeQuestion")
	beego.Router("/api/project/:projectId/task/:taskId/survey/:surveyId/question/:questionId/delete", &controllers.ProjectController{}, "POST:DeleteQuestion")

	// comments
	beego.Router("/api/project/:id/comments", &controllers.ProjectController{}, "GET:GetProjectComments")
	beego.Router("/api/project/:id/comment", &controllers.ProjectController{}, "POST:CreateProjectComment")
	beego.Router("/api/project/:id/comment/:commentId/delete", &controllers.ProjectController{}, "POST:DeleteProjectComment")

	// resource
	beego.Router("/api/project/:projectId/resource/:resourceId", &controllers.ProjectController{}, "GET:GetResource")
	beego.Router("/api/project/:projectId/resource", &controllers.ProjectController{}, "POST:CreateResource")
	beego.Router("/api/project/:projectId/resource/:resourceId", &controllers.ProjectController{}, "POST:UpdateResource")
	beego.Router("/api/project/:projectId/resource/:resourceId/content", &controllers.ProjectController{}, "POST:UpdateResourceContent")

	// submit
	beego.Router("/api/project/:projectId/task/:taskId/submit", &controllers.ProjectController{}, "POST:CreateSubmit")
	beego.Router("/api/project/:projectId/task/:taskId/submit/:submitId/file", &controllers.ProjectController{}, "POST:UploadSubmitFile")
	beego.Router("/api/project/:projectId/task/:taskId/submit/:submitId/file/:fileId/update", &controllers.ProjectController{}, "POST:UpdateSubmitFile")
	beego.Router("/api/project/:projectId/task/:taskId/submit/:submitId/file/:fileId/delete", &controllers.ProjectController{}, "POST:DeleteSubmitFile")
	beego.Router("/api/project/:projectId/task/:taskId/submit/:submitId", &controllers.ProjectController{}, "POST:UpdateSubmit")
	beego.Router("/api/project/:projectId/task/:taskId/submit/:submitId/files", &controllers.ProjectController{}, "GET:GetSubmitFiles")

	// task
	beego.Router("/api/project/:projectId/section/:sectionId/tasks", &controllers.ProjectController{}, "GET:GetSectionTasksDetail")
	beego.Router("/api/project/:projectId/tasks", &controllers.ProjectController{}, "GET:GetProjectTasks")
	beego.Router("/api/project/:projectId/tasks-detail", &controllers.ProjectController{}, "GET:GetProjectTasksDetail")
	beego.Router("/api/project/:projectId/task", &controllers.ProjectController{}, "POST:CreateTask")
	beego.Router("/api/project/:projectId/task/:taskId", &controllers.ProjectController{}, "POST:UpdateTask")
	beego.Router("/api/project/:projectId/task/:taskId/delete", &controllers.ProjectController{}, "POST:DeleteTask")
	beego.Router("/api/project/:projectId/tasks/exchange", &controllers.ProjectController{}, "POST:ExchangeTask")

	// message
	beego.Router("/api/message", &controllers.MessageController{}, "GET:GetUserMessages")
	beego.Router("/api/message/:messageId/read", &controllers.MessageController{}, "POST:ReadUserMessage")
	beego.Router("/api/message/:messageId/delete", &controllers.MessageController{}, "POST:DeleteUserMessage")
	beego.Router("/api/message/read-all", &controllers.MessageController{}, "POST:ReadAllUserMessage")

	// project-list
	beego.Router("/api/project-list/:projectType", &controllers.ProjectListController{}, "GET:GetUserProjectList")


	// home
	beego.Router("/api/home/projects/latest", &controllers.HomeController{}, "GET:GetLatestProjects")


	// student
	beego.Router("/api/student/learn/:projectId", &controllers.StudentController{}, "POST:LearnProject")
	beego.Router("/api/student/exit/:projectId", &controllers.StudentController{}, "POST:ExitProject")
	beego.Router("/api/student/learn/:projectId", &controllers.StudentController{}, "POST:LearnProject")
	beego.Router("/api/student/project/:projectId/section/:sectionId", &controllers.StudentController{}, "GET:GetLearnSection")
	beego.Router("/api/student/project/:projectId/section/:sectionId", &controllers.StudentController{}, "POST:UpdateLearnSection")
	beego.Router("/api/student/last-learn/project/:projectId", &controllers.StudentController{}, "GET:GetLastLearnSection")

}
