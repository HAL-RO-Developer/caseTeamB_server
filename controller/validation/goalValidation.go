package validation

import (
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/response"
	"github.com/gin-gonic/gin"
)

type Goal struct {
	DeviceId string `json:"device_id"`
	Content  string `json:"goal"`
}

func GoalRegistrationCheck(c *gin.Context) (Goal, bool) {
	var req Goal
	err := c.BindJSON(&req)
	if err != nil || req.DeviceId == "" || req.Content == "" {
		response.BadRequest(gin.H{"error": "device_idもしくはcontentsが未入力です。"}, c)
		return req, false
	}

	return req, true
}
