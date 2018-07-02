package router

import (
	. "github.com/HAL-RO-Developer/caseTeamB_server/controller"
	"github.com/HAL-RO-Developer/caseTeamB_server/middleware"
	"github.com/gin-gonic/gin"
)

func userRouter(user *gin.RouterGroup) {
	// ユーザー登録、サインアップ、削除
	user.POST("/signup", User.Create)
	user.POST("/signin", middleware.Login)
	user.DELETE("user", User.UserDeleteForGoal)
	// 子ども情報の登録、取得、削除
	user.POST("/child", User.Child)
	user.GET("/child", User.GetChildren)
	user.DELETE("/child/:child_id", User.DeleteChild)
	// デバイスID発行、取得、削除
	user.POST("/device", Device.CreateNewDevice)
	user.GET("/device", Device.ListDevice)
	user.DELETE("/device/:device_id", Device.DeleteDevice)

	// BOCCOAPI
	user.POST("/bocco", Bocco.RegistBocco)
	user.GET("/bocco", Bocco.GetBoccoInfo)
	user.DELETE("/bocco", Bocco.DeleteBoccoInfo)
}

func workRouter(work *gin.RouterGroup) {
	// 回答記録取得
	work.GET("/record/:device_id", Record.WorkRecord)
}

func goalRouter(goal *gin.RouterGroup) {
	// 目標登録、取得、削除
	goal.POST("/goal", Goal.CreateGoal)
	goal.PUT("/goal", Goal.UpdateGoal)
	goal.GET("/goal", Goal.GetGoal)
	goal.DELETE("/goal/:goal_id", Goal.DeleteGoal)

	// 目標達成操作
	goal.PUT("/approval", Approval.ApprovalGoal)

	// メッセージ登録、取得、削除
	goal.POST("/message", Message.EditMessage)
	goal.GET("/message", Message.GetMessage)
	goal.DELETE("/message/:goal_id/:message_call")
}

func thingRouter(thing *gin.RouterGroup) {
	// デバイスID紐付け
	thing.POST("/registration", Device.DeviceRegistration)
	// ICリーダー
	thing.POST("/reader", Reader.SendTag)
	// プッシュ回数プラス1
	thing.PUT("/button", Button.DeviceIncrement)
}
