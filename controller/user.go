package controller

import (
	"time"

	"github.com/HAL-RO-Developer/caseTeamB_server/controller/response"
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/validation"
	"github.com/HAL-RO-Developer/caseTeamB_server/model"
	"github.com/HAL-RO-Developer/caseTeamB_server/service"
	"github.com/gin-gonic/gin"
	"github.com/makki0205/gojwt"
)

var User = userimpl{}

type userimpl struct {
	ChildId  int       `json:"child_id"`
	BirthDay time.Time `json:"birthday"`
	NickName string    `json:"nickname"`
	Sex      int       `json:"sex"`
}

func (u *userimpl) Create(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		response.BadRequest(gin.H{"error": "ユーザー名またはパスワードが未入力です。"}, c)
		return
	}
	if service.User.ExisByName(user.Name) {
		response.BadRequest(gin.H{"error": "登録済みのユーザー名です。"}, c)
	} else {
		user = service.User.Store(user)
		response.Json(gin.H{"success": "ユーザー登録を行いました。"}, c)
	}
}

// ユーザー削除(チームA)
func (u *userimpl) UserDeleteForWork(c *gin.Context) {
	name, ok := authorizationCheck(c)
	if !ok {
		response.BadRequest(gin.H{"error": "ログインエラー"}, c)
		return
	}

	works, find := service.ExisByRecord(name)
	if find {
		for i := 0; i < len(works); i++ {
			service.DeleteUserAnswer(works[i].DeviceId)
		}
	}

	if service.DeleteUser(name) {
		response.Json(gin.H{"success": "ユーザー情報を削除しました。"}, c)
		return
	}
	response.BadRequest(gin.H{"error": "ユーザー情報が見つかりませんでした。"}, c)
}

// ユーザー削除(チームB)
func (u *userimpl) UserDeleteForGoal(c *gin.Context) {
	name, ok := authorizationCheck(c)
	if !ok {
		response.BadRequest(gin.H{"error": "ログインエラー"}, c)
		return
	}

	goals, find := service.GetDeviceId(name)
	if find {
		for i := 0; i < len(goals); i++ {
			service.DeleteGoal(goals[i].DeviceId)
			service.DeleteMessage(goals[i].DeviceId)
			service.DeleteButtonFirst(name)
			service.DeleteChildFirst(name)
		}
	}

	if service.DeleteUser(name) {
		response.Json(gin.H{"success": "ユーザー情報を削除しました。"}, c)
		return
	}
	response.BadRequest(gin.H{"error": "ユーザー情報が見つかりませんでした。"}, c)
}

// 子供情報取得
func (u *userimpl) GetChildren(c *gin.Context) {
	var children []userimpl
	var child userimpl

	name, ok := authorizationCheck(c)
	if !ok {
		response.BadRequest(gin.H{"error": "ログインエラー"}, c)
		return
	}

	buf, find := service.FindByUserName(name)

	if find {
		for i := 0; i < len(buf); i++ {
			child.ChildId = buf[i].ChildId
			child.BirthDay = buf[i].BirthDay
			child.NickName = buf[i].NickName
			child.Sex = buf[i].Sex
			children = append(children, child)
		}
		response.Json(gin.H{"data": children}, c)
		return
	}
	response.BadRequest(gin.H{"error": "子供情報が見つかりませんでした。"}, c)
}

// 子供情報追加
func (u *userimpl) Child(c *gin.Context) {
	name, ok := authorizationCheck(c)
	if !ok {
		response.BadRequest(gin.H{"error": "ログインエラー"}, c)
		return
	}

	req, ok := validation.AddChildInfoValidation(c)
	if !ok {
		return
	}

	success := service.AddChild(name, req)
	if !success {
		response.BadRequest(gin.H{"error": "登録失敗"}, c)
		return
	}
	response.Json(gin.H{"success": "子供の情報を追加しました。"}, c)
}

// 指定された子ども情報の削除
func (u *userimpl) DeleteChild(c *gin.Context) {
	name, ok := authorizationCheck(c)
	if !ok {
		response.BadRequest(gin.H{"error": "ログインエラー"}, c)
		return
	}

	req, ok := validation.ChildrenInfoValidation(c)
	if !ok {
		return
	}
	success := service.DeleteChild(name, req.ChildId)
	if !success {
		response.BadRequest(gin.H{"error": "子どもIDが見つかりませんでした。"}, c)
		return
	}
	response.Json(gin.H{"success": "削除しました。"}, c)
}

//	トークンの検証
func authorizationCheck(c *gin.Context) (string, bool) {
	token := c.GetHeader("Authorization")

	userInfo, err := jwt.Decode(token)
	if err != nil {
		return "error", false
	}

	_, ok := service.User.Login(userInfo["name"], userInfo["pass"])
	if !ok {
		return "error", false
	}

	return userInfo["name"], true
}
