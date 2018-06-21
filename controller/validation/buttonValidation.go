package validation

import (
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/response"
	"github.com/gin-gonic/gin"
)

type DeviceId struct {
	DeviceId string `json:"device_id"`
}

func ButtonCheck(c *gin.Context) (DeviceId, bool) {
	var req DeviceId
	err := c.BindJSON(&req)
	if err != nil || req.DeviceId == "" {
		response.BadRequest(gin.H{"error": "デバイスIDが入力されていません。"}, c)
		return req, false
	}
	return req, true
}
