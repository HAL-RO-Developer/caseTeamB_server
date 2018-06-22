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

	if service.SendUserSolution(req.DeviceId, req.Data[0].BookId, req.Data[0].QuestionNo, req.Data[0].UserAns) != nil {
		response.BadRequest(gin.H{"error": "データベースエラー"}, c)
		return
	}
	response.Json(gin.H{"success": "送信しました。"}, c)
}
