package controllers

import (
	"SimpleBlog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get() {
	c.Data["IsTopic"] = true
	c.TplName = "topic.html"
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	topics, err := models.GetAllTopics()
	if err != nil {
		logs.Error(err)
	} else {
		c.Data["Topics"] = topics
	}
}

func (c *TopicController) Add() {
	c.TplName = "topic_add.html"
	c.Data["IsLogin"] = checkAccount(c.Ctx)
}

func (c *TopicController) Post() {
	res := checkAccount(c.Ctx)
	logs.Info(res)
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}

	title := c.Input().Get("title")
	content := c.Input().Get("content")

	err := models.AddTopic(title, content)
	if err != nil {
		logs.Error(err)
	}
	c.Redirect("/topic", 302)
}
