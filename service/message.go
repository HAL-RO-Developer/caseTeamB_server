package service

import (
	"github.com/HAL-RO-Developer/caseTeamB_server/model"
)

// メッセージ新規登録
func RegistrationMessage(buttonId string, condition string, message string) error {
	registration := model.Message{
		ButtonId:  buttonId,
		Condition: condition,
		Message:   message,
	}
	err := db.Create(&registration).Error
	return err
}

// 目標取得
func ExisByButtonIdFromMessage(buttonId string) ([]model.Message, bool) {
	var messages []model.Message
	db.Where("button_id = ?", buttonId).Find(&messages)
	return messages, len(messages) != 0
}

// 目標削除
func DeleteMessage(buttonId string) bool {
	var message model.Message
	err := db.Where("button_id = ?", buttonId).First(&message).Error
	if err != nil {
		return false
	}
	db.Delete(message)
	return true
}
