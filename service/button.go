package service

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/HAL-RO-Developer/caseTeamB_server/model"
)

var db = model.GetDBConn()

// ボタンID新規登録
func CreateButton(name string) (string, error) {
	deviceID := createId()
	pin := createPin()

	button := model.Button{
		Name:     name,
		ButtonId: deviceID,
		PushOn:   "0",
		Pin:      pin,
	}
	err := db.Create(&button).Error
	return pin, err
}

// ボタンID作成
func createId() string {
	var buttonId string
	for {
		buttonId = createUuid(12, []rune("ABCDEFGHRJKLNMOPQRSTUPWXYZabcdefghijklmnopqrstuvwxyz0123456789"))
		if !ExisByButtonId(buttonId) {
			break
		}
	}
	return buttonId
}

// ランダム文字列作成
func createUuid(length int, letters []rune) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// ピン作成
func createPin() string {
	var pin string
	for {
		pin = createUuid(4, []rune("0123456789"))
		if !ExisByPin(pin) {
			break
		}
	}
	return pin
}

// データベースからボタンID検索
func ExisByButtonId(buttonId string) bool {
	var buttons []model.Button
	db.Where("button_id = ?", buttonId).Find(&buttons)
	return len(buttons) != 0
}

// データベースからPin検索
func ExisByPin(pin string) bool {
	var buttons []model.Button
	db.Where("pin = ?", pin).Find(&buttons)
	return len(buttons) != 0
}

// データベースからボタンID一覧取得(ユーザー名から)
func GetButtonId(name string) ([]model.Button, bool) {
	var buttons []model.Button
	db.Where("name = ?", name).Find(&buttons)
	return buttons, len(buttons) != 0
}

// 指定されたボタンIDの削除
func DeleteButtonId(name string, buttonId string) bool {
	var buttons model.Button
	db.Where("name = ? and button_id = ?", name, buttonId).First(&buttons)
	if buttons.ButtonId == "" {
		fmt.Println("fail")
		return false
	}

	db.Delete(buttons)
	return true
}

// 1ユーザーの最初"のボタンIDの削除
func DeleteButtonFirst(name string) {
	var buttons model.Button
	db.Where("name = ?", name).First(&buttons)
	db.Delete(buttons)
}
