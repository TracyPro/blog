package controllers

import (
	"github.com/astaxie/beego"
	"log"
	"tracy_blog/blog/models"
)

type VideoUploadController struct {
	beego.Controller
}

func (c *VideoUploadController) Get() {
	c.TplName = "uploadvideo.html"
}
func (c *VideoUploadController) Post() {
	f, h, err := c.GetFile("uploadname")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	models.CheckPath(models.UPLOAD_VIDEO_PATH)
	err = c.SaveToFile("uploadname", models.UPLOAD_VIDEO_PATH+h.Filename)
	if err != nil {
		log.Fatal("save err ", err)
	}

	models.StorVideos(models.UPLOAD_VIDEO_PATH + h.Filename)

	c.Data["json"] = `{"status":"success"}`
	c.TplName = "uploadvideo.html"
	c.ServeJSON()
	f.Close()
}
