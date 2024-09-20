package presentation

import (
	"github.com/finn-inc/finn-server-tutorial/dip/models"
	"github.com/samber/lo"
)

type PostsPresentation struct{}

func NewPostsPresentation() PostsPresentation {
	return PostsPresentation{}
}

func (p PostsPresentation) Index(posts []models.Post) map[string]interface{} {
	return map[string]interface{}{
		"posts": lo.Map(posts, func(post models.Post, _ int) map[string]interface{} {
			return map[string]interface{}{
				"id":    post.Id,
				"title": post.Title,
				"body":  post.Body,
			}
		}),
	}
}

func (p PostsPresentation) Create() map[string]string {
	return map[string]string{
		"msg": "ok",
	}
}
