package views

import (
	"github.com/finn-inc/finn-server-tutorial/layered/services"
	"github.com/samber/lo"
)

type PostsView struct{}

func NewPostsView() PostsView {
	return PostsView{}
}

func (v *PostsView) Index(posts []services.Post) map[string]interface{} {
	return map[string]interface{}{
		"posts": lo.Associate(posts, func(post services.Post) (string, string) {
			return post.Id, post.Title
		}),
	}
}

func (v *PostsView) Create() map[string]string {
	return map[string]string{
		"msg": "ok",
	}
}
