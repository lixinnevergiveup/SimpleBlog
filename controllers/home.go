package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.TplName = "home.html"
	c.Data["IsHome"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)
}
