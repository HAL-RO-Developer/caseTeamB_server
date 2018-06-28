package service

import (
	"time"

	"github.com/HAL-RO-Developer/caseTeamB_server/controller/validation"
	"github.com/HAL-RO-Developer/caseTeamB_server/model"
)

// 目標の新規登録
func RegistrationGoal(name string, info validation.Goal) (string, error) {
	var registration model.GoalData
	var deadline time.Time
	var err error

	goalId := createGoalId()

	// 目標に期限がない時
	if info.Deadline == "" {
		registration = model.GoalData{
			Name:     name,
			ChildId:  info.ChildId,
			GoalId:   goalId,
			Content:  info.Content,
			Run:      0,
			Criteria: info.Criteria,
		}
	} else {
		deadline, err = time.Parse("2006-01-02", info.Deadline)
		if err != nil {
			return "", err
		}
		registration = model.GoalData{
			Name:     name,
			ChildId:  info.ChildId,
			GoalId:   goalId,
			Content:  info.Content,
			Run:      0,
			Criteria: info.Criteria,
			Deadline: &deadline,
		}
	}
	err = db.Create(&registration).Error
	return goalId, err
}

// ボタン登録
func UpdateGoal(name string, info validation.UpdateGoal) error {
	var goal model.GoalData
	err := db.Where("name = ? and goal_id = ?", name, info.GoalId).First(&goal).Error
	if err != nil {
		return err
	}

	err = db.Model(&goal).Update(&goal).Update("device_id", info.DeviceId).Error
	return err
}

// 目標取得(ユーザー毎)
func GetGoal(name string) ([]model.GoalData, bool) {
	var goals []model.GoalData
	db.Where("name = ?", name).Find(&goals)
	return goals, len(goals) != 0
}

// 目標取得(デバイスID)
func GetGoalFromDeviceId(deviceId string) ([]model.GoalData, bool) {
	var goals []model.GoalData
	db.Where("device_id = ?", deviceId).First(&goals)
	return goals, len(goals) != 0
}

// 目標取得(子ども毎)
func GetGoalForChild(name string, childId int) ([]model.GoalData, bool) {
	var goals []model.GoalData
	db.Where("name = ? and child_id = ?", name, childId).Find(&goals)
	return goals, len(goals) != 0
}

// 目標取得(特定)
func GetOneGoal(goalId string) (model.GoalData, bool) {
	var goal model.GoalData
	err := db.Where("goal_id = ?", goalId).First(&goal).Error
	if err != nil {
		return goal, false
	}
	return goal, true
}

// 目標削除
func DeleteGoal(goalId string) bool {
	var goal model.GoalData
	err := db.Where("goal_id = ?", goalId).First(&goal).Error
	if err != nil {
		return false
	}
	db.Delete(goal)
	return true
}

// 目標ID作成
func createGoalId() string {
	var goalId string
	for {
		goalId = createUuid(16, []rune("ABCDEFGHRJKLNMOPQRSTUPWXYZabcdefghijklmnopqrstuvwxyz0123456789"))
		_, find := GetOneGoal(goalId)
		if !find {
			break
		}
	}
	return goalId
}
