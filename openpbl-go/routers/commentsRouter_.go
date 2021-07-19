package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["openpbl-go/controllers:ObjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ObjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ObjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:objectId",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ObjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:objectId",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ObjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:objectId",
            AllowHTTPMethods: []string{"delete"},
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
            Method: "GetProjectForStudent",
            Router: "/student/:id",
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

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "GetProjectForTeacher",
            Router: "/teacher/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectListController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectListController"],
        beego.ControllerComments{
            Method: "GetStudentPublicProjects",
            Router: "/public/student/:sid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectListController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectListController"],
        beego.ControllerComments{
            Method: "GetTeacherPublicProjects",
            Router: "/public/teacher/:tid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectListController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectListController"],
        beego.ControllerComments{
            Method: "GetProjectListBySid",
            Router: "/student/:type/:sid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:ProjectListController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:ProjectListController"],
        beego.ControllerComments{
            Method: "GetProjectListByTid",
            Router: "/teacher/:type/:tid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:StudentController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:StudentController"],
        beego.ControllerComments{
            Method: "CreateStudent",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:StudentController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:StudentController"],
        beego.ControllerComments{
            Method: "GetStudent",
            Router: "/:sid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:StudentController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:StudentController"],
        beego.ControllerComments{
            Method: "CheckStudentEmail",
            Router: "/checkout/:email",
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
            Router: "/learn",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["openpbl-go/controllers:TeacherController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:TeacherController"],
        beego.ControllerComments{
            Method: "CreateTeacher",
            Router: "/",
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

    beego.GlobalControllerRouter["openpbl-go/controllers:TeacherController"] = append(beego.GlobalControllerRouter["openpbl-go/controllers:TeacherController"],
        beego.ControllerComments{
            Method: "CheckTeacherEmail",
            Router: "/teacher/checkout/:email",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
