package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"tracy_blog/blog/models"
)

const (
	PER_IMGS = 12
)

type PictureController struct {
	beego.Controller
}

func (c *PictureController) Get() {
	jpgs, total := models.GetImgs()
	po := pagination.NewPaginator(c.Ctx.Request, PER_IMGS, total)
	c.Data["per"] = PER_IMGS
	c.Data["current"] = po.Page()
	c.Data["total"] = total
	c.Data["paginator"] = po
	per_imgs := models.FilterImgs(PER_IMGS, po.Page(), jpgs)
	c.Data["imgs"] = per_imgs

	c.TplName = "listpic.html"
}
