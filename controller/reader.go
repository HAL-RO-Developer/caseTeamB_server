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
		return
	}

	if service.SendUserAnswer(req.DeviceId, req.Uuid) {
		response.BadRequest(gin.H{"error": "送信に失敗しました。"}, c)
		return
	}
	response.Json(gin.H{"success": "送信しました。"}, c)
}
