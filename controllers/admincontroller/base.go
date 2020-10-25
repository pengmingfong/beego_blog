package admincontroller

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {
	userId := c.GetSession("userId")
	if userId == nil {
		c.Ctx.Redirect(302, "/login") //若Session中无用户ID则302重定向至登陆页面
	}
	c.Data["userId"] = userId
}
