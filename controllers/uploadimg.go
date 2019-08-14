package controllers

import (
	"github.com/astaxie/beego"
	"log"
	"tracy_blog/blog/models"
)

type ImgUploadController struct {
	beego.Controller
}

func (c *ImgUploadController) Get() {
	c.TplName = "uploadimg.html"
}

func (c *ImgUploadController) Post() {
	f, h, err := c.GetFile("uploadname")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	models.CheckPath(models.UPLOAD_IMG_PATH)
	err = c.SaveToFile("uploadname", models.UPLOAD_IMG_PATH+h.Filename)
	if err != nil {
		log.Fatal("save err ", err)
	}

	models.StorImgs(models.UPLOAD_IMG_PATH + h.Filename)

	c.Data["json"] = `{"status":"success"}`
	c.TplName = "uploadimg.html"
	c.ServeJSON()
	f.Close()
}
