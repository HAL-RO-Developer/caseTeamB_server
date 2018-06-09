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

	// ボタン登録、プッシュ回数追加
	api.POST("/device", Device.DeviceRegistration)
	api.PUT("/device", Device.DeviceIncrement)

	// 目標登録、取得、削除
	api.POST("/goal", Goal.CreateGoal)
	api.GET("/goal")
	api.DELETE("/goal")

}
