package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"gopkg.in/russross/blackfriday.v2"
	"tracy_blog/blog/models"
)

type IndexController struct {
	beego.Controller
}

const (
	PER    = 6 // 每页显示文章数
	CLICK  = 0 // 点击分类，显示该分类所有文章
	Search = 1 // 搜索条件检索
)

func (c *IndexController) Get() {
	// 数据库读取的文章信息
	// --------------summary && pagination-----------------
	total, arts := models.TotalArticles()
	po := pagination.NewPaginator(c.Ctx.Request, PER, total)
	c.Data["per"] = PER
	c.Data["current"] = po.Page()
	c.Data["total"] = total
	c.Data["paginator"] = po
	per_arts := models.FilterArticles(PER, po.Page(), arts)
	c.Data["articles"] = per_arts

	// 数据库读取的文章类别信息
	cates := models.TotalCategories()
	c.Data["cates"] = cates
	// --------------Mine message--------------------------
	mine := models.InsertMine()
	c.Data["mine"] = mine
	// ----------------------------------------------------

	// --------------Rank----------------------------------
	rank_arts := models.ViewsRank()
	c.Data["rank_arts"] = rank_arts
	// ----------------------------------------------------
	recommends := models.GetRecommendData()
	c.Data["recommends"] = recommends
	c.TplName = "index.html"
}

func (c *IndexController) Post() {
	title := c.GetString("title")
	category := c.GetString("category")
	author := c.GetString("author")
	meta := c.GetString("detail")
	source := c.GetString("source")
	link := c.GetString("link")
	html_content := string(blackfriday.Run([]byte(meta)))
	models.AddArticleContent(title, link, html_content, meta, author, category, source)
	c.TplName = "index.html"
}
