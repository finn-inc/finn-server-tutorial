package registry

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Client *gorm.DB
}

func NewDB() DB {
	dsn := "postgresql://postgres:postgres@localhost:5432/ft_dev"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Postgresに接続できませんでした: %v\n", err))
	}

	return DB{
		Client: db,
	}
}
