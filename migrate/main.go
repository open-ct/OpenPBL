package main

import (
	"fmt"
	"openpbl-go/models"
	"openpbl-go/models/db"
)

func main() {
	_db := db.GetEngine()
	// db.GetEngine().DropTables()
	err := _db.Sync2(
		new(models.Project),
		new(models.Teacher),
		new(models.Student),
		new(models.LearnProject),

		new(models.Chapter),
		new(models.Section),
		new(models.SubmitFile),
		new(models.ProjectSkill),
		new(models.ProjectSubject),
		)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}