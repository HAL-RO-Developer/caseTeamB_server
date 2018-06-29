package service

import "github.com/HAL-RO-Developer/caseTeamB_server/model"

// 端末ごとの回答情報取得
func ExisByRecord(deviceId string) ([]model.Record, bool) {
	var records []model.Record
	db.Where("device_id = ?", deviceId).Find(&records)
	return records, len(records) != 0
}

// 問題番号から答えを取得
func ExisByCorrect(bookId int, questionNo int) string {
	var question []model.Question
	err := db.Where("book_id = ? and q_no = ?", bookId, questionNo).Find(&question).Error

	if err != nil {
		return ""
	}
	return question[0].Correct
}

// 回答情報の削除
func DeleteUserAnswer(deviceId string) bool {
	var records []model.Record
	err := db.Where("device_id = ?", deviceId).First(&records).Error
	if err != nil {
		return false
	}
	db.Delete(records)
	return true
}
