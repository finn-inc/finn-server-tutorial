package implements

import (
	"database/sql"
	"fmt"

	"github.com/finn-inc/finn-server-tutorial/dip/models"
	"github.com/finn-inc/finn-server-tutorial/dip/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostsRepositoryImpl struct {
	db *gorm.DB
}

func NewPostsRepository(dbConn *sql.DB) (repository.PostsRepository, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: dbConn,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &PostsRepositoryImpl{
		db: db,
	}, nil
}

func (r *PostsRepositoryImpl) Index(page int, pageSize int) ([]models.Post, error) {
	var posts []models.Post
	if res := r.db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&posts); res.Error != nil {
		return nil, fmt.Errorf("error on getting posts: %w", res.Error)
	}

	return posts, nil
}

func (r *PostsRepositoryImpl) Save(post models.Post) error {
	if res := r.db.Save(&post); res.Error != nil {
		return fmt.Errorf("error on creating post: %w", res.Error)
	}

	return nil
}
