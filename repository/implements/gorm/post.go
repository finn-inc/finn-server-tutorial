package gorm

import (
	"log"

	"github.com/finn-inc/finn-server-tutorial/models"
	"github.com/finn-inc/finn-server-tutorial/registry"
	"github.com/finn-inc/finn-server-tutorial/repository"
	"github.com/samber/lo"
)

type GormPostsRepository struct {
	DB registry.DB
}

func NewGormPostsRepository(db registry.DB) GormPostsRepository {
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

func (r GormPostsRepository) Index(page int, pageSize int) []repository.Post {
	var posts []models.Post
	res := r.DB.Client.Limit(pageSize).Offset((page - 1) * pageSize).Find(&posts)
	if res.Error != nil {
		log.Fatalf("Postデータを取得することができませんでした: %v\n", res.Error)
	}

	return lo.Map(posts, func(post models.Post, _ int) repository.Post {
		return r.modelToOutput(post)
	})
}
