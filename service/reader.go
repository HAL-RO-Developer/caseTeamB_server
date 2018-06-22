package service

import "github.com/HAL-RO-Developer/caseTeamB_server/model"

func SendUserSolution(device_id string, bookId int, questionNo int, answer int) error {
	record := model.Record{
		DeviceId:   device_id,
		BookId:     bookId,
		QuestionNo: questionNo,
		Answer:     answer,
	}

	return db.Create(&record).Error
}
