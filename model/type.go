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

type Button struct {
	Model
	Name     string `json:"name"`
	ButtonId string `json:"buttonId"`
	PushOn   string `json:"pushOn"`
	Pin      string `json:"pin"`
}

type Goal struct {
	Model
	Name     string `json:"name"`
	Contents string `json:"contents"`
}

type Message struct {
	Model
	Name     string `json:"name"`
	ButtonId string `json:"buttonId"`
	Rule     string `json:"rule"`
	Message  string `json:"message"`
}
