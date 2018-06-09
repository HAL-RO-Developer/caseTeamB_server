package service

import "github.com/HAL-RO-Developer/caseTeamB_server/model"

var User = user{}

type user struct {
}

func (u *user) Store(user model.User) model.User {
	db.Create(&user)
	return user
}
func (u *user) ExisByEmail(email string) bool {
	var users []model.User
	db.Where("email = ?", email).Find(&users)
	return len(users) != 0
}

func (u *user) ExisByName(name string) bool {
	var users []model.User
	db.Where("name = ?", name).Find(&users)
	return len(users) != 0
}

func (u *user) Login(name, pass string) (*model.User, bool) {
	var users []model.User
	db.Where("name = ?", name).Find(&users)
	if len(users) == 0 {
		return nil, false
	}
	return &users[0], users[0].Password == pass
}

// ユーザー情報削除
func DeleteUser(name string) {
	var user model.User
	db.Where("name = ?", name).First(&user)
	db.Delete(&user)
}
