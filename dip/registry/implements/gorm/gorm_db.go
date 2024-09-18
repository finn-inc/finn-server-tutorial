package gorm

import (
	"fmt"

	"github.com/finn-inc/finn-server-tutorial/dip/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormDB struct {
	client *gorm.DB
}

func (d *GormDB) GetClient() *gorm.DB {
	return d.client
}

func NewGormDB(cfg *utils.Config) GormDB {
	// postgresライブラリのpostgres接続インスタンスを渡す
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Postgresに接続できませんでした: %v\n", err))
	}

	return GormDB{
		client: db,
	}
}
