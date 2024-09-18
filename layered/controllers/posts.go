package controllers

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/beego/beego/validation"
	"github.com/finn-inc/finn-server-tutorial/layered/registry"
	"github.com/finn-inc/finn-server-tutorial/layered/services"
	"github.com/finn-inc/finn-server-tutorial/layered/views"
	"github.com/samber/lo"
)

type PostsController struct {
	BaseController
}

func NewPostsController(reg *registry.Registry) *PostsController {
	return &PostsController{
		BaseController: BaseController{
			reg: reg,
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

	posts, err := services.NewPostsService(c.reg.DB().Client()).Index(page)
	if err != nil {
		c.Data["json"] = map[string]string{
			"error": fmt.Sprintf("Error on Get: %s\n", err),
		}
		return
	}

	v := views.NewPostsView()

	c.Data["json"] = v.Index(posts)
}

type post struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (p *post) toServiceInput() services.CreatePostInput {
	return services.CreatePostInput{
		Title: p.Title,
		Body:  p.Body,
	}
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

	if err := services.NewPostsService(c.reg.DB().Client()).Create(input.toServiceInput()); err != nil {
		c.Data["json"] = map[string]string{
			"error": fmt.Sprintf("Error on Create: %s\n", err),
		}
		return
	}

	p := views.PostsView{}

	c.Data["json"] = p.Create()
}
