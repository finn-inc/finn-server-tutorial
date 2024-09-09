package usecase

import (
	"fmt"

	"github.com/finn-inc/finn-server-tutorial/hexagonal/repository"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type PostsUsecase struct {
	r repository.PostRepository
}

type CreatePostInput struct {
	Title string
	Body  string
}

func (i *CreatePostInput) toRepoInput(id string) repository.CreatePostInput {
	return repository.CreatePostInput{
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

func NewPostsUsecase(r repository.PostRepository) *PostsUsecase {
	return &PostsUsecase{
		r: r,
	}
}

func (u PostsUsecase) repoToOutput(post repository.Post) Post {
	return Post{
		Id:    post.Id,
		Title: post.Title,
		Body:  post.Body,
	}
}

func (u *PostsUsecase) Index() ([]Post, error) {
	posts, err := u.r.Index(1, 10)
	if err != nil {
		return nil, fmt.Errorf("error on indexing posts: %v", err)
	}

	return lo.Map(posts, func(post repository.Post, _ int) Post {
		return u.repoToOutput(post)
	}), nil
}

func (u PostsUsecase) Create(input CreatePostInput) error {
	if err := u.r.Create(input.toRepoInput(uuid.New().String())); err != nil {
		return fmt.Errorf("error on create usecase: %v", err)
	}

	return nil
}
