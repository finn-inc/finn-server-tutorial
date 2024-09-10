package services

import (
	"fmt"

	"github.com/finn-inc/finn-server-tutorial/layered/models"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type PostsService struct {
	client *gorm.DB
}

func NewPostsService(db *gorm.DB) PostsService {
	return PostsService{
		client: db,
	}
}

type Post struct {
	Id    string
	Title string
	Body  string
}

type CreatePostInput struct {
	Title string
	Body  string
}

func (s PostsService) modelToOutput(post models.Post) Post {
	return Post{
		Id:    post.Id,
		Title: post.Title,
		Body:  post.Body,
	}
}

func (s PostsService) Index(page int) ([]Post, error) {
	pageSize := 10

	var posts []models.Post
	res := s.client.Limit(pageSize).Offset((page - 1) * pageSize).Find(&posts)
	if res.Error != nil {
		return nil, fmt.Errorf("error on getting posts: %v", res.Error)
	}

	return lo.Map(posts, func(post models.Post, _ int) Post {
		return s.modelToOutput(post)
	}), nil
}

func (s PostsService) createInputToModel(input CreatePostInput) models.Post {
	return models.Post{
		Id:    uuid.New().String(),
		Title: input.Title,
		Body:  input.Body,
	}
}

func (s PostsService) Create(input CreatePostInput) error {
	post := s.createInputToModel(input)
	res := s.client.Create(&post)
	if res.Error != nil {
		return fmt.Errorf("error on creating post: %v", res.Error)
	}

	return nil
}
