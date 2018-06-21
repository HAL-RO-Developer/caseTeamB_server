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
func DeleteUser(name string) bool {
	var user model.User
	err := db.Where("name = ?", name).First(&user).Error
	if err != nil {
		return false
	}
	db.Delete(&user)
	return true
}

// こどもID検索
func ExisByChildId(name string, childId int) bool {
	var child []model.UserChild
	db.Where("name = ? and child_id = ?", name, childId).Find(&child)
	return len(child) != 0
}

// 最初に見つけたこどもID情報削除
func DeleteChildFirst(name string) bool {
	var child model.UserChild
	err := db.Where("name = ?", name).First(&child).Error
	if err != nil {
		return false
	}
	db.Delete(child)
	return true
}

// こどもID情報削除
func DeleteChild(name string, childId int) bool {
	var child model.UserChild
	err := db.Where("name = ? and child_id = ?", name, childId).First(&child).Error
	if err != nil {
		return false
	}
	db.Delete(&child)
	return true
}
