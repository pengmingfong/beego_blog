package admincontroller

import (
	"fmt"
	"gold/models"

	"github.com/astaxie/beego"
)

type BannerController struct {
	beego.Controller
}

func (c *BannerController) List() {
	page, _ := c.GetInt("page")
	code, _ := c.GetInt("code")
	pagisize := 10
	result, total := models.BannerLists(page, pagisize)
	list := make([]map[string]interface{}, total)

	for k, v := range result {
		row := make(map[string]interface{})
		row["image"] = v.Image
		row["id"] = v.Id
		row["level"] = v.Level
		list[k] = row
	}

	c.Data["code"] = code
	c.Data["total"] = total
	c.Data["clist"] = list
	c.TplName = "admin/banner/lists.html"
}

func (c *BannerController) Add() {
	c.TplName = "admin/banner/add.html"
}

func (c *BannerController) AddOne() {
	image := ""
	level, _ := c.GetInt("level")

	// 获取上传文件
	f, h, err := c.GetFile("filename")
	if err == nil {
		defer f.Close()
		path := CreateDateDir()
		c.SaveToFile("filename", path+h.Filename)
		image = path + h.Filename
	}

	banner := new(models.Banner)
	banner.Image = image
	banner.Level = level

	id, err := models.AddBanner(banner)
	if err != nil || id == 0 {
		c.Ctx.Redirect(302, "/admin/banner/list?page=1&code=1")
	}
	c.Ctx.Redirect(302, "/admin/banner/list?page=1&code=2")

}

func (c *BannerController) Edit() {
	id, _ := c.GetInt("id")
	category, _ := models.BannerGetById(id)

	c.Data["data"] = category
	c.TplName = "admin/banner/edit.html"
}

func (c *BannerController) EditOne() {
	id, _ := c.GetInt("id")
	level, _ := c.GetInt("level")

	bannber, _ := models.BannerGetById(id)
	bannber.Level = level
	image := bannber.Image
	// 获取上传文件
	f, h, err := c.GetFile("filename")
	if err == nil {
		defer f.Close()
		path := CreateDateDir()
		c.SaveToFile("filename", path+h.Filename)
		image = path + h.Filename
	}
	bannber.Image = image

	if err := bannber.Update(); err != nil {
		c.Ctx.Redirect(302, "/admin/banner/list?page=1&code=1")
	}

	c.Ctx.Redirect(302, "/admin/banner/list?page=1&code=2")
}

func (c *BannerController) Delete() {
	id, _ := c.GetInt("id")
	result, _ := models.BannerDelete(id)
	fmt.Println(result)
	c.Ctx.Redirect(302, "/admin/banner/list?page=1&code=2")
}
