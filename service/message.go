package service

import (
	"github.com/HAL-RO-Developer/caseTeamB_server/model"
)

// メッセージ新規登録
func RegistrationMessage(goalId string, messageCall int, message string) error {
	goalData, _ := GetOneGoal(goalId)

	registration := model.CustomMessage{
		Name:        goalData.Name,
		ChildId:     goalData.ChildId,
		GoalId:      goalId,
		MessageCall: messageCall,
		Message:     message,
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

// 目標取得
func GetMessageFromNameChild(name string, childId int) ([]model.CustomMessage, bool) {
	var messages []model.CustomMessage
	db.Where("name = ? and child_id = ?", name, childId).Find(&messages)
	return messages, len(messages) != 0
}

// メッセージ取得
func GetMessageFromGoal(goalId string) ([]model.CustomMessage, bool) {
	var messages []model.CustomMessage
	db.Where("goal_id = ?", goalId).First(&messages)
	return messages, len(messages) != 0
}

// メッセージ重複確認
func ExisMessageFromGoal(goalId string, messageCall int) bool {
	var messages []model.CustomMessage
	db.Where("goal_id = ? and message_call = ?", goalId, messageCall).Find(&messages)

	return len(messages) != 0
}

// メッセージ
func UpdateMessage(goalId string, messageCall int, message string) error {
	var data model.CustomMessage
	err := db.Where("goal_id = ? and message_call = ?", goalId, messageCall).First(&data).Error
	if err != nil {
		return err
	}
	err = db.Model(&data).Update(&data).Update("message", message).Error
	return err
}
