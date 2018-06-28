package validation

import (
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/response"
	"github.com/gin-gonic/gin"
)

type Approval struct {
	GoalId   string `json:"goal_id"`
	Approval int    `json:"approval"`
}

func ApprovalCheck(c *gin.Context) (Approval, bool) {
	var req Approval
	err := c.BindJSON(&req)
	if err != nil || req.GoalId == "" {
		response.BadRequest(gin.H{"error": "未入力の項目があります。"}, c)
		return req, false
	}
	return req, true
}
