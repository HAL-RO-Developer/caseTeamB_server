package model

import "time"

// 目標
type GoalData struct {
	Model
	Name     string     `json:"name"`
	ChildId  int        `json:"child_id"`
	GoalId   string     `json:"goal_id"`
	DeviceId string     `json:"device_id"`
	Content  string     `json:"content"`  // 目標内容
	Run      int        `json:"run"`      // 実行回数
	Criteria int        `json:"criteria"` // 目標達成基準
	Deadline *time.Time `json:"deadline"` // 目標達成期限
}

// 編集メッセージ
type CustomMessage struct {
	Model
	Name      string `json:"name"`
	ChildId   int    `json:"child_id"`
	GoalId    string `json:"goal_id"`
	Condition int    `json:"condition"` // メッセージ発信条件
	Message   string `json:"message"`   // メッセージ内容
}
