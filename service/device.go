package service

import (
	"strconv"

	"github.com/HAL-RO-Developer/caseTeamB_server/model"
)

var db = model.GetDBConn()

// macAddr登録
func RegistrationButton(pin string, mac string) (string, error) {
	button := model.Button{}
	err := db.Where("pin = ?", pin).First(&button).Error
	if err != nil {
		return "not found", err
	}

	button.Mac = mac
	err = db.Model(&button).Update(&button).Update("pin", "").Error
	return button.ButtonId, err
}

// プッシュ回数追加
// Todo BOCCOAPI追記
func IncrementButton(button_id string) error {
	button := model.Button{}
	err := db.Where("button_id = ?", button_id).First(&button).Error
	if err != nil {
		return err
	}

	number, _ := strconv.Atoi(button.PushOn)
	number++
	err = db.Model(&button).Update(&button).Update("push_on", strconv.Itoa(number)).Error
	return err
}
