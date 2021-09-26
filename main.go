// Copyright 2021 The OpenPBL Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
