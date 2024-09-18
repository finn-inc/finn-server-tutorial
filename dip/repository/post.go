package repository

import (
	"github.com/finn-inc/finn-server-tutorial/dip/models"
)

type PostsRepository interface {
	Index(page int, pageSize int) ([]models.Post, error)
	Save(post models.Post) error
}
