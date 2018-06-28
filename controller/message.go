package controller

import (
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/response"
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/validation"
	"github.com/HAL-RO-Developer/caseTeamB_server/model"
	"github.com/HAL-RO-Developer/caseTeamB_server/service"
	"github.com/gin-gonic/gin"
)

var Message = messageimpl{}

type messageimpl struct {
	GoalId    string `json:"goal_id"`
	Content   string `json:"content"`   // 目標内容
	Condition int    `json:"condition"` // メッセージ発信条件
	Message   string `json:"message"`   // メッセージ内容
}

// メッセージ編集
func (m *messageimpl) EditMessage(c *gin.Context) {
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
	_, find := service.GetOneGoal(req.GoalId)
	if !find {
		response.BadRequest(gin.H{"error": "目標が登録されていません。"}, c)
		return
	}

	// 新規メッセージ登録
	err := service.RegistrationMessage(req.GoalId, req.Condition, req.Message)
	if err != nil {
		response.BadRequest(gin.H{"error": "データベースエラー"}, c)
		return
	}
	response.Json(gin.H{"success": "メッセージを登録しました。"}, c)
}

// メッセージ取得
func (m *messageimpl) GetMessage(c *gin.Context) {
	var messages []messageimpl
	var message messageimpl
	var buf model.GoalData

	name, ok := authorizationCheck(c)
	if !ok {
		response.BadRequest(gin.H{"error": "アクセストークンが不正です。"}, c)
		return
	}

	data, find := service.GetMessageFromName(name)
	if !find {
		response.BadRequest(gin.H{"error": "メッセージが見つかりません。"}, c)
		return
	}

	for i := 0; i < len(data); i++ {
		message.GoalId = data[i].GoalId
		buf, _ = service.GetOneGoal(data[i].GoalId)
		message.Content = buf.Content
		message.Condition = data[i].Condition
		message.Message = data[i].Message
		messages = append(messages, message)
	}
	response.Json(gin.H{"messages": messages}, c)
}
