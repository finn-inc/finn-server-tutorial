package main

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	"github.com/finn-inc/finn-server-tutorial/dip/config"
	"github.com/finn-inc/finn-server-tutorial/dip/controllers"
	"github.com/finn-inc/finn-server-tutorial/dip/registry"
	"github.com/finn-inc/finn-server-tutorial/dip/repository/implements"
	"github.com/finn-inc/finn-server-tutorial/dip/usecase"
	_ "github.com/lib/pq"
)

func main() {
	web.BConfig.CopyRequestBody = true

	web.InsertFilter("*", web.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	env, err := config.LoadEnv()
	if err != nil {
		panic(fmt.Errorf("環境変数を取得できませんでした: %w", err))
	}

	reg, err := registry.NewRegistryImpl(env)
	if err != nil {
		panic(fmt.Errorf("registryの初期化に失敗しました: %w", err))
	}

	postsRepository, err := implements.NewPostsRepository(reg.DBConn())
	if err != nil {
		panic(fmt.Errorf("データベースクライアントの初期化に失敗しました: %w", err))
	}

	uc := usecase.NewPostsUsecase(postsRepository)
	web.Router("/posts", controllers.NewPostsController(reg, uc))
	web.Run()
}
