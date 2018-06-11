package controller

import (
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/response"
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/validation"
	"github.com/HAL-RO-Developer/caseTeamB_server/service"
	"github.com/gin-gonic/gin"
)

var Message = messageimpl{}

type messageimpl struct {
}

// 新規メッセージ登録
func (m *messageimpl) NewMessage(c *gin.Context) {
	_, ok := authorizationCheck(c)
	if !ok {
		response.BadRequest(gin.H{"error": "ログインエラー"}, c)
		return
	}

	req, ok := validation.MessageValidation(c)
	if !ok {
		return
	}

	// データベースへの重複防止
	_, find := service.ExisByButtonIdFromGoal(req.ButtonId)
	if !find {
		response.BadRequest(gin.H{"error": "目標が登録されていません。"}, c)
		return
	}
	_, find = service.ExisByButtonIdFromMessage(req.ButtonId)
	if find {
		response.BadRequest(gin.H{"error": "メッセージ登録済みです。"}, c)
		return
	}

	// 新規メッセージ登録
	err := service.RegistrationMessage(req.ButtonId, req.Condition, req.Message)
	if err != nil {
		response.BadRequest(gin.H{"error": "データベースエラー"}, c)
		return
	}
	response.Json(gin.H{"success": "メッセージを登録しました。"}, c)
}

// メッセージ取得
func (m *messageimpl) GetMessage(c *gin.Context) {
	_, ok := authorizationCheck(c)
	if !ok {
		response.BadRequest(gin.H{"error": "ログインエラー"}, c)
		return
	}

	buttonId := c.Param("button_id")

	message, find := service.ExisByButtonIdFromMessage(buttonId)
	if !find {
		response.BadRequest(gin.H{"error": "ボタンIDが見つかりません。"}, c)
		return
	}
	response.Json(gin.H{message[0].Message: message[0].Condition}, c)
}

// メッセージ削除
func (m *messageimpl) DeleteMessage(c *gin.Context) {
	_, ok := authorizationCheck(c)
	if !ok {
		response.BadRequest(gin.H{"error": "ログインエラー"}, c)
		return
	}

	buttonId := c.Param("button_id")

	_, find := service.ExisByButtonIdFromMessage(buttonId)
	if !find {
		response.BadRequest(gin.H{"error": "ボタンIDが見つかりません。"}, c)
		return
	}

	if service.DeleteMessage(buttonId) {
		response.Json(gin.H{"success": "メッセージを削除しました。"}, c)
		return
	}
	response.BadRequest(gin.H{"error": "データベースエラー"}, c)
}
