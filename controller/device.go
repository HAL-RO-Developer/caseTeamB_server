package controller

import (
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/response"
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/validation"
	"github.com/HAL-RO-Developer/caseTeamB_server/service"
	"github.com/gin-gonic/gin"
)

var Device = deviceimpl{}

type deviceimpl struct {
}

// ボタンとIDの紐付け
func (d *deviceimpl) DeviceRegistration(c *gin.Context) {
	req, ok := validation.ButtonRegistrationCheck(c)
	if !ok {
		return
	}
	if !service.ExisByPin(req.Pin) {
		response.BadRequest("pinが見つかりません。", c)
	} else {
		button_id, err := service.RegistrationButton(req.Pin, req.Mac)
		if err != nil {
			response.BadRequest("データベースエラー", c)
			return
		}
		response.Json(gin.H{"button_id": button_id}, c)
	}
}

// ボタン押下回数インクリメント
func (d *deviceimpl) DeviceIncrement(c *gin.Context) {
	req, ok := validation.ButtonCheck(c)
	if !ok {
		return
	}

	if !service.ExisByButtonId(req.ButtonId) {
		response.BadRequest("button_idが見つかりません。", c)
	} else {
		err := service.IncrementButton(req.ButtonId)
		if err != nil {
			response.BadRequest("データベースエラー", c)
			return
		}
		response.Json(gin.H{"success": "プッシュ回数を追加しました。"}, c)
	}
}
