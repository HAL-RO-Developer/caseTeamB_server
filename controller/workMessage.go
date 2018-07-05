package controller

import (
	"strconv"

	"github.com/HAL-RO-Developer/caseTeamB_server/controller/response"
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/validation"
	"github.com/HAL-RO-Developer/caseTeamB_server/model"
	"github.com/HAL-RO-Developer/caseTeamB_server/service"
	"github.com/gin-gonic/gin"
)

var WorkMessage = workmessageimpl{}

type workmessageimpl struct {
}

type workMessageInfo struct {
	ChildId  int               `json:"child_id"`
	Nickname string            `json:"nickname"`
	Messages []workMessageData `json:"child_messages"`
}

type workMessageData struct {
	Condition   int            `json:"condition"`
	MessageCall int            `json:"message_call"` // メッセージ発信条件
	Message     []workMessages `json:"messages"`
}

type workMessages struct {
	MessageId string `json:"message_id"`
	Message   string `json:"message"`
}

// メッセージ編集
func (w *workmessageimpl) EditMessage(c *gin.Context) {
	registData := model.WorkCustomMessage{}
	name, ok := authorizationCheck(c)
	if !ok {
		response.TokenError(gin.H{"error": "アクセストークンが不正です。"}, c)
		return
	}

	req, ok := validation.WorkMessageValidation(c)
	if !ok {
		return
	}

	if service.GetWorkMessageFromMessage(name, req.ChildId, req.Message) {
		response.BadRequest(gin.H{"error": "同じメッセージが登録されています。"}, c)
		return
	}
	registData.Name = name
	registData.ChildId = req.ChildId
	registData.Message = req.Message
	registData.Conditions = req.Condition
	registData.MessageCall = req.MessageCall
	// 新規メッセージ登録
	err := service.RegistrationWorkMessage(registData)
	if err != nil {
		response.BadRequest(gin.H{"error": "データベースエラー"}, c)
		return
	}

	response.Json(gin.H{"success": "メッセージを登録しました。"}, c)
}

// メッセージ取得
func (w *workmessageimpl) GetMessage(c *gin.Context) {
	var userMessage []workMessageInfo
	var childMsg workMessageInfo
	var message workMessageData
	var sames []model.WorkCustomMessage
	var sameMessage workMessages
	var messages []model.WorkCustomMessage
	var childData []model.UserChild

	name, ok := authorizationCheck(c)
	if !ok {
		response.TokenError(gin.H{"error": "アクセストークンが不正です。"}, c)
		return
	}
	buf := c.Param("child_id")
	childId, err := strconv.Atoi(buf)
	if err != nil {
		response.BadRequest(gin.H{"error": "childIdが不正です。"}, c)
		return
	}

	messages, find := service.GetWorkMessageFromNameChild(name, childId)
	if !find {
		response.Json(gin.H{"messages": []string{}}, c)
		return
	} else {
		childMsg.ChildId = childId
		childData, _ = service.GetByChildInfo(name, childId)
		childMsg.Nickname = childData[0].NickName
		/* メッセージの数繰り返し */
		for i := 0; i < len(messages); i++ {
			message.Condition = messages[i].Conditions
			message.MessageCall = messages[i].MessageCall
			// 同一条件メッセージの取得
			sames, _ = service.GetMessageInfoFromSame(name, childId, message.Condition, message.MessageCall)
			for j := 0; j < len(sames); j++ {
				sameMessage.MessageId = messages[j].MessageId
				sameMessage.Message = messages[j].Message
				message.Message = append(message.Message, sameMessage)
				sameMessage = workMessages{}
				i++
			}
			childMsg.Messages = append(childMsg.Messages, message)
			message = workMessageData{}
		}
		userMessage = append(userMessage, childMsg)
		childMsg.Messages = nil
	}
	response.Json(gin.H{"messages": userMessage}, c)
}

// メッセージ削除
func (w *workmessageimpl) DeleteMessage(c *gin.Context) {
	_, ok := authorizationCheck(c)
	if !ok {
		response.TokenError(gin.H{"error": "アクセストークンが不正です。"}, c)
		return
	}
	messageId := c.Param("message_id")

	success := service.DeleteWorkMessage(messageId)
	if !success {
		response.BadRequest(gin.H{"error": "メッセージの削除に失敗しました。"}, c)
		return
	}
	response.Json(gin.H{"success": "メッセージを削除しました。"}, c)
}
