package admincontroller

import (
	"fmt"
	"gold/models"
	"time"
)

type CategoryController struct {
	BaseController
}

func (c *CategoryController) Add() {
	result := models.CategoryListParent()
	list := make([]map[string]interface{}, len(result))

	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["name"] = v.Name
		list[k] = row
	}

	c.Data["category"] = list
	c.TplName = "admin/category/add.html"
}

func (c *CategoryController) AddOne() {
	name := c.GetString("name")
	pid, _ := c.GetInt("pid")
	level, _ := c.GetInt("level")
	index, _ := c.GetInt("index")

	category := new(models.Category)
	category.Name = name
	category.Pid = pid
	category.Level = level
	category.Index = index
	category.Created = time.Now().Format("2006-01-02 15:04:05")
	category.Updated = time.Now().Format("2006-01-02 15:04:05")

	id, err := models.AddCategory(category)
	if err != nil || id == 0 {
		c.Ctx.Redirect(302, "/admin/category/list?page=1&code=1")
	}
	c.Ctx.Redirect(302, "/admin/category/list?page=1&code=2")

}

func (c *CategoryController) List() {
	page, _ := c.GetInt("page")
	code, _ := c.GetInt("code")
	pagisize := 100
	result, total := models.CategoryLists(page, pagisize)
	list := make([]map[string]interface{}, total)

	for k, v := range result {
		row := make(map[string]interface{})
		row["name"] = v.Name
		row["id"] = v.Id
		row["pid"] = v.Pid
		row["insex"] = v.Index
		row["level"] = v.Level
		list[k] = row
	}

	c.Data["code"] = code
	c.Data["total"] = total
	c.Data["clist"] = list
	c.TplName = "admin/category/lists.html"
}

func (c *CategoryController) Edit() {
	id, _ := c.GetInt("id")
	category, _ := models.CategoryGetById(id)
	result := models.CategoryListParent()
	list := make([]map[string]interface{}, len(result))

	for k, v := range result {
		row := make(map[string]interface{})

		row["id"] = v.Id
		row["name"] = v.Name
		list[k] = row
	}

	c.Data["category"] = list
	c.Data["data"] = category
	c.TplName = "admin/category/edit.html"
}

func (c *CategoryController) EditOne() {
	id, _ := c.GetInt("id")
	name := c.GetString("name")
	pid, _ := c.GetInt("pid")
	level, _ := c.GetInt("level")
	index, _ := c.GetInt("index")

	category, _ := models.CategoryGetById(id)
	category.Index = index
	category.Name = name
	category.Level = level
	category.Pid = pid

	if err := category.Update(); err != nil {
		c.Ctx.Redirect(302, "/admin/category/list?page=1&code=1")
	}

	c.Ctx.Redirect(302, "/admin/category/list?page=1&code=2")
}

func (c *CategoryController) Delete() {
	id, _ := c.GetInt("id")
	result, _ := models.CategoryDelete(id)
	fmt.Println(result)
	c.Ctx.Redirect(302, "/admin/category/list?page=1&code=2")
}
