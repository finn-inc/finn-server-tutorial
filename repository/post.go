package repository

type Post struct {
	Id    string
	Title string
	Body  string
}

type PostRepository interface {
	Index(int, int) []Post
}
