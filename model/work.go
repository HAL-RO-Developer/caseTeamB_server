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
	BookId     int    `json:"book_id"`
	QuestionNo int    `json:"q_no"`
	Sentence   string `json:"sentence"` // TagId
	Correct    string `json:"correct"`  // TagId
}

// 回答の記録
type Record struct {
	Model
	Name       string `json:"name"`
	ChildId    int    `json:"child_id"`
	BookId     int    `json:"book_id"`
	QuestionNo int    `json:"q_no"`
	UserAnswer string `json:"user_answer"` // ユーザーの回答
	Challenge  int    `json:"challenge"`   // チャレンジ実行タイプ
}

// タグ情報
type Tag struct {
	Model
	TagId      string `json:"tag_id"`
	Uuid       string `json:"uuid"`
	BookId     int    `json:"book_id"`
	QuestionNo int    `json:"q_no"`
	Sentence   string `json:"sentence"`
	Answer     string `json:"answer"`
}
