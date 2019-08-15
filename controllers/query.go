package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"tracy_blog/blog/models"
)

type QueryController struct {
	beego.Controller
}

func (c *QueryController) Get() {
	po := pagination.NewPaginator(c.Ctx.Request, PER, len(models.Final))
	arts := models.FilterArticles(PER, po.Page(), models.Final)
	if len(arts) == 0 {
		c.Redirect("/search/failed", 302)
		return
	}
	c.Data["articles"] = arts
	c.Data["per"] = PER
	c.Data["current"] = po.Page()
	c.Data["total"] = len(models.Final)
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
	c.TplName = "query.html"
}

func (c *QueryController) Post() {
	if len(models.Final) != 0 {
		models.Final = nil
	}
	search_content := c.GetString("keyboard")
	models.Search(search_content)
	c.TplName = "query.html"
	c.Redirect("/article/query", 302)
	return
}
