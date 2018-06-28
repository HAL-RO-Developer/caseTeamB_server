package controller

import (
	"time"

	"github.com/HAL-RO-Developer/caseTeamB_server/controller/response"
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/validation"
	"github.com/HAL-RO-Developer/caseTeamB_server/service"
	"github.com/gin-gonic/gin"
)

var Goal = goalimpl{}

type goalimpl struct {
	Created  time.Time  `json:"created_at"`
	ChildId  int        `json:"child_id"`
	GoalId   string     `json:"goal_id"`
	DeviceId string     `json:"device_id"`
	Run      int        `json:"run"`
	Content  string     `json:"content"`
	Criteria int        `json:"criteria"`
	Deadline *time.Time `json:"deadline"`
	Status   int        `json:"status"`
	Updated  time.Time  `json:"updated_at"`
}

// 目標の新規追加
func (g *goalimpl) CreateGoal(c *gin.Context) {
	name, ok := authorizationCheck(c)
	if !ok {
		response.BadRequest(gin.H{"error": "ログインエラー"}, c)
		return
	}

	req, ok := validation.GoalRegistrationValidation(c)
	if !ok {
		return
	}

	goalId, err := service.RegistrationGoal(name, req)
	if err != nil {
		response.BadRequest(gin.H{"error": "データベースエラー"}, c)
		return
	}
	response.Json(gin.H{"goal_id": goalId}, c)
}

// 目標更新
func (g *goalimpl) UpdateGoal(c *gin.Context) {
	name, ok := authorizationCheck(c)
	if !ok {
		response.BadRequest(gin.H{"error": "ログインエラー"}, c)
		return
	}

	req, ok := validation.GoalUpdateValidation(c)
	if !ok {
		return
	}

	// デバイスIDを登録
	err := service.UpdateGoal(name, req)
	if err != nil {
		response.BadRequest(gin.H{"error": "デバイスIDを登録できませんでした。"}, c)
		return
	}
	response.Json(gin.H{"success": "登録しました"}, c)
}

// 目標取得
func (g *goalimpl) GetGoal(c *gin.Context) {
	var goals []goalimpl
	var goal goalimpl
	name, ok := authorizationCheck(c)
	if !ok {
		response.BadRequest(gin.H{"error": "ログインエラー"}, c)
		return
	}

	// ボタンIDを検索
	data, find := service.GetGoal(name)
	if !find {
		response.BadRequest(gin.H{"error": "目標が登録されていません。"}, c)
		return
	}

	if find {
		for i := 0; i < len(data); i++ {
			status := 0
			if data[i].Run < data[i].Criteria {
				status = 1
			} else if data[i].Run >= data[i].Criteria && data[i].CreatedAt.Unix() > data[i].Deadline.Unix() {
				status = 3
			} else {
				status = 2
			}

			goal.Created = data[i].CreatedAt
			goal.ChildId = data[i].ChildId
			goal.GoalId = data[i].GoalId
			goal.DeviceId = data[i].DeviceId
			goal.Run = data[i].Run
			goal.Content = data[i].Content
			goal.Criteria = data[i].Criteria
			goal.Deadline = data[i].Deadline
			goal.Status = status
			goal.Updated = data[i].UpdatedAt
			goals = append(goals, goal)
		}
		response.Json(gin.H{"goals": goals}, c)
		return
	}
	response.BadRequest(gin.H{"error": "目標が見つかりませんでした。"}, c)
}

// 目標削除
func (g *goalimpl) DeleteGoal(c *gin.Context) {
	_, ok := authorizationCheck(c)
	if !ok {
		response.BadRequest(gin.H{"error": "ログインエラー"}, c)
		return
	}

	goalId := c.Param("goal_id")

	// 目標の削除
	if service.DeleteGoal(goalId) {
		response.Json(gin.H{"success": "目標を削除しました。"}, c)
		return
	}
	response.BadRequest(gin.H{"error": "目標が見つかりません。"}, c)
}
