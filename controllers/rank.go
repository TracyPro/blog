package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"tracy_blog/blog/models"
)

var rank_arts models.Article

type RankController struct {
	beego.Controller
}

func (c *RankController) Get() {
	c.Data["rank_arts"] = models.DetailArt
	c.TplName = "rank.html"
}

func (c *RankController) GetAjaxRid() {
	rid, _ := strconv.Atoi(c.GetString("id"))
	models.ViewDetail(rid)
	c.TplName = "rank.html"
}
