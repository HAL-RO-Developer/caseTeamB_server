package main

import "github.com/HAL-RO-Developer/caseTeamB_server/model"

func main() {
	db := model.GetDBConn()

	db.DropTableIfExists(&model.User{})
	db.DropTableIfExists(&model.Token{})
	db.DropTableIfExists(&model.Recipi{})

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Token{})
	db.AutoMigrate(&model.Recipi{})
}
