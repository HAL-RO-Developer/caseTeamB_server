package model

import (
	"time"
)

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

type User struct {
	Model
	Name     string `json:"name" binding:"required"`
	Password string `json:"pass" binding:"required"`
}
type Token struct {
	Model
	UserID uint
	Body   string
}

type Recipi struct {
	Model
}
