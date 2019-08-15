package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
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
	// 判断当前登陆用户，文章列表只显示当前用户登陆时创建的文章
	// 如果当前登陆用户是管理员，显示所有文章
	var arts []*models.Article
	var total int
	if VerifyCookie(c.Ctx) {
		ck, _ := c.Ctx.Request.Cookie("uname")
		// 筛选出文章User字段为uname的文章
		if IsAdmin(c.Ctx) {
			total, arts = models.TotalArticles() // 如果是管理员
		} else {
			arts, total = models.FilterUserArts(ck.Value)
		}
	} else {
		arts = nil
		total = 0
	}
	po := pagination.NewPaginator(c.Ctx.Request, LIST_PER, total)
	c.Data["per"] = LIST_PER
	c.Data["current"] = po.Page()
	c.Data["total"] = total
	c.Data["paginator"] = po
	per_arts := models.FilterArticles(LIST_PER, po.Page(), arts)
	c.Data["articles"] = per_arts
	c.Data["IsLogin"] = VerifyCookie(c.Ctx)
	c.TplName = "articles.html"
}

func (c *ArticleController) DelArticle() {
	del_id := c.GetString("id")
	status := models.DelArticle(del_id)
	c.Data["json"] = status
	c.TplName = "articles.html"
	c.ServeJSON()
}

func IsAdmin(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}
	uname := ck.Value

	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := ck.Value
	return uname == "admin" && pwd == "antiy?1024"
}
