package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"tracy_blog/blog/models"
)

type SearchController struct {
	beego.Controller
}

func (c *SearchController) Get() {
	po := pagination.NewPaginator(c.Ctx.Request, PER, len(models.Arts))
	arts := models.FilterArticles(PER, po.Page(), models.Arts)
	c.Data["articles"] = arts
	c.Data["per"] = PER
	c.Data["current"] = po.Page()
	c.Data["total"] = len(models.Arts)
	c.Data["paginator"] = po
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
	c.Data["IsLogin"] = VerifyCookie(c.Ctx)
	c.TplName = "results.html"
}

func (c *SearchController) Post() {
	cid := c.GetString("id")
	models.StorClickCate(cid)
	c.TplName = "results.html"
}
