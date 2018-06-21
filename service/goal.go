package service

import (
	"github.com/HAL-RO-Developer/caseTeamB_server/model"
)

// 目標の新規登録
func RegistrationGoal(goal string, deviceId string) error {
	registration := model.GoalDate{
		DeviceId: deviceId,
		Content:  goal,
		Run:      0,
	}
	err := db.Create(&registration).Error
	return err
}

// 目標取得
func ExisByDeviceIdFromGoal(deviceId string) ([]model.GoalDate, bool) {
	var goals []model.GoalDate
	db.Where("device_id = ?", deviceId).Find(&goals)
	return goals, len(goals) != 0
}

// 目標削除
func DeleteGoal(deviceId string) bool {
	var goals model.Device
	err := db.Where("device_id = ?", deviceId).First(&goals).Error
	if err != nil {
		return false
	}
	db.Delete(goals)
	return true
}
