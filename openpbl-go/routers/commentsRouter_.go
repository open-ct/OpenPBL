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
            Method: "CreateProjectChapter",
            Router: "/chapter",
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
            Method: "GetSection",
            Router: "/chapter/section/:sid",
            AllowHTTPMethods: []string{"get"},
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
            Method: "UpdateProject",
            Router: "/info",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "GetProjectOutline",
            Router: "/outline/:pid",
            AllowHTTPMethods: []string{"get"},
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
            Method: "CreateSubmitFile",
            Router: "/submit-files",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "GetSubmitFiles",
            Router: "/submit-files/:pid/:sid",
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

    beego.GlobalControllerRouter["openpbl-go/controllers:TeacherController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:TeacherController"],
        beego.ControllerComments{
            Method: "PublishProject",
            Router: "/publish",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
