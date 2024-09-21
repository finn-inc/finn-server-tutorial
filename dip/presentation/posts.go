package presentation

import (
	"github.com/finn-inc/finn-server-tutorial/dip/models"
	"github.com/samber/lo"
)

func NewIndexPostsResponse(posts []models.Post) map[string]interface{} {
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

func NewCreatePostResponse() map[string]string {
	return map[string]string{
		"msg": "ok",
	}
}
