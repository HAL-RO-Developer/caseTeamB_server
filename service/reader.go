package service

import "github.com/HAL-RO-Developer/caseTeamB_server/model"

// ユーザーの回答データ送信
func SendUserAnswer(device_id string, uuid string) bool {
	deviceInfo, find := GetDeviceInfoFromDeviceId(device_id)
	if !find {
		return false
	}

	//tagInfo, find := GetTagDataFromTagId(uuid)
	record := model.Record{
		Name:    deviceInfo[0].Name,
		ChildId: deviceInfo[0].ChildId,
		//BookId:
	}

	err := db.Create(&record).Error
	if err != nil {
		return false
	}
	return true
}
