package main

import (
	"log"

	"github.com/beego/beego/v2/server/web"
	"github.com/finn-inc/finn-server-tutorial/layered/controllers"
	"github.com/finn-inc/finn-server-tutorial/layered/registry"
	"github.com/joho/godotenv"
)

func main() {
	initConfig()

	reg := registry.NewRegistry()

	web.Router("/posts", controllers.NewPostsController(&reg))
	web.Run()
}

func initConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	web.BConfig.CopyRequestBody = true
}
