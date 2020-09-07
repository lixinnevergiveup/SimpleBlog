package controllers

import (
	"SimpleBlog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.TplName = "home.html"
	c.Data["IsHome"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	topics, err := models.GetAllTopics(true)
	if err != nil {
		logs.Error(err)
	} else {
		c.Data["Topics"] = topics
	}
}
