package service

import "github.com/HAL-RO-Developer/caseTeamB_server/model"

// 端末ごとの回答情報取得
func ExisByRecord(deviceId string) ([]model.Record, bool) {
	var records []model.Record
	db.Where("device_id = ?", deviceId).Find(&records)
	return records, len(records) != 0
}

// 問題番号から答えを取得
func GetByCorrect(bookId int, questionNo int) string {
	var question []model.Question
	err := db.Where("book_id = ? and q_no = ?", bookId, questionNo).Find(&question).Error

	if err != nil {
		return ""
	}
	return question[0].Correct
}

// ジャンル名称の取得
func GetGenreName(bookId int) string {
	book := getBookData(bookId)
	if book == nil {
		return ""
	}
	genre := getGenreData(book[0].GenreId)
	if genre == nil {
		return ""
	}

	return genre[0].GenreName
}

// 本情報の取得
func getBookData(bookId int) []model.Book {
	var book []model.Book
	err := db.Where("book_id = ?", bookId).Find(&book).Error
	if err != nil {
		return nil
	}
	return book
}

// ジャンル情報の取得
func getGenreData(genreId int) []model.Genre {
	var genre []model.Genre
	err := db.Where("genre_id = ?", genreId).Find(&genre).Error
	if err != nil {
		return nil
	}
	return genre
}

// 問題情報の取得
func getQuestionData(bookId int) []model.Question {
	var question []model.Question
	err := db.Where("book_id = ?", bookId).Find(&question).Error
	if err != nil {
		return nil
	}
	return question
}

// タグ情報の取得(タグIDから)
func GetTagDataFromTagId(tagId string) []model.Tag {
	var tag []model.Tag
	err := db.Where("tag_id = ?", tagId).Find(&tag).Error
	if err != nil {
		return nil
	}
	return tag
}

// タグ情報の取得(bookId&questionNoから)
func GetTagDataFromBookId(bookId int, questionId int) []model.Tag {
	var tag []model.Tag
	err := db.Where("book_id = ? and q_no = ?", bookId, questionId).Find(&tag).Error
	if err != nil {
		return nil
	}
	return tag
}

// タグ情報の取得(uuidから)
func GetTagDataFromUuid(uuid string) ([]model.Tag, bool) {
	var tag []model.Tag
	err := db.Where("uuid = ?", uuid).Find(&tag).Error
	if err != nil {
		return nil, false
	}
	return tag, true
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
