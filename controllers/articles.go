package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"tracy_blog/blog/models"
)

const (
	LIST_PER = 10 // 文章列表每页文章数
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {
	total, arts := models.TotalArticles()
	po := pagination.NewPaginator(c.Ctx.Request, LIST_PER, total)
	c.Data["per"] = LIST_PER
	c.Data["current"] = po.Page()
	c.Data["total"] = total
	c.Data["paginator"] = po
	per_arts := models.FilterArticles(LIST_PER, po.Page(), arts)
	c.Data["articles"] = per_arts
	c.TplName = "articles.html"
}

func (c *ArticleController) DelArticle() {
	del_id := c.GetString("id")
	status := models.DelArticle(del_id)
	c.Data["json"] = status
	c.TplName = "articles.html"
	c.ServeJSON()
}
