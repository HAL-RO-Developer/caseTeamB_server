package controller

import (
	"github.com/HAL-RO-Developer/caseTeamB_server/model"
	"github.com/HAL-RO-Developer/caseTeamB_server/service"
	"github.com/gin-gonic/gin"
)

var User = userimpl{}

type userimpl struct {
}

func (u *userimpl) Create(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		BatRequest(err.Error(), c)
		return
	}

	if service.User.ExisByName(user.Name) {
		Json("登録済みユーザーです", c)
	} else {
		user = service.User.Store(user)
		Json("ユーザー登録を行いました。", c)
	}
}
