package controller

import (
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/response"
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/validation"
	"github.com/HAL-RO-Developer/caseTeamB_server/service"
	"github.com/gin-gonic/gin"
)

var Button = buttonimpl{}

type buttonimpl struct {
	ButtonId string `json:"button_id"`
}

// ボタンID発行
func (b *buttonimpl) CreateNewButton(c *gin.Context) {
	name, ok := authorizationCheck(c)
	if !ok {
		response.BadRequest("ログインエラー", c)
		return
	}

	pin, err := service.CreateButton(name)
	if err != nil {
		response.BadRequest("データベースエラー", c)
		return
	}
	response.Json(gin.H{"pin": pin}, c)
}

// ボタン一覧取得
func (b *buttonimpl) ListButton(c *gin.Context) {
	var buttonId []string
	name, ok := authorizationCheck(c)
	if !ok {
		response.BadRequest("ログインエラー", c)
		return
	}

	buttons, find := service.GetButtonId(name)
	if find {
		for i := 0; i < len(buttons); i++ {
			buttonId = append(buttonId, buttons[i].ButtonId)
		}
		response.Json(gin.H{"button_id": buttonId}, c)
		return
	}
	response.BadRequest("ボタンが登録されていません。", c)
	return
}

// ボタンID削除
func (b *buttonimpl) DeleteButton(c *gin.Context) {
	name, ok := authorizationCheck(c)
	if !ok {
		response.BadRequest("ログインエラー", c)
		return
	}

	req, ok := validation.ButtonCheck(c)
	if !ok {
		return
	}

	success := service.DeleteButtonId(name, req.ButtonId)
	if success {
		response.Json(gin.H{"success": "ボタンIDを削除しました。"}, c)
		return
	}
	response.BadRequest("ボタンIDが見つかりません。", c)
}
