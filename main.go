package main

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/finn-inc/finn-server-tutorial/controllers"
	"github.com/finn-inc/finn-server-tutorial/registry"
)

func main() {
	web.BConfig.CopyRequestBody = true

	reg := registry.NewRegistry()

	web.Router("/posts", controllers.NewPostsController(&reg))
	web.Run()
}
