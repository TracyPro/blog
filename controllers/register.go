package controllers

import (
	"github.com/astaxie/beego"
	"tracy_blog/blog/models"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	c.TplName = "register.html"
}

func (c *RegisterController) Post() {
	uname := c.GetString("name")
	pwd := c.GetString("pwd")
	status := models.Register(uname, pwd)
	c.Data["json"] = status
	c.ServeJSON()
	c.TplName = "register.html"
}
