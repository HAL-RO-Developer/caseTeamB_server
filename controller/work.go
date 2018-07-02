package controller

import (
	"time"

	"github.com/HAL-RO-Developer/caseTeamB_server/controller/response"
	"github.com/HAL-RO-Developer/caseTeamB_server/service"
	"github.com/gin-gonic/gin"
)

var Record = recordimpl{}

type recordimpl struct {
	Date       time.Time `json:"date"`
	GenreName  string    `json:"genre_name"`
	Sentence   string    `json:"sentence"`
	UserAnswer string    `json:"user_answer"`
	Correct    string    `json:"correct"`
	Result     bool      `json:"result"`
}

// 回答記録取得
func (r *recordimpl) WorkRecord(c *gin.Context) {
	var record []recordimpl
	var userRecord recordimpl

	_, ok := authorizationCheck(c)
	if !ok {
		response.TokenError(gin.H{"error": "アクセストークンが不正です。"}, c)
		return
	}

	readerId := c.Param("device_id")
	records, find := service.ExisByRecord(readerId)

	// 回答情報が見つかった時
	if find {
		// 回答情報分繰り返し
		for i := 0; i < len(records); i++ {
			correctId := service.GetByCorrect(records[i].BookId, records[i].QuestionNo)
			if correctId == "" {
				response.BadRequest(gin.H{"error": "問題が見つかりませんでした。"}, c)
				return
			}

			userRecord.Date = records[i].UpdatedAt
			userRecord.GenreName = service.GetGenreName(records[i].BookId)
			tagData := service.GetTagDataFromBookId(records[i].BookId, records[i].QuestionNo)
			userRecord.Sentence = tagData[0].Sentence
			userRecord.UserAnswer = records[i].UserAnswer
			correct := service.GetTagDataFromTagId(correctId)
			userRecord.Correct = correct[0].Answer
			if records[i].UserAnswer == correct[0].Answer {
				userRecord.Result = true
			} else {
				userRecord.Result = false
			}
			record = append(record, userRecord)
		}
		response.Json(gin.H{"records": record}, c)
		return
	}
	response.BadRequest(gin.H{"error": "回答情報が見つかりませんでした。"}, c)
}
