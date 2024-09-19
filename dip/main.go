package main

import (
	"fmt"
	"log"

	"github.com/beego/beego/v2/server/web"
	"github.com/finn-inc/finn-server-tutorial/dip/config"
	"github.com/finn-inc/finn-server-tutorial/dip/controllers"
	"github.com/finn-inc/finn-server-tutorial/dip/registry"
	"github.com/finn-inc/finn-server-tutorial/dip/repository/implements"
	"github.com/finn-inc/finn-server-tutorial/dip/usecase"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	web.BConfig.CopyRequestBody = true

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	env, err := config.LoadEnv()
	if err != nil {
		panic(fmt.Errorf("環境変数を取得できませんでした: %w", err))
	}

	reg, err := registry.NewRegistryImpl(env)
	if err != nil {
		panic(fmt.Errorf("registryの初期化に失敗しました: %w", err))
	}

	uc := usecase.NewPostsUsecase(implements.NewPostsRepository(reg.DBConn()))
	web.Router("/posts", controllers.NewPostsController(reg, uc))
	web.Run()
}
