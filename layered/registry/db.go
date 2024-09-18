package registry

import (
	"fmt"

	"github.com/finn-inc/finn-server-tutorial/layered/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	client *gorm.DB
}

func NewDB() DB {
	cfg, err := utils.LoadEnv()
	if err != nil {
		panic(fmt.Errorf("環境変数を取得できませんでした: %w", err))
	}
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("postgresに接続できませんでした: %w", err))
	}

	return DB{
		client: db,
	}
}

func (db DB) Client() *gorm.DB {
	return db.client
}
