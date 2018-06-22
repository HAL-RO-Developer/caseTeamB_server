package model

// ドリル情報
type Book struct {
	Model
	BookId  int `json:"book_id"`
	GenreId int `json:"genre_id"`
}

// ドリル分野
type Genre struct {
	Model
	GenreId   int    `json:"genre_id"`
	GenreName string `json:"genre_name"`
}

// ドリル問題情報
type Question struct {
	Model
	DeviceId   string `json:"device_id"` // 読み取り機のデバイスID
	QuestionNo int    `json:"q_no"`
	Sentence   string `json:"sentence"` // 問題文
	Correct    int    `json:"correct"`
}

// 回答の記録
type Record struct {
	Model
	DeviceId   string `json:"device_id"` // 読み取り機のデバイスID
	BookId     int    `json:"book_id"`
	QuestionNo int    `json:"q_no"`
	Answer     int    `json:"user_answer"` // ユーザーの回答
}
