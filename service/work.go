package service

import "github.com/HAL-RO-Developer/caseTeamB_server/model"

// 端末ごとの回答情報取得
func ExisByRecord(readerId string) ([]model.Record, bool) {
	var records []model.Record
	db.Where("reader_id = ?", readerId).Find(&records)
	return records, len(records) != 0
}

// 問題番号から答えを取得
func ExisByCorrect(bookId int, questionNo int) int {
	var question []model.Question
	err := db.Where("book_id = ? and q_no = ?", bookId, questionNo).Find(&question).Error

	if err != nil {
		return 0
	}
	return question[0].Correct
}
