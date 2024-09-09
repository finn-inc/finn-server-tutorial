package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/finn-inc/finn-server-tutorial/services"
	"github.com/finn-inc/finn-server-tutorial/views"
	"gorm.io/gorm"
)

type PostsController struct {
	web.Controller
	DB *gorm.DB
}

func (c *PostsController) Get() {
	s := services.PostsService{
		DB: c.DB,
	}
	posts := s.Index(1, 10)

	v := views.PostsView{}

	c.Data["json"] = v.Index(posts)
	c.ServeJSON()
}
