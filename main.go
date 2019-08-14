package main

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"tracy_blog/blog/models"
	_ "tracy_blog/blog/routers"
)

func init() {
	models.RegisterDB()
}

func main() {
	beego.AddFuncMap("CatesTitle", models.GetCategory)
	beego.Run()
}
