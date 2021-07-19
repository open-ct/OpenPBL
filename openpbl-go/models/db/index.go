package db

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"time"
	"xorm.io/xorm"
)

var engine *xorm.Engine
var loadDbOnce sync.Once

func GetEngine() *xorm.Engine {
	loadDbOnce.Do(func() {
		var (
			err      error
			host     string
			port     string
			user     string
			password string
			name     string
			dsn      string
		)
		host = beego.AppConfig.String("db.host")
		port = beego.AppConfig.String("db.port")
		user = beego.AppConfig.String("db.user")
		password = beego.AppConfig.String("db.password")
		name = beego.AppConfig.String("db.name")

		fmt.Sprintln("host=%s, port=%s, user=%s, password=%s, name=%s, ",
			host, port, user, password, name)

		dsn = user + ":" + password + "@tcp(" + host + ":" + port + ")/" + name + "?charset=utf8"
		engine, err = xorm.NewEngine("mysql", dsn)

		if err != nil {
			panic(err.Error())
			return
		}

		engine.DB().SetMaxIdleConns(10)
		engine.DB().SetMaxOpenConns(100)
		engine.DatabaseTZ = time.Local
		engine.TZLocation = time.Local
	})

	return engine
}
