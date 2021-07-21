package db

import (
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
		driver := beego.AppConfig.String("driverName")
		dataSource := beego.AppConfig.String("dataSourceName")
		db := beego.AppConfig.String("dbName")

		engine, err := xorm.NewEngine(driver, dataSource + db)

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
