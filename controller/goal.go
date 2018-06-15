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
		response.BadRequest(gin.H{"error": "ログインエラー"}, c)
		return
	}

	req, ok := validation.GoalRegistrationCheck(c)
	if !ok {
		return
	}
	// 目標の重複チェック
	_, find := service.ExisByButtonIdFromGoal(req.ButtonId)
	if find {
		response.BadRequest(gin.H{"error": "そのボタンは目標登録済みです。"}, c)
		return
	}
	err := service.RegistrationGoal(req.Contents, req.ButtonId)
	if err != nil {
		response.BadRequest(gin.H{"error": "データベースエラー"}, c)
		return
	}
	response.Json(gin.H{"success": "目標を追加しました。"}, c)
}

// 目標取得
func (g *goalimpl) GetGoal(c *gin.Context) {
	_, ok := authorizationCheck(c)
	if !ok {
		response.BadRequest(gin.H{"error": "ログインエラー"}, c)
		return
	}

	buttonId := c.Param("button_id")
	// ボタンIDを検索
	goal, find := service.ExisByButtonIdFromGoal(buttonId)
	if !find {
		response.BadRequest(gin.H{"error": "目標が登録されていません。"}, c)
		return
	}
	response.Json(gin.H{"archive": goal[0].Run, "goal": goal[0].Contents}, c)
}

// 目標削除
func (g *goalimpl) DeleteGoal(c *gin.Context) {
	_, ok := authorizationCheck(c)
	if !ok {
		response.BadRequest(gin.H{"error": "ログインエラー"}, c)
		return
	}

	buttonId := c.Param("button_id")
	// ボタンIDを検索
	_, find := service.ExisByButtonIdFromGoal(buttonId)
	if !find {
		response.BadRequest(gin.H{"error": "ボタンIDが見つかりません。"}, c)
		return
	}

	// 目標の削除
	if service.DeleteGoal(buttonId) {
		response.Json(gin.H{"success": "目標を削除しました。"}, c)
		return
	}
	response.BadRequest(gin.H{"error": "データベースエラー"}, c)
}
