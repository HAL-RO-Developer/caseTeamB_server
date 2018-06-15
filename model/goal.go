package model

// ボタン
type Button struct {
	Model
	Name     string `json:"name"`
	ButtonId string `json:"button_id"`
	Pin      string `json:"pin"` // ボタン登録するための一時的な情報
	Mac      string `json:"mac"` // ボタンのmacアドレス
}

// 目標
type Goal struct {
	Model
	ButtonId string `json:"button_id"`
	Contents string `json:"goal"`  // 目標内容
	Run      int    `json:"run"`   // 実行回数
	Value    int    `json:"value"` // 目標達成基準
}

// 登録メッセージ
type Message struct {
	Model
	ButtonId  string `json:"button_id"`
	Condition int    `json:"condition"` // メッセージ発信条件
	Message   string `json:"message"`   // メッセージ内容
}
