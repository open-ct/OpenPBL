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
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/auth",
			beego.NSInclude(
				&controllers.AuthController{})),

		beego.NSNamespace("/project",
			beego.NSInclude(
				&controllers.ProjectController{})),

		beego.NSNamespace("/resource",
			beego.NSInclude(
				&controllers.ResourceController{})),

		beego.NSNamespace("/project-list",
			beego.NSInclude(
				&controllers.ProjectListController{})),

		beego.NSNamespace("/student",
			beego.NSInclude(
				&controllers.StudentController{})),

		beego.NSNamespace("/message",
			beego.NSInclude(
				&controllers.MessageController{})),

	)
	beego.AddNamespace(ns)
}
