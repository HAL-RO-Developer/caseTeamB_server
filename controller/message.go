package controller

import (
	"strconv"

	"github.com/HAL-RO-Developer/caseTeamB_server/controller/response"
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/validation"
	"github.com/HAL-RO-Developer/caseTeamB_server/model"
	"github.com/HAL-RO-Developer/caseTeamB_server/service"
	"github.com/gin-gonic/gin"
)

var Message = messageimpl{}

type messageimpl struct {
}

type messageInfo struct {
	ChildId  int           `json:"child_id"`
	Nickname string        `json:"nickname"`
	Messages []messageData `json:"child_messages"`
}

type messageData struct {
	GoalId      string `json:"goal_id"`
	Content     string `json:"content"`      // 目標内容
	MessageCall int    `json:"message_call"` // メッセージ発信条件
	Message     string `json:"message"`      // メッセージ内容
}

// メッセージ編集
func (m *messageimpl) EditMessage(c *gin.Context) {
	_, ok := authorizationCheck(c)
	if !ok {
		response.TokenError(gin.H{"error": "アクセストークンが不正です。"}, c)
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

	// メッセージ登録確認
	find = service.ExisMessageFromGoal(req.GoalId, req.MessageCall)
	if find {
		// メッセージを更新
		err := service.UpdateMessage(req.GoalId, req.MessageCall, req.Message)
		if err != nil {
			response.BadRequest(gin.H{"error": "メッセージの更新に失敗しました。"}, c)
			return
		}
	} else {
		// 新規メッセージ登録
		err := service.RegistrationMessage(req.GoalId, req.MessageCall, req.Message)
		if err != nil {
			response.BadRequest(gin.H{"error": "データベースエラー"}, c)
			return
		}
	}

	response.Json(gin.H{"success": "メッセージを更新しました。"}, c)
}

// メッセージ取得
func (m *messageimpl) GetMessage(c *gin.Context) {
	var userMessage []messageInfo
	var childMsg messageInfo
	var message messageData
	var messages []model.CustomMessage
	var buf model.GoalData
	var childData []model.UserChild

	name, ok := authorizationCheck(c)
	if !ok {
		response.TokenError(gin.H{"error": "アクセストークンが不正です。"}, c)
		return
	}

	_, find := service.GetMessageFromName(name)
	if !find {
		response.Json(gin.H{"messages": []messageInfo{}}, c)
		return
	}
	children, _ := service.GetChildInfo(name)
	/* 子どもIDの数繰り返し */
	for i := 0; i < len(children); i++ {
		messages, find = service.GetMessageFromNameChild(name, children[i].ChildId)
		if !find {

		} else {
			childMsg.ChildId = children[i].ChildId
			childData, _ = service.GetByChildInfo(name, children[i].ChildId)
			childMsg.Nickname = childData[0].NickName
			/* メッセージの数繰り返し */
			for j := 0; j < len(messages); j++ {
				message.GoalId = messages[j].GoalId
				buf, _ = service.GetOneGoal(messages[j].GoalId)
				message.Content = buf.Content
				message.MessageCall = messages[j].MessageCall
				message.Message = messages[j].Message
				childMsg.Messages = append(childMsg.Messages, message)
				message = messageData{}
			}
			userMessage = append(userMessage, childMsg)
			childMsg.Messages = nil
		}
	}

	response.Json(gin.H{"messages": userMessage}, c)
}

// メッセージ削除
func (m *messageimpl) DeleteMessage(c *gin.Context) {
	_, ok := authorizationCheck(c)
	if !ok {
		response.TokenError(gin.H{"error": "アクセストークンが不正です。"}, c)
		return
	}
	goalId := c.Param("goal_id")
	messageCall := c.Param("message_call")
	message, err := strconv.Atoi(messageCall)
	if err != nil {
		response.BadRequest(gin.H{"error": "メッセージの出力条件が不正です。"}, c)
		return
	}

	success := service.DeleteMessage(goalId, message)
	if !success {
		response.BadRequest(gin.H{"error": "メッセージの削除に失敗しました。"}, c)
		return
	}
	response.Json(gin.H{"success": "メッセージを削除しました。"}, c)
}
