package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"tracy_blog/blog/models"
)

type TechnologyController struct {
	beego.Controller
}

func (c *TechnologyController) Get() {
	// 数据库读取的文章信息
	// --------------summary && pagination-----------------
	total, arts := models.TotalArticles()
	tech_arts, tech_total := models.GetTechnologyArts(arts)
	po := pagination.NewPaginator(c.Ctx.Request, PER, tech_total)
	c.Data["per"] = PER
	c.Data["current"] = po.Page()
	c.Data["total"] = total
	c.Data["paginator"] = po
	per_arts := models.FilterArticles(PER, po.Page(), tech_arts)
	// 过滤掉非技术文章
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
