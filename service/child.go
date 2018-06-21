package service

import (
	"time"

	"github.com/HAL-RO-Developer/caseTeamB_server/controller/validation"
	"github.com/HAL-RO-Developer/caseTeamB_server/model"
)

// 子供ID検索
func FindByUserName(name string) ([]model.UserChild, bool) {
	var children []model.UserChild
	db.Where("name = ?", name).Find(&children)
	return children, len(children) != 0
}

// 子供情報追加
func AddChild(name string, info validation.UserChildren) bool {
	var i int
	buf, find := FindByUserName(name)
	if !find {
		i = 1
	}

	i = len(buf) + 1
	t, err := time.Parse("2006-01-02", info.BirthDay)

	if err != nil {
		return false
	}
	child := model.UserChild{
		Name:     name,
		ChildId:  i,
		BirthDay: t,
		NickName: info.NickName,
		Sex:      info.Sex,
	}

	err = db.Create(&child).Error
	if err != nil {
		return false
	}
	return true
}
