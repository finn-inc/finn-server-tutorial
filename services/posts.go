package services

import (
	"log"

	"github.com/finn-inc/finn-server-tutorial/models"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type PostsService struct {
	DB *gorm.DB
}

type Post struct {
	Id    string
	Title string
	Body  string
}

func (s *PostsService) modelToOutput(post models.Post) Post {
	return Post{
		Id:    post.Id,
		Title: post.Title,
		Body:  post.Body,
	}
}

func (s *PostsService) Index(page int, pageSize int) []Post {
	var posts []models.Post
	res := s.DB.Limit(pageSize).Offset((page - 1) * pageSize).Find(&posts)
	if res.Error != nil {
		log.Fatalf("Postデータを取得することができませんでした: %v\n", res.Error)
	}

	return lo.Map(posts, func(post models.Post, _ int) Post {
		return s.modelToOutput(post)
	})
}
