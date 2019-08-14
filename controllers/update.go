package controllers

import (
	"github.com/astaxie/beego"
	"gopkg.in/russross/blackfriday.v2"
	"tracy_blog/blog/models"
)

var EditId string

type UpdateController struct {
	beego.Controller
}

func (c *UpdateController) Get() {
	c.Data["article"] = models.EditArt
	c.TplName = "update.html"
}

func (c *UpdateController) Post() {
	EditId = c.GetString("id")
	models.GetEditArticle(EditId)
	c.TplName = "update.html"
}

func (c *UpdateController) PostUpdate() {
	// 更新文章信息
	title := c.GetString("title")
	category := c.GetString("category")
	author := c.GetString("author")
	meta := c.GetString("detail")
	source := c.GetString("source")
	link := c.GetString("link")
	html_content := string(blackfriday.Run([]byte(meta)))
	models.UpdateArticle(EditId, title, link, html_content, meta, author, category, source)
	c.TplName = "update.html"
}
