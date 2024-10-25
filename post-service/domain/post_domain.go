package domain

type Post struct {
	Id			int32		`json:"id"`
	Title		string		`json:"title"`
	Content		string		`json:"content"`
}