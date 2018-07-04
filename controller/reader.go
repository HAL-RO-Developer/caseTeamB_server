package controller

import (
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/response"
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/validation"
	"github.com/HAL-RO-Developer/caseTeamB_server/service"
	"github.com/gin-gonic/gin"
)

var Reader = readerimpl{}

type readerimpl struct {
}

// リーダーで読み取った情報の送信
func (r *readerimpl) SendTag(c *gin.Context) {
	req, ok := validation.ReaderValidation(c)
	if !ok {
		response.BadRequest(gin.H{"error": "不正なリクエストです。"}, c)
		return
	}

	result := service.SendUserAnswer(req.DeviceId, req.Uuid)
	if result == -2 {
		response.BadRequest(gin.H{"error": "BOCCOAPI情報が登録されていません。"}, c)
		return
	} else if result == -1 {
		response.BadRequest(gin.H{"error": "送信に失敗しました。"}, c)
		return
	} else if result == 0 {
		response.Json(gin.H{"success": false}, c)
		return
	}
	response.Json(gin.H{"success": true}, c)
}
