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
	Mac        string `json:"mac"` // 読み取り機のmacアドレス
	QuestionNo int    `json:"q_no"`
	Sentence   string `json:"sentence"` // 問題文
	Correct    int    `json:"correct"`
}

// 回答の記録
type Record struct {
	Model
	Mac        string `json:"mac"` // 読み取り機のmacアドレス
	BookId     int    `json:"book_id"`
	QuestionNo int    `json:"q_no"`
	Answer     int    `json:"answer"` // ユーザーの回答
}
