package service

import (
	"github.com/HAL-RO-Developer/caseTeamB_server/model"
	"github.com/satori/go.uuid"
	_ "github.com/satori/go.uuid"
)

// ユーザーの回答データ送信
func SendUserAnswer(deviceId string, tagUuid string) int {
	deviceInfo, find := GetDeviceInfoFromDeviceId(deviceId)
	if !find {
		return -1
	}

	tagInfo := GetTagDataFromUuid(tagUuid)
	if tagInfo == nil {
		return -1
	}

	// 問題タグ時
	if tagInfo.Sentence != "" {
		bocco, find := GetDeviceInfoFromDeviceId(deviceId)
		if !find {
			return -1
		}
		boccoInfo, find := ExisByBoccoAPI(bocco[0].Name)
		if !find {
			return -2
		}
		boccoToken, _ := GetBoccoToken(boccoInfo[0].Email, boccoInfo[0].Key, boccoInfo[0].Pass)
		roomId, _ := GetRoomId(boccoToken)
		uuid := uuid.Must(uuid.NewV4()).String()
		SendMessage(uuid, roomId, boccoToken, tagInfo.Sentence)
	} else {
		record := model.Record{
			Name:       deviceInfo[0].Name,
			ChildId:    deviceInfo[0].ChildId,
			BookId:     tagInfo.BookId,
			QuestionNo: tagInfo.QuestionNo,
			UserAnswer: tagInfo.Answer,
		}
		err := db.Create(&record).Error
		if err != nil {
			return -1
		}
	}
	correct := GetQuestionData(tagInfo.BookId, tagInfo.QuestionNo)
	if tagInfo.TagId != correct[0].Correct {
		return 0
	}
	return 1
}

// チャレンジ判定
