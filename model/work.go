package model

// ドリル情報
type Book struct {
	Model
	BookId  string `json:"book_id"`
	GenreId int    `json:"genre_id"`
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
	Sentence   string `json:"sentence"`
	Answer     string `json:"answer"`
}

// 回答の記録
type Record struct {
	Model
	Name       string `json:"name" binding:"required"`
	BookId     string `json:"book_id"`
	QuestionNo int    `json:"q_no"`
	Reply      string `json:"reply"`
	Correct    bool   `json:"correct"`
}
