package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/makki0205/gojwt"
	"github.com/HAL-RO-Developer/caseTeamB_server/controller"
	"github.com/HAL-RO-Developer/caseTeamB_server/model"
	"github.com/HAL-RO-Developer/caseTeamB_server/service"
)

func Jwt(salt string, exp int) gin.HandlerFunc {
	jwt.SetSalt(salt)
	jwt.SetExp(exp)
	return func(c *gin.Context) {
		token := c.Query("token")
		claims, err := jwt.Decode(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"err": err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("user_id", claims["id"])
		c.Set("email", claims["email"])
		c.Next()
	}
}

func Login(c *gin.Context) {
	var req model.User
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	user, ok := service.User.Login(req.Name, req.Password)
	if !ok {
		controller.BatRequest("ログイン失敗", c)
	}
	claims := map[string]string{
		"id":    strconv.Itoa(int(user.ID)),
		"name": user.Name,
	}
	token := jwt.Generate(claims)
	controller.Json(gin.H{"token": token}, c)
}
