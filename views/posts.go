package views

import (
	"github.com/finn-inc/finn-server-tutorial/usecase"
	"github.com/samber/lo"
)

type PostsView struct{}

func (v *PostsView) Index(posts []usecase.Post) map[string]interface{} {
	return map[string]interface{}{
		"posts": lo.Associate(posts, func(post usecase.Post) (string, string) {
			return post.Id, post.Title
		}),
	}
}
