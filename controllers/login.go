package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"tracy_blog/blog/models"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	is_exit := c.Input().Get("exit") == "true"
	if is_exit {
		c.Ctx.SetCookie("uname", "", -1, "/")
		c.Ctx.SetCookie("pwd", "", -1, "/")
		c.Redirect("/", 302)
		return
	}
	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	uname := c.GetString("name")
	pwd := c.GetString("pwd")
	status := models.Login(uname, pwd)
	if status {
		maxAge := 0
		c.Ctx.SetCookie("uname", uname, maxAge, "/")
		c.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	}
	c.Data["json"] = status
	c.ServeJSON()
	c.TplName = "login.html"
}

func VerifyCookie(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}
	uname := ck.Value

	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := ck.Value
	return models.Login(uname, pwd)
}
