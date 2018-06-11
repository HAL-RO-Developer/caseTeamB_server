package validation

import (
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/response"
	"github.com/gin-gonic/gin"
)

type Message struct {
	ButtonId  string `json:"button_id"`
	Condition string `json:"condition"`
	Message   string `json:"message"`
}

func MessageValidation(c *gin.Context) (Message, bool) {
	var req Message
	err := c.BindJSON(&req)
	if err != nil {
		response.BadRequest(gin.H{"error": "入力されていないデータがあります。"}, c)
		return req, false
	}
	return req, true
}
