package router

import (
	. "github.com/HAL-RO-Developer/caseTeamB_server/controller"
	"github.com/HAL-RO-Developer/caseTeamB_server/middleware"
	"github.com/gin-gonic/gin"
)

func userRouter(user *gin.RouterGroup) {
	// ユーザー登録、サインアップ
	user.POST("/signup", User.Create)
	user.POST("/signin", middleware.Login)
	// 子ども情報の登録、取得、削除
	user.POST("/child", User.Child)
	user.GET("/child", User.GetChildren)
	user.DELETE("/child", User.DeleteChild)
	// デバイスID発行、取得、削除
	user.POST("/device", Device.CreateNewDevice)
	user.GET("/device", Device.ListDevice)
	user.DELETE("/device/:device_id", Device.DeleteDevice)

	// デバイスID紐付け
	user.POST("/registration", Device.DeviceRegistration)
}

func workRouter(work *gin.RouterGroup) {
	// ユーザー情報削除
	work.DELETE("/user")
	// ICリーダー
	work.POST("/reader", Reader.SendTag)
	// 回答記録取得
	work.GET("/record/:device_id", Record.WorkRecord)

}

func goalRouter(goal *gin.RouterGroup) {
	// ユーザー情報削除
	goal.DELETE("/user", User.UserDeleteForGoal)

	// プッシュ回数変更
	goal.POST("/push", Button.DeviceIncrement)

	// 目標登録、取得、削除
	goal.POST("/goal", Goal.CreateGoal)
	goal.GET("/goal/:device_id", Goal.GetGoal)
	goal.DELETE("/goal/:device_id", Goal.DeleteGoal)

	// 目標達成操作
	goal.PUT("/approval", Approval.ApprovalGoal)

	// メッセージ登録、取得、削除
	goal.POST("/message", Message.NewMessage)
	goal.GET("/message/:device_id", Message.GetMessage)
	goal.DELETE("/message/:device_id", Message.DeleteMessage)
}
