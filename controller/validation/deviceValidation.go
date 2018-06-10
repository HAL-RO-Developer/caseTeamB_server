package validation

import (
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/response"
	"github.com/gin-gonic/gin"
)

type PostDevice struct {
	Pin string `json:"pin"`
	Mac string `json:"mac"`
}

func ButtonRegistrationCheck(c *gin.Context) (PostDevice, bool) {
	var req PostDevice
	err := c.BindJSON(&req)
	if err != nil {
		response.BadRequest(gin.H{"error": "pinもしくはmacアドレスが未入力です。"}, c)
		return req, false
	}
	return req, true
}
