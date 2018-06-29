package main

import (
	"github.com/HAL-RO-Developer/caseTeamB_server/model"
)

func main() {
	db := model.GetDBConn()

	db.DropTableIfExists(&model.User{})
	db.DropTableIfExists(&model.UserChild{})
	db.DropTableIfExists(&model.Device{})
	db.DropTableIfExists(&model.Bocco{})
	db.DropTableIfExists(&model.GoalData{})
	db.DropTableIfExists(&model.CustomMessage{})
	db.DropTableIfExists(&model.Book{})
	db.DropTableIfExists(&model.Genre{})
	db.DropTableIfExists(&model.Question{})
	db.DropTableIfExists(&model.Record{})
	db.DropTableIfExists(&model.Tag{})

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.UserChild{})
	db.AutoMigrate(&model.Device{})
	db.AutoMigrate(&model.Bocco{})
	db.AutoMigrate(&model.GoalData{})
	db.AutoMigrate(&model.CustomMessage{})
	db.AutoMigrate(&model.Book{})
	db.AutoMigrate(&model.Genre{})
	db.AutoMigrate(&model.Question{})
	db.AutoMigrate(&model.Record{})
	db.AutoMigrate(&model.Tag{})
}
