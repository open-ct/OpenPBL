package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["OpenPBL/controllers:AuthController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:AuthController"],
        beego.ControllerComments{
            Method: "GetAccount",
            Router: "/account",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:AuthController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:AuthController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:AuthController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:AuthController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: "/logout",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "CreateProject",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "UpdateProject",
            Router: "/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "GetProjectDetail",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "CreateProjectChapter",
            Router: "/:id/chapter",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "GetProjectChapters",
            Router: "/:id/chapters",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "CloseProject",
            Router: "/:id/close",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "DeleteProject",
            Router: "/:id/delete",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "PublishProject",
            Router: "/:id/publish",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "UpdateProjectWeight",
            Router: "/:id/weight",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "UpdateProjectChapter",
            Router: "/:projectId/chapter/:chapterId",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "DeleteProjectChapter",
            Router: "/:projectId/chapter/:chapterId/delete",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "CreateChapterSection",
            Router: "/:projectId/chapter/:chapterId/section",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "UpdateChapterSection",
            Router: "/:projectId/chapter/:chapterId/section/:sectionId",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "DeleteChapterSection",
            Router: "/:projectId/chapter/:chapterId/section/:sectionId/delete",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "GetChapterSections",
            Router: "/:projectId/chapter/:chapterId/sections",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "ExchangeChapterSection",
            Router: "/:projectId/chapter/:chapterId/sections/exchange",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "ExchangeProjectChapter",
            Router: "/:projectId/chapters/exchange",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "RemoveStudent",
            Router: "/:projectId/remove/:studentId",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "GetSectionDetail",
            Router: "/:projectId/section/:sectionId",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "GetSectionTasks",
            Router: "/:projectId/section/:sectionId/tasks",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "GetProjectStudents",
            Router: "/:projectId/students",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "CreateTask",
            Router: "/:projectId/task",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "UpdateTask",
            Router: "/:projectId/task/:taskId",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "DeleteTask",
            Router: "/:projectId/task/:taskId/delete",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "CreateSubmit",
            Router: "/:projectId/task/:taskId/submit",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "UpdateSubmit",
            Router: "/:projectId/task/:taskId/submit/:submitId",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "GetSurveyDetailByTaskId",
            Router: "/:projectId/task/:taskId/survey",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "CreateSurvey",
            Router: "/:projectId/task/:taskId/survey",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "UpdateSurvey",
            Router: "/:projectId/task/:taskId/survey/:sid",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "CreateQuestion",
            Router: "/:projectId/task/:taskId/survey/:surveyId/question",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "UpdateQuestion",
            Router: "/:projectId/task/:taskId/survey/:surveyId/question/:questionId",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "DeleteQuestion",
            Router: "/:projectId/task/:taskId/survey/:surveyId/question/:questionId/delete",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "ExchangeQuestion",
            Router: "/:projectId/task/:taskId/survey/:surveyId/questions/exchange",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "GetProjectTasks",
            Router: "/:projectId/tasks",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "ExchangeTask",
            Router: "/:projectId/tasks/exchange",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "SaveTaskWeight",
            Router: "/:projectId/tasks/exchange",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ProjectListController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ProjectListController"],
        beego.ControllerComments{
            Method: "GetUserProjectList",
            Router: "/:projectType",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ResourceController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ResourceController"],
        beego.ControllerComments{
            Method: "CreateResource",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ResourceController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ResourceController"],
        beego.ControllerComments{
            Method: "GetResource",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ResourceController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ResourceController"],
        beego.ControllerComments{
            Method: "UpdateResource",
            Router: "/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:ResourceController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:ResourceController"],
        beego.ControllerComments{
            Method: "UpdateResourceContent",
            Router: "/:id/content",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:StudentController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:StudentController"],
        beego.ControllerComments{
            Method: "ExitProject",
            Router: "/exit/:projectId",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:StudentController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:StudentController"],
        beego.ControllerComments{
            Method: "FinishedProject",
            Router: "/finished",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:StudentController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:StudentController"],
        beego.ControllerComments{
            Method: "LearnProject",
            Router: "/learn/:projectId",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:StudentController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:StudentController"],
        beego.ControllerComments{
            Method: "CreateLearnSection",
            Router: "/section/:sectionId",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenPBL/controllers:StudentController"] = append(beego.GlobalControllerRouter["OpenPBL/controllers:StudentController"],
        beego.ControllerComments{
            Method: "UpdateLearnSection",
            Router: "/section/:sectionId",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
