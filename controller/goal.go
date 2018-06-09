package controller

import (
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/response"
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/validation"
	"github.com/HAL-RO-Developer/caseTeamB_server/service"
	"github.com/gin-gonic/gin"
)

var Goal = goalimpl{}

type goalimpl struct {
}

// 目標の新規追加
func (g *goalimpl) CreateGoal(c *gin.Context) {
	_, ok := authorizationCheck(c)
	if !ok {
		response.BadRequest("ログインエラー", c)
		return
	}

	req, ok := validation.GoalRegistrationCheck(c)
	if !ok {
		return
	}
	// 目標の重複チェック
	if service.ExisByGoal(req.Contents) {
		response.BadRequest("その目標は登録済みです。", c)
		return
	}
	err := service.RegistrationGoal(req.Contents, req.ButtonId)
	if err != nil {
		response.BadRequest("データベースエラー", c)
		return
	}
	response.Json(gin.H{"success": "目標を追加しました。"}, c)

}
