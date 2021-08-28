package main

import (
	"OpenPBL/controllers"
	"OpenPBL/models"
	"OpenPBL/routers"
	_ "OpenPBL/routers"
	"OpenPBL/util"
	"flag"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/astaxie/beego/session/redis"
	"log"
)


func main() {
	var mode string
	flag.StringVar(&mode, "RUNMODE", "default", "运行模式")
	flag.Parse()
	var err error
	configPath := util.GetConfigFile(mode)
	err = beego.LoadAppConfig("ini", configPath)
	if err != nil {
		panic(err)
	}
	log.Println("App start with runmode: " + mode)
	log.Println("Load config file: " + configPath)
	models.InitAdapter()
	controllers.InitCasdoor()
	models.StartTask()

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
	beego.SetStaticPath("/static", "web/build/static")
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.InsertFilter("/", beego.BeforeRouter, routers.TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, routers.TransparentStatic)

	beego.BConfig.WebConfig.Session.SessionName = "openpbl_session_id"
	if beego.AppConfig.String("redisEndpoint") == "" {
		beego.BConfig.WebConfig.Session.SessionProvider = "file"
		beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"
	} else {
		beego.BConfig.WebConfig.Session.SessionProvider = "redis"
		beego.BConfig.WebConfig.Session.SessionProviderConfig = beego.AppConfig.String("redisEndpoint")
	}
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600 * 24 * 30

	beego.Run()
}
