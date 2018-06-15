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
}

func workRouter(work *gin.RouterGroup) {

}

func goalRouter(goal *gin.RouterGroup) {
	// ユーザー情報削除
	goal.DELETE("/user", User.UserDeleteForGoal)

	// ユーザー情報追加登録

	// ボタンID発行、取得、削除
	goal.POST("/button", Button.CreateNewButton)
	goal.GET("/button", Button.ListButton)
	goal.DELETE("/button/:device_id", Button.DeleteButton)

	// ボタン登録、プッシュ回数追加
	goal.POST("/device", Device.DeviceRegistration)
	goal.PUT("/device", Device.DeviceIncrement)

	// 目標登録、取得、削除
	goal.POST("/goal", Goal.CreateGoal)
	goal.GET("/goal/:button_id", Goal.GetGoal)
	goal.DELETE("/goal/:button_id", Goal.DeleteGoal)

	// 目標達成操作
	goal.PUT("/approval", Approval.ApprovalGoal)

	// メッセージ登録、取得、削除
	goal.POST("/message", Message.NewMessage)
	goal.GET("/message/:button_id", Message.GetMessage)
	goal.DELETE("/message/:button_id", Message.DeleteMessage)
}
