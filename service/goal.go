package service

import (
	"github.com/HAL-RO-Developer/caseTeamB_server/model"
)

// 目標の新規登録
func RegistrationGoal(goal string, deviceId string) error {
	registration := model.Goal{
		ButtonId: deviceId,
		Contents: goal,
		Run:      "0",
		Approval: "0",
	}
	err := db.Create(&registration).Error
	return err
}

// 目標取得
func ExisByButtonIdFromGoal(buttonId string) ([]model.Goal, bool) {
	var goals []model.Goal
	db.Where("button_id = ?", buttonId).Find(&goals)
	return goals, len(goals) != 0
}

// 目標削除
func DeleteGoal(buttonId string) bool {
	var goals model.Goal
	err := db.Where("button_id = ?", buttonId).First(&goals).Error
	if err != nil {
		return false
	}
	db.Delete(goals)
	return true
}
