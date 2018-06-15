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

	req, ok := validation.ApprovalCheck(c)
	if !ok {
		return
	}

	_, find := service.ExisByButtonIdFromGoal(req.ButtonId)
	if !find {
		response.BadRequest(gin.H{"error": "そのボタンは見つかりません。"}, c)
		return
	}

	// 目標の達成承認
	fail := service.ApprovalGoal(req.ButtonId, req.Approval)
	if !fail {
		response.BadRequest(gin.H{"error": "データベースエラー"}, c)
		return
	}
	response.Json(gin.H{"success": "目標達成数を変更しました。"}, c)
}
