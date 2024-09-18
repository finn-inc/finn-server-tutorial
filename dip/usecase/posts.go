package usecase

import (
	"fmt"

	"github.com/finn-inc/finn-server-tutorial/dip/models"
	"github.com/finn-inc/finn-server-tutorial/dip/repository"
	"github.com/google/uuid"
)

type PostsUsecase struct {
	r repository.PostsRepository
}

func NewPostsUsecase(r repository.PostsRepository) *PostsUsecase {
	return &PostsUsecase{
		r: r,
	}
}

type CreatePostInput struct {
	Title string
	Body  string
}

func (i *CreatePostInput) toModel(id string) models.Post {
	return models.Post{
		Id:    id,
		Title: i.Title,
		Body:  i.Body,
	}
}

type Post struct {
	Id    string
	Title string
	Body  string
}

func (u *PostsUsecase) Index(page int) ([]models.Post, error) {
	posts, err := u.r.Index(page, 10)
	if err != nil {
		return nil, fmt.Errorf("error on indexing posts: %v", err)
	}

	return posts, nil
}

func (u PostsUsecase) Create(input CreatePostInput) error {
	post := input.toModel(uuid.New().String())
	if err := u.r.Save(post); err != nil {
		return fmt.Errorf("error on create usecase: %v", err)
	}

	return nil
}
