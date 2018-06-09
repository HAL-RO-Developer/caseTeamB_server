package controller

import (
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
		BadRequest("ログインエラー", c)
		return
	}

	pin, err := service.CreateButton(name)
	if err != nil {
		BadRequest("データベースエラー", c)
		return
	}
	Json(gin.H{"pin": pin}, c)
}

// ボタン一覧取得
func (b *buttonimpl) ListButton(c *gin.Context) {
	var buttonId []string
	name, ok := authorizationCheck(c)
	if !ok {
		BadRequest("ログインエラー", c)
		return
	}

	buttons, find := service.GetButtonId(name)
	if find {
		for i := 0; i < len(buttons); i++ {
			buttonId = append(buttonId, buttons[i].ButtonId)
		}
		Json(gin.H{"button_id": buttonId}, c)
		return
	}
	BadRequest("ボタンが登録されていません。", c)
	return
}

// ボタンID削除
func (b *buttonimpl) DeleteButton(c *gin.Context) {
	name, ok := authorizationCheck(c)
	if !ok {
		BadRequest("ログインエラー", c)
		return
	}

	err := c.BindJSON(b)
	if err != nil {
		BadRequest("ボタンIDが入力されていません。", c)
		return
	}

	success := service.DeleteButtonId(name, b.ButtonId)
	if success {
		Json(gin.H{"success": "ボタンIDを削除しました。"}, c)
		return
	}
	BadRequest("ボタンIDが見つかりません。", c)
}
