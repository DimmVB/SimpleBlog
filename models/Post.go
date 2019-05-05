package models

//Post(Id, Title, Content string)
type Post struct {
	ID      string
	Title   string
	Content string
}

//NewPost(id, title, content string)
func NewPost(id, title, content string) *Post {
	return &Post{id, title, content}
}
