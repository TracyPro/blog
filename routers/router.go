package routers

import (
	"github.com/astaxie/beego"
	"tracy_blog/blog/controllers"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/about", &controllers.AboutController{})
	beego.Router("/listpic", &controllers.PictureController{})
	beego.Router("/listvideo", &controllers.VideoController{})
	beego.Router("/article/new", &controllers.EditController{})
	beego.Router("/article/content", &controllers.IndexController{})
	beego.Router("/article/detail", &controllers.DetailController{})
	beego.Router("/article/summary", &controllers.IndexController{})
	beego.Router("/article/aid", &controllers.DetailController{}, "POST:GetAjaxId")
	beego.Router("/article/rank/detail", &controllers.RankController{})
	beego.Router("/article/rid", &controllers.RankController{}, "POST:GetAjaxRid")
	beego.Router("/article/search/cid", &controllers.SearchController{})
	beego.Router("/article/search/category", &controllers.SearchController{})
	beego.Router("/error", &controllers.ErrController{})
	beego.Router("/search/failed", &controllers.FailedController{})
	beego.Router("/article/query", &controllers.QueryController{})
	beego.Router("/upload/images", &controllers.ImgUploadController{})
	beego.Router("/upload/videos", &controllers.VideoUploadController{})
	beego.Router("/developing", &controllers.DevelopController{})
	beego.Router("/articles/technology", &controllers.TechnologyController{})
	beego.Router("/article/list", &controllers.ArticleController{})
	beego.Router("/article/delete/did", &controllers.ArticleController{}, "POST:DelArticle")
	beego.Router("/article/update", &controllers.UpdateController{})
	beego.Router("/article/update/submit", &controllers.UpdateController{}, "POST:PostUpdate")
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})
}
