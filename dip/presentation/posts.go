package presentation

import (
	"github.com/finn-inc/finn-server-tutorial/dip/models"
	"github.com/samber/lo"
)

type PostsPresentation struct{}

func NewPostsPresentation() PostsPresentation {
	return PostsPresentation{}
}

func (p *PostsPresentation) Index(posts []models.Post) map[string]interface{} {
	return map[string]interface{}{
		"posts": lo.Associate(posts, func(post models.Post) (string, string) {
			return post.Id, post.Title
		}),
	}
}

func (p *PostsPresentation) Create() map[string]string {
	return map[string]string{
		"msg": "ok",
	}
}
