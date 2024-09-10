package presentation

import (
	"github.com/finn-inc/finn-server-tutorial/hexagonal/usecase"
	"github.com/samber/lo"
)

type PostsPresentation struct{}

func NewPostsPresentation() PostsPresentation {
	return PostsPresentation{}
}

func (v *PostsPresentation) Index(posts []usecase.Post) map[string]interface{} {
	return map[string]interface{}{
		"posts": lo.Associate(posts, func(post usecase.Post) (string, string) {
			return post.Id, post.Title
		}),
	}
}

func (v *PostsPresentation) Create() map[string]string {
	return map[string]string{
		"msg": "ok",
	}
}
