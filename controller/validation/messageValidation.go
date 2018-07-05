package validation

import (
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/response"
	"github.com/gin-gonic/gin"
)

type Message struct {
	GoalId      string `json:"goal_id"`
	MessageCall int    `json:"message_call"`
	Message     string `json:"message"`
}

type WorkMessge struct {
	ChildId     int    `json:"child_id"`
	MessageCall int    `json:"message_call"` // メッセージ発信条件
	Condition   int    `json:"condition"`
	Message     string `json:"message"` // メッセージ内容
}

func MessageValidation(c *gin.Context) (Message, bool) {
	var req Message
	err := c.BindJSON(&req)
	if err != nil || req.GoalId == "" || req.Message == "" {
		response.BadRequest(gin.H{"error": "入力されていないデータがあります。"}, c)
		return req, false
	}
	return req, true
}

func WorkMessageValidation(c *gin.Context) (WorkMessge, bool) {
	var req WorkMessge
	err := c.BindJSON(&req)
	if err != nil || req.Message == "" {
		response.BadRequest(gin.H{"error": "入力されていないデータがあります。"}, c)
		return req, false
	}
	return req, true
}
