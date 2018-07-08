package router

import (
	. "github.com/HAL-RO-Developer/caseTeamB_server/controller"
	"github.com/HAL-RO-Developer/caseTeamB_server/middleware"
	"github.com/gin-gonic/gin"
)

func apiRouter(api *gin.RouterGroup) {
	// ユーザー登録、サインアップ、削除
	api.POST("/signup", User.Create)
	api.POST("/signin", middleware.Login)
	api.DELETE("user", User.UserDeleteForGoal)
	// 子ども情報の登録、取得、削除
	api.POST("/child", User.Child)
	api.GET("/child", User.GetChildren)
	api.DELETE("/child/:child_id", User.DeleteChild)
	// デバイスID発行、取得、削除
	api.POST("/device", Device.CreateNewDevice)
	api.GET("/device", Device.ListDevice)
	api.DELETE("/device/:device_id", Device.DeleteDevice)

	// BOCCOAPI
	api.POST("/bocco", Bocco.RegistBocco)
	api.GET("/bocco", Bocco.GetBoccoInfo)
	api.DELETE("/bocco", Bocco.DeleteBoccoInfo)

	// 目標登録、取得、削除
	api.POST("/goal", Goal.CreateGoal)
	api.GET("/goal", Goal.GetGoal)
	api.DELETE("/goal/:goal_id", Goal.DeleteGoal)

	// 目標達成操作
	api.PUT("/approval", Approval.ApprovalGoal)

	// メッセージ登録、取得、削除
	api.POST("/message", Message.EditMessage)
	api.GET("/message", Message.GetMessage)
	api.DELETE("/message/:goal_id/:message_call", Message.DeleteMessage)
}

func thingRouter(thing *gin.RouterGroup) {
	// デバイスID紐付け
	thing.POST("/registration", Device.DeviceRegistration)
	// プッシュ回数プラス1
	thing.PUT("/button", Button.DeviceIncrement)
}
