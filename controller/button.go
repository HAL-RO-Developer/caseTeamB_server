package controller

import (
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/response"
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
		response.BadRequest(gin.H{"error": "ログインエラー"}, c)
		return
	}

	pin, err := service.CreateButton(name)
	if err != nil {
		response.BadRequest(gin.H{"error": "データベースエラー"}, c)
		return
	}
	response.Json(gin.H{"pin": pin}, c)
}

// ボタン一覧取得
func (b *buttonimpl) ListButton(c *gin.Context) {
	var buttonId []string
	name, ok := authorizationCheck(c)
	if !ok {
		response.BadRequest(gin.H{"error": "ログインエラー"}, c)
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
	response.BadRequest(gin.H{"error": "ボタンが登録されていません。"}, c)
	return
}

// ボタンID削除
func (b *buttonimpl) DeleteButton(c *gin.Context) {
	name, ok := authorizationCheck(c)
	if !ok {
		response.BadRequest(gin.H{"error": "ログインエラー"}, c)
		return
	}

	buttonId := c.PostForm("button_id")

	success := service.DeleteButtonId(name, buttonId)
	if success {
		response.Json(gin.H{"success": "ボタンIDを削除しました。"}, c)
		return
	}
	response.BadRequest(gin.H{"error": "ボタンIDが見つかりません。"}, c)
}
