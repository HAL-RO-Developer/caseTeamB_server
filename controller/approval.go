package controller

import (
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/response"
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/validation"
	"github.com/HAL-RO-Developer/caseTeamB_server/service"
	"github.com/gin-gonic/gin"
)

var Approval = approvalimpl{}

type approvalimpl struct {
}

// 目標承認
func (A *approvalimpl) ApprovalGoal(c *gin.Context) {
	_, ok := authorizationCheck(c)
	if !ok {
		response.BadRequest(gin.H{"error": "ログインエラー"}, c)
		return
	}

	req, ok := validation.ButtonCheck(c)
	if !ok {
		return
	}

	_, find := service.ExisByButtonIdFromGoal(req.ButtonId)
	if !find {
		response.BadRequest(gin.H{"error": "そのボタンは見つかりません。"}, c)
		return
	}

	// 目標の達成承認
	err := service.ApprovalGoal(req.ButtonId)
	if err != nil {
		response.BadRequest(gin.H{"error": "データベースエラー"}, c)
		return
	}
	response.Json(gin.H{"success": "目標達成を承認しました。"}, c)
}

// 目標非承認
func (A *approvalimpl) NotApprovalGoal(c *gin.Context) {
	_, ok := authorizationCheck(c)
	if !ok {
		response.BadRequest(gin.H{"error": "ログインエラー"}, c)
		return
	}

	buttonId := c.PostForm("button_id")
	_, find := service.ExisByButtonIdFromGoal(buttonId)
	if !find {
		response.BadRequest(gin.H{"error": "ボタンIDは見つかりません。"}, c)
		return
	}

	// 目標達成非承認
	err := service.NotApprovalGoal(buttonId)
	if err != nil {
		response.BadRequest(gin.H{"error": "データベースエラー"}, c)
		return
	}
	response.Json(gin.H{"success": "目標達成を非承認にしました。"}, c)
}
