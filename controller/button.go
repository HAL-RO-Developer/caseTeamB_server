package controller

import (
	"fmt"

	"github.com/HAL-RO-Developer/caseTeamB_server/controller/response"
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/validation"
	"github.com/HAL-RO-Developer/caseTeamB_server/service"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	_ "github.com/satori/go.uuid"
)

var Button = buttonimpl{}

type buttonimpl struct {
}

// ボタン押下回数変更
// Todo BOCCOのメッセージ自動生成追加
func (b *buttonimpl) DeviceIncrement(c *gin.Context) {
	req, ok := validation.ButtonCheck(c)
	if !ok {
		return
	}

	bocco, find := service.GetDeviceInfoFromDeviceId(req.DeviceId)
	if find {
		data, find := service.GetGoalFromDeviceId(req.DeviceId)
		if !find {
			response.Json(gin.H{"angle": "目標が見つかりません。"}, c)
			return
		}

		success := service.ApprovalGoal(data[0].GoalId, 1)
		if !success {
			response.BadRequest(gin.H{"error": "押下回数を追加できませんでした"}, c)
			return
		}
		// 押下回数追加後データ取得
		data, _ = service.GetGoalFromDeviceId(req.DeviceId)
		// サーボモーターの移動角度計算
		progress := (float64(data[0].Run) / float64(data[0].Criteria)) * 100 * 10 / 8
		message, find := service.GetMessageFromGoal(data[0].GoalId)
		if !find {
			response.Json(gin.H{"angle": int(progress)}, c)
			return
		}
		// 目標の実行回数がメッセージの発信条件を満たした時
		if data[0].Run == message[0].MessageCall {
			boccoInfo, find := service.ExisByBoccoAPI(bocco[0].Name)
			if !find {
				fmt.Println("test")
				response.Json(gin.H{"angle": int(progress)}, c)
				return
			}
			boccoToken, _ := service.GetBoccoToken(boccoInfo[0].Email, boccoInfo[0].Key, boccoInfo[0].Pass)
			roomId, _ := service.GetRoomId(boccoToken)
			uuid := uuid.Must(uuid.NewV4()).String()
			service.SendMessage(uuid, roomId, boccoToken, message[0].Message)
		}
		response.Json(gin.H{"angle": int(progress)}, c)
		return
	}
	response.Json(gin.H{"angle": "デバイスIDが見つかりません。"}, c)
}
