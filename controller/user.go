package controller

import (
	"github.com/HAL-RO-Developer/caseTeamB_server/model"
	"github.com/HAL-RO-Developer/caseTeamB_server/service"
	"github.com/gin-gonic/gin"
	jwt "github.com/makki0205/gojwt"
)

var User = userimpl{}

type userimpl struct {
}

func (u *userimpl) Create(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		BadRequest("ユーザー名またはパスワードが未入力です。", c)
		return
	}

	if service.User.ExisByName(user.Name) {
		BadRequest("登録済みのユーザー名です。", c)
	} else {
		user = service.User.Store(user)
		Json(gin.H{"success": "ユーザー登録を行いました。"}, c)
	}
}

// TODO 目標関連処理実装後追記
func (u *userimpl) UserDelete(c *gin.Context) {
	name, ok := authorizationCheck(c)
	if !ok {
		BadRequest("ログインエラー", c)
		return
	}

	buttons, find := service.GetButtonId(name)
	if find {
		for i := 0; i < len(buttons); i++ {
			service.DeleteButtonFirst(name)
		}
	}
	service.DeleteUser(name)
	Json(gin.H{"success": "ユーザー情報を削除しました。"}, c)
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
