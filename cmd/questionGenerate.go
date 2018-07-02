package main

import "github.com/HAL-RO-Developer/caseTeamB_server/model"

func main() {
	questionGenerate()
}

func questionGenerate() {
	db := model.GetDBConn()

	question := model.Question{
		BookId:     1,
		QuestionNo: 1,
		Sentence:   "question1",
		Correct:    "answer1",
	}
	db.Create(&question)
}

func tagGenerate() {
	db := model.GetDBConn()

	tag := model.Tag{
		TagId:      "question1",
		Uuid:       "sample",
		BookId:     1,
		QuestionNo: 1,
		Sentence:   "1+1は",
	}
}
