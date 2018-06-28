package service

import (
	"github.com/HAL-RO-Developer/caseTeamB_server/model"
)

// メッセージ新規登録
func RegistrationMessage(goalId string, condition int, message string) error {
	goalData, _ := GetOneGoal(goalId)

	registration := model.CustomMessage{
		Name:      goalData.Name,
		ChildId:   goalData.ChildId,
		GoalId:    goalId,
		Condition: condition,
		Message:   message,
	}
	err := db.Create(&registration).Error
	return err
}

// 目標取得
func GetMessageFromName(name string) ([]model.CustomMessage, bool) {
	var messages []model.CustomMessage
	db.Where("name = ?", name).Find(&messages)
	return messages, len(messages) != 0
}
