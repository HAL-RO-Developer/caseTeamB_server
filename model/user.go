package model

import "time"

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

// ユーザー基本情報
type User struct {
	Model
	Name     string `json:"name" binding:"required"`
	Password string `json:"pass" binding:"required"`
}

// ユーザー任意情報
type UserOption struct {
	Model
	Name  string `json:"name" binding:"required"`
	Year  uint   `json:"year"`
	Month uint   `json:"month"`
	Date  uint   `json:"date"`
	Mac   string `json:"mac"` // チームA用
}
