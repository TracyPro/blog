package controllers

import "github.com/astaxie/beego"

type ErrController struct {
	beego.Controller
}

func (c *ErrController) Get() {
	c.TplName = "404.html"
}
