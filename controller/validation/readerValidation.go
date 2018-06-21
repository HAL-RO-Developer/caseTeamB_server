package validation

import (
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/response"
	"github.com/gin-gonic/gin"
)

type Tag struct {
	DeviceId string `json:"device_id"` // 読み取り機のmacアドレス
	Data     []Info `json:"data"`
}

type Info struct {
	BookId     int `json:"book_id"`
	QuestionNo int `json:"q_no"`
	Solution   int `json:"solution"`
}

func ReaderValidation(c *gin.Context) (Tag, bool) {
	var req Tag
	err := c.BindJSON(&req)
	if err != nil || len(req.Data) == 0 || req.DeviceId == "" {
		response.BadRequest(gin.H{"error": "入力されていないデータがあります。"}, c)
		return req, false
	}

	return req, true
}
