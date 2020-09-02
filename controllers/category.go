package controllers

import (
	"SimpleBlog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	op := c.Input().Get("op")
	switch op {
	case "add":
		name := c.Input().Get("name")
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			logs.Error(err)
		}
		c.Redirect("/category", 301)
		return

	case "del":
		name := c.Input().Get("id")
		if len(name) == 0 {
			break
		}

	}

	c.Data["IsCategory"] = true
	c.TplName = "category.html"
}
