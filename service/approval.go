package service

import (
	"github.com/HAL-RO-Developer/caseTeamB_server/model"
)

// 目標達成数変更
func ApprovalGoal(button_id string, approval int) bool {
	goal := model.Goal{}
	err := db.Where("button_id = ?", button_id).First(&goal).Error
	if err != nil {
		return false
	}

	goal.Run += approval
	if goal.Run < 0 {
		return false
	}
	err = db.Model(&goal).Update(&goal).Update("run", goal.Run).Error
	return true
}
