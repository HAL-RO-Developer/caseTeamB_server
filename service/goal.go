package service

import (
	"github.com/HAL-RO-Developer/caseTeamB_server/model"
)

// 目標の新規登録
func RegistrationGoal(goal string, deviceId string) error {
	registration := model.Goal{
		ButtonId: deviceId,
		Contents: goal,
	}
	err := db.Create(&registration).Error
	return err
}

// 目標検索
func ExisByGoal(goal string) bool {
	var goals []model.Goal
	db.Where("contents = ?", goal).Find(&goals)
	return len(goals) != 0
}
