package routers

import (
	"gold/controllers"
	"gold/controllers/admincontroller"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{}, "get:Index")
	beego.Router("/index", &controllers.IndexController{}, "get:Index")
	beego.Router("/index1", &controllers.IndexController{}, "get:Index1")
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
	beego.Router("/admin/article/delete", &admincontroller.ArticleController{}, "get:Delete")
	// 分类
	beego.Router("/admin/category/list", &admincontroller.CategoryController{}, "get:List")
	beego.Router("/admin/category/add", &admincontroller.CategoryController{}, "post:AddOne")
	beego.Router("/admin/category/add", &admincontroller.CategoryController{}, "get:Add")
	beego.Router("/admin/category/edit", &admincontroller.CategoryController{}, "get:Edit")
	beego.Router("/admin/category/edit", &admincontroller.CategoryController{}, "post:EditOne")
	beego.Router("/admin/category/delete", &admincontroller.CategoryController{}, "get:Delete")
	// banner
	beego.Router("/admin/banner/list", &admincontroller.BannerController{}, "get:List")
	beego.Router("/admin/banner/add", &admincontroller.BannerController{}, "post:AddOne")
	beego.Router("/admin/banner/add", &admincontroller.BannerController{}, "get:Add")
	beego.Router("/admin/banner/edit", &admincontroller.BannerController{}, "get:Edit")
	beego.Router("/admin/banner/edit", &admincontroller.BannerController{}, "post:EditOne")
	beego.Router("/admin/banner/delete", &admincontroller.BannerController{}, "get:Delete")
}
