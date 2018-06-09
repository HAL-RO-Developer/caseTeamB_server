package main

import "github.com/HAL-RO-Developer/caseTeamB_server/model"

func main() {
	db := model.GetDBConn()

	db.DropTableIfExists(&model.User{})
	db.DropTableIfExists(&model.Button{})
	db.DropTableIfExists(&model.Goal{})
	db.DropTableIfExists(&model.Message{})

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Button{})
	db.AutoMigrate(&model.Goal{})
	db.AutoMigrate(&model.Message{})
}
