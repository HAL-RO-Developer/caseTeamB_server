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
	api.DELETE("/button/:device_id", Button.DeleteButton)

	// ボタン登録、プッシュ回数追加
	api.POST("/device", Device.DeviceRegistration)
	api.PUT("/device", Device.DeviceIncrement)

	// 目標登録、取得、削除
	api.POST("/goal", Goal.CreateGoal)
	api.GET("/goal/:button_id", Goal.GetGoal)
	api.DELETE("/goal/:button_id", Goal.DeleteGoal)

	// 目標達成承認、非承認
	api.PUT("/approval", Approval.ApprovalGoal)
	api.DELETE("/approval/:button_id", Approval.NotApprovalGoal)

	// メッセージ登録、取得、削除
	api.POST("/message", Message.NewMessage)
	api.GET("/message/:button_id", Message.GetMessage)
	api.DELETE("/message/:button_id", Message.DeleteMessage)
}
