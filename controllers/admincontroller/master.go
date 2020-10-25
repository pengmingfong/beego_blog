package admincontroller

import (
	"gold/models"
)

type MasterController struct {
	BaseController
}

func (c *MasterController) List() {
	page, _ := c.GetInt("page")
	code, _ := c.GetInt("code")
	pagisize := 10
	result, total := models.MasterLists(page, pagisize)
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
	c.TplName = "admin/master/lists.html"
}

func (a *MasterController) Add() {
	a.TplName = "admin/master/add.html"
}

func (a *MasterController) Edit() {
	id, _ := a.GetInt("id")
	product, _ := models.MasterGetById(id)

	a.Data["data"] = product
	a.TplName = "admin/master/edit.html"
}

func (c *MasterController) EditOne() {
	id, _ := c.GetInt("id")
	url := c.GetString("url")

	bannber, _ := models.MasterGetById(id)
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
		c.Ctx.Redirect(302, "/admin/master/list?page=1&code=1")
	}

	c.Ctx.Redirect(302, "/admin/master/list?page=1&code=2")
}
