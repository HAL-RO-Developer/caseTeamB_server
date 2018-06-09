package middleware

import (
	"net/http"

	"github.com/HAL-RO-Developer/caseTeamB_server/controller"
	"github.com/HAL-RO-Developer/caseTeamB_server/model"
	"github.com/HAL-RO-Developer/caseTeamB_server/service"
	"github.com/gin-gonic/gin"
	"github.com/makki0205/gojwt"
)

func Login(c *gin.Context) {
	var req model.User
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	user, ok := service.User.Login(req.Name, req.Password)
	if !ok {
		controller.BadRequest("ログインエラー", c)
		return
	}
	claims := map[string]string{
		"name": user.Name,
		"pass": user.Password,
	}
	token := jwt.Generate(claims)

	controller.Json(gin.H{"token": token}, c)
}
