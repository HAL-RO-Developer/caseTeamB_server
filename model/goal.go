package model

// 目標
type GoalDate struct {
	Model
	DeviceId string `json:"device_id"`
	Content  string `json:"content"` // 目標内容
	Run      int    `json:"run"`     // 実行回数
	Value    uint   `json:"value"`   // 目標達成基準
}

// 登録メッセージ
type CustomMessage struct {
	Model
	DeviceId  string `json:"device_id"`
	Condition int    `json:"condition"` // メッセージ発信条件
	Message   string `json:"message"`   // メッセージ内容
}
