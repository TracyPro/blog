package controllers

import (
	"github.com/astaxie/beego"
)

type AboutController struct {
	beego.Controller
}

func (c *AboutController) Get() {
	c.Data["IsLogin"] = VerifyCookie(c.Ctx)
	c.TplName = "about.html"
}
