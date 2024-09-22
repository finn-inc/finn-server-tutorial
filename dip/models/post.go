package models

import "github.com/google/uuid"

type Post struct {
	Id    string `gorm:"column:id"`
	Title string `gorm:"column:title"`
	Body  string `gorm:"column:body"`
}

type NewPostInput struct {
	Title string
	Body  string
}

func NewPost(input NewPostInput) Post {
	return Post{
		Id:    uuid.New().String(),
		Title: input.Title,
		Body:  input.Body,
	}
}
