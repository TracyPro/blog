package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"tracy_blog/blog/models"
)

var article models.Article

type DetailController struct {
	beego.Controller
}

func (c *DetailController) Get() {
	c.Data["article"] = models.DetailArt
	c.Data["IsLogin"] = VerifyCookie(c.Ctx)
	c.TplName = "detail.html"
}

func (c *DetailController) GetAjaxId() {
	aid, _ := strconv.Atoi(c.GetString("id"))
	models.ViewDetail(aid)
	c.TplName = "detail.html"
}
