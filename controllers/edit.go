package controllers

import (
	"github.com/astaxie/beego"
)

type EditController struct {
	beego.Controller
}

func (c *EditController) Get() {
	c.TplName = "markdown.html"
}

func (c *EditController) Post() {
	// 新建文章前要先登陆
	is_login := VerifyCookie(c.Ctx)
	c.Data["json"] = is_login
	c.ServeJSON()
	c.TplName = "markdown.html"
}
