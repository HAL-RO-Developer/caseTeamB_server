package service

import (
	"github.com/HAL-RO-Developer/caseTeamB_server/model"
)

var db = model.GetDBConn()

// macAddr登録
func RegistrationButton(pin string, mac string) (string, error) {
	button := model.Device{}
	err := db.Where("pin = ?", pin).First(&button).Error
	if err != nil {
		return "not found", err
	}

	button.Mac = mac
	err = db.Model(&button).Update(&button).Update("pin", "").Error
	return button.DeviceId, err
}

// プッシュ回数追加
// Todo BOCCOAPI追記
func IncrementButton(goal_id string) error {
	goal := model.GoalDate{}
	err := db.Where("goal_id = ?", goal_id).First(&goal).Error
	if err != nil {
		return err
	}

	goal.Run++
	err = db.Model(&goal).Update(&goal).Update("run", goal.Run).Error
	return err
}
