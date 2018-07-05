package service

import (
	"fmt"

	"github.com/HAL-RO-Developer/caseTeamB_server/model"
)

// メッセージ新規登録
func RegistrationWorkMessage(registData model.WorkCustomMessage) error {
	messageId := CreateWorkMessageId()
	registration := model.WorkCustomMessage{
		Name:        registData.Name,
		ChildId:     registData.ChildId,
		MessageId:   messageId,
		MessageCall: registData.MessageCall,
		Conditions:  registData.Conditions,
		Message:     registData.Message,
	}

	err := db.Create(&registration).Error
	return err
}

// メッセージ削除
func DeleteWorkMessage(messageId string) bool {
	var data model.WorkCustomMessage
	err := db.Where("message_id = ?", messageId).First(&data).Error
	if err != nil {
		return false
	}
	db.Delete(data)
	return true
}

// メッセージ取得(本文の重複チェック)
func GetWorkMessageFromMessage(name string, childId int, message string) bool {
	var messages []model.WorkCustomMessage
	db.Where("name = ? and child_id = ? and message = ?", name, childId, message).Find(&messages)
	return len(message) != 0
}

// メッセージ削除(子どもごと)
func DeleteWorkMessageFromChild(goalId string, childId int) bool {
	var data model.WorkCustomMessage
	err := db.Where("name = ? and child_id = ?", goalId, childId).First(&data).Error
	if err != nil {
		return false
	}
	db.Delete(data)
	return true
}

// メッセージ取得
func GetWorkMessageFromName(name string) ([]model.WorkCustomMessage, bool) {
	var messages []model.WorkCustomMessage
	db.Where("name = ?", name).Find(&messages)
	return messages, len(messages) != 0
}

// データベースからメッセージ情報取得
func GetMessageInfoFromId(messageId string) ([]model.WorkCustomMessage, bool) {
	var messages []model.WorkCustomMessage
	db.Where("message_id = ?", messageId).Find(&messages)
	return messages, len(messages) != 0
}

// メッセージ取得
func GetWorkMessageFromNameChild(name string, childId int) ([]model.WorkCustomMessage, bool) {
	var messages []model.WorkCustomMessage
	db.Where("name = ? and child_id = ?", name, childId).Find(&messages)
	return messages, len(messages) != 0
}

// データベースからメッセージ情報取得
func GetMessageInfoFromSame(name string, childId int, condition int, messageCall int) ([]model.WorkCustomMessage, bool) {
	var messages []model.WorkCustomMessage
	db.Where("conditions = ?", condition).Find(&messages)
	fmt.Println(messages)
	return messages, len(messages) != 0
}

// データベースからメッセージ情報取得
func GetMessageInfoFromTrue(name string, childId int, result int) ([]model.WorkCustomMessage, bool) {
	var messages []model.WorkCustomMessage
	db.Where("name = ? and child_id = ? and condition = ?", name, childId, result).Find(&messages)
	return messages, len(messages) != 0
}

// メッセージID作成
func CreateWorkMessageId() string {
	var messageId string
	for {
		messageId = createUuid(12, []rune("ABCDEFGHRJKLNMOPQRSTUPWXYZabcdefghijklmnopqrstuvwxyz0123456789"))
		_, find := GetMessageInfoFromId(messageId)
		if !find {
			break
		}
	}
	return messageId
}
