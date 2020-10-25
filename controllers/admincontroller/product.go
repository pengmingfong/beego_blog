package admincontroller

import (
	"gold/models"
)

type ProductController struct {
	BaseController
}

func (c *ProductController) List() {
	page, _ := c.GetInt("page")
	code, _ := c.GetInt("code")
	pagisize := 10
	result, total := models.ProductLists(page, pagisize)
	list := make([]map[string]interface{}, total)

	for k, v := range result {
		row := make(map[string]interface{})
		row["image"] = v.Image
		row["id"] = v.Id
		row["url"] = v.Url
		list[k] = row
	}

	c.Data["code"] = code
	c.Data["total"] = total
	c.Data["clist"] = list
	c.TplName = "admin/product/lists.html"
}

func (a *ProductController) Add() {
	a.TplName = "admin/product/add.html"
}

func (a *ProductController) Edit() {
	id, _ := a.GetInt("id")
	product, _ := models.ProductGetById(id)

	a.Data["data"] = product
	a.TplName = "admin/product/edit.html"
}

func (c *ProductController) EditOne() {
	id, _ := c.GetInt("id")
	url := c.GetString("url")

	bannber, _ := models.ProductGetById(id)
	bannber.Url = url
	image := bannber.Image
	// 获取上传文件
	f, h, err := c.GetFile("filename")
	if err == nil {
		defer f.Close()
		path := CreateDateDir()
		c.SaveToFile("filename", path+h.Filename)
		image = "/" + path + h.Filename
	}
	bannber.Image = image

	if err := bannber.Update(); err != nil {
		c.Ctx.Redirect(302, "/admin/product/list?page=1&code=1")
	}

	c.Ctx.Redirect(302, "/admin/product/list?page=1&code=2")
}
