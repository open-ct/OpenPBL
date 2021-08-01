package models

import (
	"fmt"
	"runtime"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var adapter *Adapter

func InitAdapter() {
	adapter = NewAdapter(beego.AppConfig.String("driverName"), beego.AppConfig.String("dataSourceName"), beego.AppConfig.String("dbName"))
	adapter.createTable()
}

// Adapter represents the MySQL adapter for policy storage.
type Adapter struct {
	driverName     string
	dataSourceName string
	dbName         string
	Engine         *xorm.Engine
}

// finalizer is the destructor for Adapter.
func finalizer(a *Adapter) {
	err := a.Engine.Close()
	if err != nil {
		panic(err)
	}
}

// NewAdapter is the constructor for Adapter.
func NewAdapter(driverName string, dataSourceName string, dbName string) *Adapter {
	a := &Adapter{}
	a.driverName = driverName
	a.dataSourceName = dataSourceName
	a.dbName = dbName

	// Open the DB, create it if not existed.
	a.open()

	// Call the destructor when the object is released.
	runtime.SetFinalizer(a, finalizer)

	return a
}

func (a *Adapter) createDatabase() error {
	engine, err := xorm.NewEngine(a.driverName, a.dataSourceName)
	if err != nil {
		return err
	}
	defer engine.Close()

	_, err = engine.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s default charset utf8 COLLATE utf8_general_ci", a.dbName))
	return err
}

func (a *Adapter) open() {
	if a.driverName != "postgres" {
		if err := a.createDatabase(); err != nil {
			panic(err)
		}
	}

	engine, err := xorm.NewEngine(a.driverName, a.dataSourceName+a.dbName)
	if err != nil {
		panic(err)
	}

	a.Engine = engine
}

func (a *Adapter) close() {
	a.Engine.Close()
	a.Engine = nil
}

func (a *Adapter) createTable() {
	err := a.Engine.Sync2(
		new(Project),

		new(LearnProject),
		new(LearnSection),

		new(Chapter),
		new(Section),

		new(Resource),
		new(Task),
		new(Survey),
		new(Question),

		new(Submit),
		new(Choice),

		new(ProjectSkill),
		new(ProjectSubject),
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
