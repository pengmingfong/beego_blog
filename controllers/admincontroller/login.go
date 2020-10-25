package admincontroller

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (a *LoginController) Login() {

	a.TplName = "admin/article/login.html"
}

func (a *LoginController) LoginUser() {
	username := a.GetString("username")
	password := a.GetString("password")

	if username == "admin" && password == "wodecaifuadmin123" {
		a.SetSession("userId", 1)
		a.Ctx.Redirect(302, "/admin/article/list") //若Session中无用户ID则302重定向至登陆页面
	}
	a.Data["username"] = username
	a.Data["password"] = password
	a.Ctx.Redirect(302, "/login") //若Session中无用户ID则302重定向至登陆页面
}
