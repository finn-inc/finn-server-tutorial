package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/finn-inc/finn-server-tutorial/layered/registry"
)

type BaseController struct {
	web.Controller
	reg *registry.Registry
}
