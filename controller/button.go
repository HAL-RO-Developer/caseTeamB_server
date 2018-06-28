package controller

import (
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/response"
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/validation"
	"github.com/HAL-RO-Developer/caseTeamB_server/service"
	"github.com/gin-gonic/gin"
	_ "github.com/satori/go.uuid"
)

var Button = buttonimpl{}

type buttonimpl struct {
}

// ボタン押下回数変更
// TODO BOCCO喋る処理追加
func (b *buttonimpl) DeviceIncrement(c *gin.Context) {
	req, ok := validation.ButtonCheck(c)
	if !ok {
		return
	}

	_, find := service.ExisByDeviceId(req.DeviceId)
	if !find {
		response.BadRequest(gin.H{"error": "デバイスIDが見つかりません。"}, c)
	} else {
		data, find := service.GetGoalFromDeviceId(req.DeviceId)
		if !find {
			response.BadRequest(gin.H{"error": "目標が見つかりません。"}, c)
			return
		}
		success := service.ApprovalGoal(data[0].GoalId, 1)
		if !success {
			response.BadRequest(gin.H{"error": "押下回数を追加できませんでした"}, c)
			return
		}
		//boccoInfo, _ := service.ExisByBoccoAPI(bocco[0].Name)
		//boccoToken, _ := service.GetBoccoToken(boccoInfo[0].Email, boccoInfo[0].Key, boccoInfo[0].Pass)
		//roomId, _ := service.GetRoomId(boccoToken)
		//uuid := uuid.Must(uuid.NewV4()).String()
		//service.SendMessage(uuid, roomId, boccoToken, "おはよう")

		response.Json(gin.H{"success": "プッシュ回数を変更しました。"}, c)
	}
}
