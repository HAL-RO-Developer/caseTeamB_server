package controller

import (
	"github.com/HAL-RO-Developer/caseTeamB_server/model"
	"github.com/HAL-RO-Developer/caseTeamB_server/service"
	"github.com/gin-gonic/gin"
)

var Device = deviceimpl{}

type deviceimpl struct {
}

// ボタンとIDの紐付け
func (d *deviceimpl) DeviceRegistration(c *gin.Context) {
	var device model.Button
	err := c.BindJSON(&device)
	if err != nil {
		BadRequest("pinもしくはmacアドレスが未入力です。", c)
		return
	}

	if !service.ExisByPin(device.Pin) {
		BadRequest("pinが見つかりません。", c)
	} else {
		button_id, err := service.RegistrationButton(device.Pin, device.Mac)
		if err != nil {
			BadRequest("データベースエラー", c)
			return
		}
		Json(gin.H{"button_id": button_id}, c)
	}
}

// ボタン押下回数インクリメント
func (d *deviceimpl) DeviceIncrement(c *gin.Context) {
	var device model.Button
	err := c.BindJSON(&device)
	if err != nil {
		BadRequest("button_idが未入力です。", c)
		return
	}

	if !service.ExisByButtonId(device.ButtonId) {
		BadRequest("button_idが見つかりません。", c)
	} else {
		err := service.IncrementButton(device.ButtonId)
		if err != nil {
			BadRequest("データベースエラー", c)
			return
		}
		Json(gin.H{"success": "プッシュ回数を追加しました。"}, c)
	}
}
