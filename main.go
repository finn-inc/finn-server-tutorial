package main

import (
	"log"

	"github.com/beego/beego/v2/server/web"
	"github.com/finn-inc/finn-server-tutorial/controllers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "postgresql://postgres:postgres@localhost:15432/ft_dev"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Postgresに接続できませんでした: %v\n", err)
	}

	web.Router("/posts", &controllers.PostsController{
		DB: db,
	})
	web.Run()
}
