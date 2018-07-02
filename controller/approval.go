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
		response.TokenError(gin.H{"error": "アクセストークンが不正です。"}, c)
		return
	}

	req, ok := validation.ApprovalCheck(c)
	if !ok {
		return
	}

	_, find := service.GetOneGoal(req.GoalId)
	if !find {
		response.BadRequest(gin.H{"error": "その目標は見つかりません。"}, c)
		return
	}

	// 目標の達成承認
	success := service.ApprovalGoal(req.GoalId, req.Approval)
	if !success {
		response.BadRequest(gin.H{"error": "データベースエラー"}, c)
		return
	}
	response.Json(gin.H{"success": "目標達成数を変更しました。"}, c)
}
