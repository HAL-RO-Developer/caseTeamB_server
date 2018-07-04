package controller

import (
	"time"

	"github.com/HAL-RO-Developer/caseTeamB_server/controller/response"
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/validation"
	"github.com/HAL-RO-Developer/caseTeamB_server/model"
	"github.com/HAL-RO-Developer/caseTeamB_server/service"
	"github.com/gin-gonic/gin"
)

var Goal = goalimpl{}

type goalimpl struct {
}

type goalInfo struct {
	ChildId  int        `json:"child_id"`
	Nickname string     `json:"nickname"`
	Goals    []goalData `json:"child_goals"`
}

type goalData struct {
	Created  time.Time  `json:"created_at"`
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
		response.TokenError(gin.H{"error": "アクセストークンが不正です。"}, c)
		return
	}

	req, ok := validation.GoalRegistrationValidation(c)
	if !ok {
		return
	}

	goalId, err := service.RegistrationGoal(name, req)
	if err != nil {
		response.BadRequest(gin.H{"error": goalId}, c)
		return
	}
	response.Json(gin.H{"goal_id": goalId}, c)
}

// 目標更新
func (g *goalimpl) UpdateGoal(c *gin.Context) {
	name, ok := authorizationCheck(c)
	if !ok {
		response.TokenError(gin.H{"error": "アクセストークンが不正です。"}, c)
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
	var userGoal []goalInfo
	var childGoal goalInfo
	var goal goalData
	var childData []model.UserChild
	var status int

	name, ok := authorizationCheck(c)
	if !ok {
		response.TokenError(gin.H{"error": "アクセストークンが不正です。"}, c)
		return
	}

	children, _ := service.GetChildInfo(name)

	/* 子どもIDの数繰り返し */
	for i := 0; i < len(children); i++ {
		// ボタンIDを検索
		goals, find := service.GetGoalForChild(name, children[i].ChildId)
		if !find {

		} else {
			childGoal.ChildId = children[i].ChildId
			childData, _ = service.GetByChildInfo(name, children[i].ChildId)
			childGoal.Nickname = childData[0].NickName
			/* メッセージの数繰り返し */
			for j := 0; j < len(goals); j++ {
				status = 0
				if goals[j].Run < goals[j].Criteria {
					status = 1
				} else if goals[j].Run >= goals[j].Criteria && goals[j].CreatedAt.Unix() > goals[j].Deadline.Unix() {
					status = 3
				} else {
					status = 2
				}

				goal.Created = goals[j].CreatedAt
				goal.GoalId = goals[j].GoalId
				goal.DeviceId = goals[j].DeviceId
				goal.Run = goals[j].Run
				goal.Content = goals[j].Content
				goal.Criteria = goals[j].Criteria
				goal.Deadline = goals[j].Deadline
				goal.Status = status
				goal.Updated = goals[j].UpdatedAt
				childGoal.Goals = append(childGoal.Goals, goal)
				goal = goalData{}
			}
			userGoal = append(userGoal, childGoal)
			childGoal.Goals = nil
		}
	}
	response.Json(gin.H{"goals": userGoal}, c)
}

// 目標削除
func (g *goalimpl) DeleteGoal(c *gin.Context) {
	_, ok := authorizationCheck(c)
	if !ok {
		response.TokenError(gin.H{"error": "アクセストークンが不正です。"}, c)
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
