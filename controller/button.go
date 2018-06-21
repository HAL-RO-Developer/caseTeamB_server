package controller

import (
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/response"
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/validation"
	"github.com/HAL-RO-Developer/caseTeamB_server/service"
	"github.com/gin-gonic/gin"
)

var Button = buttonimpl{}

type buttonimpl struct {
}

// ボタン押下回数変更
func (b *buttonimpl) DeviceIncrement(c *gin.Context) {
	req, ok := validation.ButtonCheck(c)
	if !ok {
		return
	}

	if !service.ExisByDeviceId(req.DeviceId) {
		response.BadRequest(gin.H{"error": "goal_idが見つかりません。"}, c)
	} else {
		err := service.IncrementButton(req.DeviceId)
		if err != nil {
			response.BadRequest(gin.H{"error": "データベースエラー"}, c)
			return
		}
		response.Json(gin.H{"success": "プッシュ回数を変更しました。"}, c)
	}
}
