package usecase

import (
	"github.com/finn-inc/finn-server-tutorial/repository"
	"github.com/samber/lo"
)

type PostsUsecase struct {
	r repository.PostRepository
}

type Post struct {
	Id    string
	Title string
	Body  string
}

func NewPostsUsecase(r repository.PostRepository) PostsUsecase {
	return PostsUsecase{
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

func (u PostsUsecase) Index() []Post {
	posts := u.r.Index(1, 10)
	return lo.Map(posts, func(post repository.Post, _ int) Post {
		return u.repoToOutput(post)
	})
}
