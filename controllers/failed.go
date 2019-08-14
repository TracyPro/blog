package controllers

import "github.com/astaxie/beego"

type FailedController struct {
	beego.Controller
}

func (c *FailedController) Get() {
	c.TplName = "failed.html"
}
