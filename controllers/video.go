package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"tracy_blog/blog/models"
)

const (
	PER_VIDEOS = 8
)

type VideoController struct {
	beego.Controller
}

func (c *VideoController) Get() {
	videos, total := models.GetVideos()
	po := pagination.NewPaginator(c.Ctx.Request, PER_VIDEOS, total)
	c.Data["per"] = PER_VIDEOS
	c.Data["current"] = po.Page()
	c.Data["total"] = total
	c.Data["paginator"] = po
	per_videos := models.FilterVideos(PER_VIDEOS, po.Page(), videos)
	c.Data["videos"] = per_videos
	c.TplName = "listvideo.html"
}
