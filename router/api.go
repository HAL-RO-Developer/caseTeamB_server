package router

import (
	. "github.com/HAL-RO-Developer/caseTeamB_server/controller"
	"github.com/HAL-RO-Developer/caseTeamB_server/middleware"
	"github.com/gin-gonic/gin"
)

func apiRouter(api *gin.RouterGroup) {
	// ユーザー登録、サインアップ
	api.POST("/signup", User.Create)
	api.POST("/signin", middleware.Login)
	api.DELETE("/user", User.UserDelete)

	// ボタンID発行、取得、削除
	api.POST("/button", Button.CreateNewButton)
	api.GET("/button", Button.ListButton)
	api.DELETE("/button", Button.DeleteButton)

}

func authApiRouter(auth *gin.RouterGroup) {
	auth.GET("/hoge", func(c *gin.Context) {
		c.JSON(200, "ok")
	})
}
