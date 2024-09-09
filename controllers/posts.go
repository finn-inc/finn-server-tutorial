package controllers

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/server/web"
	"github.com/finn-inc/finn-server-tutorial/presentation"
	"github.com/finn-inc/finn-server-tutorial/registry"
	"github.com/finn-inc/finn-server-tutorial/repository/implements/gorm"
	"github.com/finn-inc/finn-server-tutorial/usecase"
	"github.com/samber/lo"
)

type PostsController struct {
	web.Controller
	usecase *usecase.PostsUsecase
}

func NewPostsController(reg *registry.Registry) *PostsController {
	rg := *reg

	u := usecase.NewPostsUsecase(gorm.NewGormPostsRepository(&rg.DB))

	return &PostsController{
		usecase: u,
	}
}

func (c *PostsController) Get() {
	defer c.ServeJSON()

	posts, err := c.usecase.Index()
	if err != nil {
		c.Data["json"] = map[string]string{
			"error": fmt.Sprintf("Error on Get: %s\n", err),
		}
		return
	}

	p := presentation.PostsPresentation{}

	c.Data["json"] = p.Index(posts)
}

type post struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (p *post) validate() error {
	valid := validation.Validation{}
	valid.Required(p.Title, "titl")
	valid.Required(p.Body, "body")

	if valid.HasErrors() {
		errors := lo.Map(valid.Errors, func(err *validation.Error, _ int) string {
			return err.Message
		})
		return fmt.Errorf("validation failed: %s", strings.Join(errors, ", "))
	}

	return nil
}

func (p *post) toUsecaseInput() usecase.CreatePostInput {
	return usecase.CreatePostInput{
		Title: p.Title,
		Body:  p.Body,
	}
}

func (c *PostsController) Post() {
	defer c.ServeJSON()

	input := post{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &input); err != nil {
		c.Data["json"] = map[string]string{
			"error": fmt.Sprintf("Error on json.Unmarshal: %s\n", err),
		}
		return
	}

	if err := input.validate(); err != nil {
		c.Data["json"] = map[string]string{
			"error": fmt.Sprintf("Error on Validation: %s\n", err),
		}
		return
	}

	ui := input.toUsecaseInput()
	if err := c.usecase.Create(ui); err != nil {
		c.Data["json"] = map[string]string{
			"error": fmt.Sprintf("Error on Create: %s\n", err),
		}
		return
	}

	p := presentation.PostsPresentation{}

	c.Data["json"] = p.Create()
}
