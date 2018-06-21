package service

import (
	"github.com/HAL-RO-Developer/caseTeamB_server/model"
)

// メッセージ新規登録
func RegistrationMessage(deviceId string, condition int, message string) error {
	registration := model.CustomMessage{
		DeviceId:  deviceId,
		Condition: condition,
		Message:   message,
	}
	err := db.Create(&registration).Error
	return err
}

// 目標取得
func ExisByButtonIdFromMessage(goalId string) ([]model.CustomMessage, bool) {
	var messages []model.CustomMessage
	db.Where("goal_id = ?", goalId).Find(&messages)
	return messages, len(messages) != 0
}

// 目標削除
func DeleteMessage(goalId string) bool {
	var message model.CustomMessage
	err := db.Where("goal_id = ?", goalId).First(&message).Error
	if err != nil {
		return false
	}
	db.Delete(message)
	return true
}
