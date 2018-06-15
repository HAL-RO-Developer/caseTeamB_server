package validation

import (
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/response"
	"github.com/gin-gonic/gin"
)

type Approval struct {
	ButtonId string `json:"button_id"`
	Approval int    `json:"approval"`
}

func ApprovalCheck(c *gin.Context) (Approval, bool) {
	var req Approval
	err := c.BindJSON(&req)
	if err != nil {
		response.BadRequest(gin.H{"error": "ボタンIDが入力されていません。"}, c)
		return req, false
	}
	return req, true
}
