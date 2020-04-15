package routers

import (
	"gold/controllers"
	"gold/controllers/admincontroller"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{}, "get:Index")
	beego.Router("/index", &controllers.IndexController{}, "get:Index")
	beego.Router("/detail", &controllers.ArticleController{}, "get:Detail")
	beego.Router("/list", &controllers.ArticleController{}, "get:List")
	back()
}

func back() {
	beego.Router("/admin/article/add", &admincontroller.ArticleController{}, "get:Add")
	beego.Router("/admin/article/add", &admincontroller.ArticleController{}, "post:AddOne")
	beego.Router("/upload", &admincontroller.ArticleController{}, "post:Upload")
	beego.Router("/admin/article/list", &admincontroller.ArticleController{}, "get:List")
	beego.Router("/admin/article/edit", &admincontroller.ArticleController{}, "get:Edit")
	beego.Router("/admin/article/edit", &admincontroller.ArticleController{}, "post:EditOne")
}
