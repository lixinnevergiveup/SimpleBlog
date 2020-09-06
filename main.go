package main

import (
	"SimpleBlog/controllers"
	"SimpleBlog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegisterDB()
}

func main() {
	// open orm debug mode
	orm.Debug = true
	// create default database if it not existed.
	orm.RunSyncdb("default", false, true)

	// register router
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{})

	beego.Run()
}
