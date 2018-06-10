package validation

import (
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/response"
	"github.com/gin-gonic/gin"
)

type Goal struct {
	ButtonId string `json:"button_id"`
	Contents string `json:"goal"`
}

func GoalRegistrationCheck(c *gin.Context) (Goal, bool) {
	var req Goal
	err := c.BindJSON(&req)
	if err != nil {
		response.BadRequest(gin.H{"error": "button_idもしくはcontentsが未入力です。"}, c)
		return req, false
	}
	return req, true
}
