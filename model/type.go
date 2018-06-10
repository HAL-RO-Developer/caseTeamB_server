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
	ButtonId string `json:"button_id"`
	Pin      string `json:"pin"`
	Mac      string `json:"mac"`
}

type Goal struct {
	Model
	ButtonId  string `json:"button_id"`
	Contents  string `json:"goal"`
	Run       string `json:"run"`
	Apporoval string `json:"approval"`
}

type Message struct {
	Model
	ButtonId string `json:"button_id"`
	Rule     string `json:"rule"`
	Message  string `json:"message"`
}
