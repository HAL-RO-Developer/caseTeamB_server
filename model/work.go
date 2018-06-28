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
	BookId     string `json:"book_id"`
	QuestionNo int    `json:"q_no"`
	Sentence   string `json:"sentence"` // 問題文
	Correct    int    `json:"correct"`  // 答え
}

// 回答の記録
type Record struct {
	Model
	Name       string `json:"name"`
	ChildId    int    `json:"child_id"`
	BookId     int    `json:"book_id"`
	QuestionNo int    `json:"q_no"`
	Answer     int    `json:"user_answer"` // ユーザーの回答
	Challenge  int    `json:"challege"`    // チャレンジ実行タイプ
}
