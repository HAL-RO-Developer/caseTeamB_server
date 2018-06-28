package service

import (
	"math/rand"
	"time"

	"github.com/HAL-RO-Developer/caseTeamB_server/model"
)

// デバイスID新規登録
func CreateDevice(name string, childId int) (string, bool) {
	if ExisByChildId(name, childId) == false {
		return "子どもIDが存在しません。", false
	}

	deviceID := createId()
	pin := createPin()

	device := model.Device{
		Name:     name,
		ChildId:  childId,
		DeviceId: deviceID,
		Pin:      pin,
	}
	err := db.Create(&device).Error
	if err != nil {
		return "デバイスIDが登録できませんでした。", false
	}
	return pin, true
}

// デバイスID作成
func createId() string {
	var deviceId string
	for {
		deviceId = createUuid(12, []rune("ABCDEFGHRJKLNMOPQRSTUPWXYZabcdefghijklmnopqrstuvwxyz0123456789"))
		_, find := ExisByDeviceId(deviceId)
		if !find {
			break
		}
	}
	return deviceId
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

// データベースからデバイス情報取得
func ExisByDeviceId(deviceId string) ([]model.Device, bool) {
	var devices []model.Device
	db.Where("device_id = ?", deviceId).Find(&devices)
	return devices, len(devices) != 0
}

// データベースからPin検索
func ExisByPin(pin string) bool {
	var devices []model.Device
	db.Where("pin = ?", pin).Find(&devices)
	return len(devices) != 0
}

func ExisByMac(mac string) bool {
	var devices []model.Device
	db.Where("mac = ?", mac).Find(&devices)
	return len(devices) != 0
}

// データベースからデバイステーブルの情報取得(ユーザー名から)
func GetDeviceId(name string) ([]model.Device, bool) {
	var devices []model.Device
	db.Where("name = ?", name).Find(&devices)
	return devices, len(devices) != 0
}

// 指定されたボタンIDの削除
func DeleteButtonId(name string, buttonId string) bool {
	var devices model.Device
	db.Where("name = ? and device_id = ?", name, buttonId).First(&devices)
	if devices.DeviceId == "" {
		return false
	}

	db.Delete(devices)
	return true
}

// 1ユーザーの最初のボタンIDの削除
func DeleteButtonFirst(name string) bool {
	var devices model.Device
	err := db.Where("name = ?", name).First(&devices).Error
	if err != nil {
		return false
	}
	db.Delete(devices)
	return true
}
