package validation

import (
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/response"
	"github.com/gin-gonic/gin"
)

type Button struct {
	ButtonId string `json:"button_id"`
}

func ButtonCheck(c *gin.Context) (Button, bool) {
	var req Button
	err := c.BindJSON(&req)
	if err != nil {
		response.BadRequest(gin.H{"error": "ボタンIDが入力されていません。"}, c)
		return req, false
	}
	return req, true
}
