package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["openpbl-go/controllers:AuthController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:AuthController"],
        beego.ControllerComments{
            Method: "GetAccount",
            Router: "/account",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:AuthController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:AuthController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:AuthController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:AuthController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: "/logout",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "CreateProject",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "GetProjectDetail",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "UpdateProject",
            Router: "/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "CreateProjectChapter",
            Router: "/chapter",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "UpdateProjectChapter",
            Router: "/chapter/:cid",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "DeleteProjectChapter",
            Router: "/chapter/delete/:cid",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "ExchangeProjectChapter",
            Router: "/chapter/exchange/:cid1/:cid2",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "CreateChapterSection",
            Router: "/chapter/section",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "UpdateChapterSection",
            Router: "/chapter/section/:sid",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "GetSectionDetail",
            Router: "/chapter/section/:sid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "DeleteChapterSection",
            Router: "/chapter/section/delete/:sid",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "ExchangeChapterSection",
            Router: "/chapter/section/exchange/:sid1/:sid2",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "GetChapterSections",
            Router: "/chapter/sections/:cid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "GetProjectChapters",
            Router: "/chapters/:pid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "PublishProject",
            Router: "/publish/:pid",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "GetProjectStudents",
            Router: "/students/:pid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "CreateSubmit",
            Router: "/submit",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "CreateTask",
            Router: "/task",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "UpdateTask",
            Router: "/task/:tid",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "DeleteTask",
            Router: "/task/delete/:tid",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "ExchangeTask",
            Router: "/task/exchange/:tid1/:tid2",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "CreateSurvey",
            Router: "/task/survey/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "UpdateSurvey",
            Router: "/task/survey/:sid",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "GetSurveyDetailByTaskId",
            Router: "/task/survey/:tid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "CreateQuestion",
            Router: "/task/survey/question",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "UpdateQuestion",
            Router: "/task/survey/question/:qid",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "DeleteQuestion",
            Router: "/task/survey/question/delete/:qid",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "ExchangeQuestion",
            Router: "/task/survey/question/exchange/:id1/:id2",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "GetSectionTasks",
            Router: "/tasks/:sid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectListController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectListController"],
        beego.ControllerComments{
            Method: "GetUserProjectList",
            Router: "/:projectType",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ResourceController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ResourceController"],
        beego.ControllerComments{
            Method: "CreateResource",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ResourceController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ResourceController"],
        beego.ControllerComments{
            Method: "GetResource",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ResourceController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ResourceController"],
        beego.ControllerComments{
            Method: "UpdateResource",
            Router: "/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ResourceController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ResourceController"],
        beego.ControllerComments{
            Method: "UpdateResourceContent",
            Router: "/content/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:StudentController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:StudentController"],
        beego.ControllerComments{
            Method: "FinishedProject",
            Router: "/finished",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:StudentController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:StudentController"],
        beego.ControllerComments{
            Method: "LearnProject",
            Router: "/learn/:pid",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
