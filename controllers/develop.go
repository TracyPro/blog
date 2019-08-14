package controllers

import "github.com/astaxie/beego"

type DevelopController struct {
	beego.Controller
}

func (c *DevelopController) Get() {
	c.TplName = "developing.html"
}
