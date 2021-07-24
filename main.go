package main

import (
	"OpenPBL/controllers"
	"OpenPBL/models"
	"OpenPBL/routers"
	_ "OpenPBL/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"os"
)

func main() {
	mode := os.Getenv("RUNMODE")
	var err error
	if mode == "prod" {
		err = beego.LoadAppConfig("ini", "conf/app-prod.conf")
	} else if mode == "dev" {
		err = beego.LoadAppConfig("ini", "conf/app-dev.conf")
	} else {
		err = beego.LoadAppConfig("ini", "conf/app-dev.conf")
	}
	if err != nil {
		panic(err)
	}

	fmt.Println(beego.AppConfig.String("casdoorEndpoint"))


	models.InitAdapter()
	controllers.InitCasdoor()

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders: 	  []string{"Origin"},
		ExposeHeaders:	  []string{"Content-Length"},
		AllowCredentials: true,
	}))
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.SetStaticPath("/static", "openpbl-landing/build/static")
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.InsertFilter("/", beego.BeforeRouter, routers.TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, routers.TransparentStatic)

	beego.BConfig.WebConfig.Session.SessionName = "openct_session_id"
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600 * 24 * 365

	beego.Run()
}
