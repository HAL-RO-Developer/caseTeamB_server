package service

import (
	"strconv"

	"github.com/HAL-RO-Developer/caseTeamB_server/model"
)

// 目標承認
func ApprovalGoal(button_id string) error {
	goal := model.Goal{}
	err := db.Where("button_id = ?", button_id).First(&goal).Error
	if err != nil {
		return err
	}

	number, _ := strconv.Atoi(goal.Run)
	err = db.Model(&goal).Update(&goal).Update("approval", strconv.Itoa(number)).Error
	return err
}

// 目標非承認
func NotApprovalGoal(button_id string) error {
	goal := model.Goal{}
	err := db.Where("button_id = ?", button_id).First(&goal).Error
	if err != nil {
		return err
	}

	number, _ := strconv.Atoi(goal.Approval)
	err = db.Model(&goal).Update(&goal).Update("run", strconv.Itoa(number)).Error
	return err
}
