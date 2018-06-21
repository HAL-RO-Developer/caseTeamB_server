package service

import "github.com/HAL-RO-Developer/caseTeamB_server/model"

func SendUserSolution(macAddr string, bookId int, questionNo int, answer int) error {
	record := model.Record{
		Mac:        macAddr,
		BookId:     bookId,
		QuestionNo: questionNo,
		Answer:     answer,
	}

	return db.Create(&record).Error
}
