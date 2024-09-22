package usecase

import (
	"fmt"

	"github.com/finn-inc/finn-server-tutorial/dip/models"
	"github.com/finn-inc/finn-server-tutorial/dip/repository"
)

type PostsUsecase struct {
	postRepository repository.PostsRepository
}

func NewPostsUsecase(r repository.PostsRepository) *PostsUsecase {
	return &PostsUsecase{
		postRepository: r,
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

func (u *PostsUsecase) IndexPosts(page int) ([]models.Post, error) {
	posts, err := u.postRepository.Index(page, 10)
	if err != nil {
		return nil, fmt.Errorf("error on indexing posts: %w", err)
	}

	return posts, nil
}

func (u *PostsUsecase) CreatePost(input CreatePostInput) error {
	post := models.NewPost(models.NewPostInput{
		Title: input.Title,
		Body:  input.Body,
	})
	if err := u.postRepository.Save(post); err != nil {
		return fmt.Errorf("error on create usecase: %w", err)
	}

	return nil
}
