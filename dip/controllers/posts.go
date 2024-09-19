package controllers

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/server/web"
	"github.com/finn-inc/finn-server-tutorial/dip/presentation"
	"github.com/finn-inc/finn-server-tutorial/dip/registry"
	"github.com/finn-inc/finn-server-tutorial/dip/usecase"
	"github.com/samber/lo"
)

type PostsControllerBase struct {
	web.Controller
	reg     registry.Registry
	usecase *usecase.PostsUsecase
}

type PostsController struct {
	PostsControllerBase
}

func NewPostsController(reg registry.Registry, usecase *usecase.PostsUsecase) *PostsController {
	return &PostsController{
		PostsControllerBase: PostsControllerBase{
			reg:     reg,
			usecase: usecase,
		},
	}
}

func (c *PostsController) Get() {
	defer c.ServeJSON()

	var page int
	if err := c.Ctx.Input.Bind(&page, "page"); err != nil {
		c.Data["json"] = map[string]string{
			"error": fmt.Sprintf("Error on Binding: %s\n", err),
		}
		return
	}
	if page == 0 {
		page = 1
	}

	posts, err := c.usecase.IndexPosts(page)
	if err != nil {
		c.Data["json"] = map[string]string{
			"error": fmt.Sprintf("Error on Get: %s\n", err),
		}
		return
	}

	p := presentation.NewPostsPresentation()

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

	if err := c.usecase.CreatePost(input.toUsecaseInput()); err != nil {
		c.Data["json"] = map[string]string{
			"error": fmt.Sprintf("Error on Create: %s\n", err),
		}
		return
	}

	p := presentation.NewPostsPresentation()

	c.Data["json"] = p.Create()
}
