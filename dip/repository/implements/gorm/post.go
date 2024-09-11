package gorm

import (
	"fmt"

	"github.com/finn-inc/finn-server-tutorial/dip/models"
	"github.com/finn-inc/finn-server-tutorial/dip/registry"
	"github.com/finn-inc/finn-server-tutorial/dip/repository"
	"github.com/samber/lo"
)

type GormPostsRepository struct {
	DB *registry.DB
}

func NewGormPostsRepository(db *registry.DB) GormPostsRepository {
	return GormPostsRepository{
		DB: db,
	}
}

func (r GormPostsRepository) modelToOutput(post models.Post) repository.Post {
	return repository.Post{
		Id:    post.Id,
		Title: post.Title,
		Body:  post.Body,
	}
}

func (r GormPostsRepository) Index(page int, pageSize int) ([]repository.Post, error) {
	var posts []models.Post
	res := r.DB.Client.Limit(pageSize).Offset((page - 1) * pageSize).Find(&posts)
	if res.Error != nil {
		return nil, fmt.Errorf("error on getting posts: %v", res.Error)
	}

	return lo.Map(posts, func(post models.Post, _ int) repository.Post {
		return r.modelToOutput(post)
	}), nil
}

func (r GormPostsRepository) createInputToModel(input repository.CreatePostInput) models.Post {
	return models.Post{
		Id:    input.Id,
		Title: input.Title,
		Body:  input.Body,
	}
}

func (r GormPostsRepository) Create(input repository.CreatePostInput) error {
	post := r.createInputToModel(input)
	res := r.DB.Client.Create(&post)
	if res.Error != nil {
		return fmt.Errorf("error on creating post: %v", res.Error)
	}

	return nil
}
