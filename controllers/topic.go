package controllers

import (
	"SimpleBlog/models"
	"fmt"
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
	topics, err := models.GetAllTopics(false)
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

func (c *TopicController) View() {
	c.TplName = "topic_view.html"
	topic, err := models.GetTopic(c.Ctx.Input.Param("0"))
	if err != nil {
		logs.Error(err)
		c.Redirect("/", 302)
		return
	}

	c.Data["Topic"] = topic
	c.Data["Tid"] = c.Ctx.Input.Param("0")
}

func (c *TopicController) Modify() {
	c.TplName = "topic_modify.html"

	tid := c.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		logs.Error(err)
		c.Redirect("/", 302)
		return
	}

	c.Data["Topic"] = topic
	c.Data["Tid"] = tid
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
	tid := c.Input().Get("tid")
	logs.Info(fmt.Sprintf("%v==================", tid))

	var err error
	if len(tid) == 0 {
		err = models.AddTopic(title, content)
	} else {
		err = models.ModifyTopic(tid, title, content)
	}

	if err != nil {
		logs.Error(err)
	}
	c.Redirect("/topic", 302)
}
