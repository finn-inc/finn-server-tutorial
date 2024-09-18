package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/beego/beego/v2/server/web"
	"github.com/finn-inc/finn-server-tutorial/dip/controllers"
	"github.com/finn-inc/finn-server-tutorial/dip/registry"
	"github.com/finn-inc/finn-server-tutorial/dip/repository/implements"
	"github.com/finn-inc/finn-server-tutorial/dip/usecase"
	"github.com/finn-inc/finn-server-tutorial/dip/utils"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	web.BConfig.CopyRequestBody = true

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	env, err := utils.LoadEnv()
	if err != nil {
		panic(fmt.Errorf("環境変数を取得できませんでした: %w", err))
	}

	dbConn, err := sql.Open("postgres", env.DatabaseURL)
	if err != nil {
		panic(fmt.Errorf("postgresに接続できませんでした: %w", err))
	}

	reg := registry.NewRegistryImpl(registry.RegistryConfig{
		DBConn: dbConn,
	})

	uc := usecase.NewPostsUsecase(implements.NewPostsRepository(reg.DBConn()))
	web.Router("/posts", controllers.NewPostsController(reg, uc))
	web.Run()
}
