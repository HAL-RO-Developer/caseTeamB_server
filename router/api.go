package router

import (
	. "github.com/HAL-RO-Developer/caseTeamB_server/controller"
	"github.com/HAL-RO-Developer/caseTeamB_server/middleware"
	"github.com/gin-gonic/gin"
)

func apiRouter(api *gin.RouterGroup) {
	//api.OPTIONS("/signup", preflightRequest)
	api.POST("/signup", User.Create)
	//api.OPTIONS("/signin", preflightRequest)
	api.POST("/signin", middleware.Login)

}

func authApiRouter(auth *gin.RouterGroup) {
	auth.GET("/hoge", func(c *gin.Context) {
		c.JSON(200, "ok")
	})
}
