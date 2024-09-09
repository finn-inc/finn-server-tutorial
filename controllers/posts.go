package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/finn-inc/finn-server-tutorial/registry"
	"github.com/finn-inc/finn-server-tutorial/repository/implements/gorm"
	"github.com/finn-inc/finn-server-tutorial/usecase"
	"github.com/finn-inc/finn-server-tutorial/views"
)

func NewPostsController(reg registry.Registry) *PostsController {
	repo := gorm.NewGormPostsRepository(reg.DB)

	return &PostsController{
		usecase: usecase.NewPostsUsecase(repo),
	}
}

type PostsController struct {
	web.Controller
	usecase usecase.PostsUsecase
}

func (c *PostsController) Get() {
	posts := c.usecase.Index()

	v := views.PostsView{}

	c.Data["json"] = v.Index(posts)
	c.ServeJSON()
}
