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
	BookId     int       `json:"book_id"`
	QuestionNo int       `json:"q_no"`
	Answer     int       `json:"answer"`
	Correct    int       `json:"correct"`
}

// 回答記録取得
func (r *recordimpl) WorkRecord(c *gin.Context) {
	var record []recordimpl
	var userRecord recordimpl

	_, ok := authorizationCheck(c)
	if !ok {
		response.BadRequest(gin.H{"error": "ログインエラー"}, c)
		return
	}

	readerId := c.Param("device_id")
	records, find := service.ExisByRecord(readerId)

	if find {
		for i := 0; i < len(records); i++ {
			correct := service.ExisByCorrect(records[i].BookId, records[i].QuestionNo)
			if correct == 0 {
				response.BadRequest(gin.H{"error": "問題が見つかりませんでした。"}, c)
				return
			}

			userRecord.Date = records[i].UpdatedAt
			userRecord.BookId = records[i].BookId
			userRecord.QuestionNo = records[i].QuestionNo
			userRecord.Answer = records[i].Answer
			userRecord.Correct = correct
			record = append(record, userRecord)
		}
		response.Json(gin.H{"data": record}, c)
		return
	}
	response.BadRequest(gin.H{"error": "回答情報が見つかりませんでした。"}, c)
}