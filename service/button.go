package service

import (
	"sync"

	"github.com/HAL-RO-Developer/caseTeamB_server/model"
)

var db = model.GetDBConn()
var approvalM = new(sync.Mutex)

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

// 目標達成数変更
// Todo BOCCOAPI追記
func ApprovalGoal(device_id string, approval int) bool {
	approvalM.Lock()
	goal := model.GoalDate{}
	err := db.Where("device_id = ?", device_id).First(&goal).Error
	if err != nil {
		approvalM.Unlock()
		return false
	}

	goal.Run += approval
	if goal.Run < 0 {
		approvalM.Unlock()
		return false
	}
	err = db.Model(&goal).Update(&goal).Update("run", goal.Run).Error
	approvalM.Unlock()
	return true
}
