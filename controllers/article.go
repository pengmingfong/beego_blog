package controllers

import (
	"gold/models"

	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

func (a *ArticleController) Detail() {
	a.getmenu()
	a.getfooter()
	id, _ := a.GetInt("id")
	article, _ := models.ArticleById(id)

	category, _ := models.CategoryGetById(article.Cid)

	a.Data["title"] = article.Title
	a.Data["content"] = article
	a.Data["category"] = category
	a.TplName = "home/detail.html"
}

func (a *ArticleController) List() {
	a.getmenu()
	a.getfooter()
	cid, _ := a.GetInt("cid")
	page, _ := a.GetInt("page")
	pagesize := 10

	// 初始化查询条件
	condition := make(map[string]int)
	condition["cid"] = cid
	condition["pagesize"] = cid
	result, _ := models.ArticleByCids(page, pagesize, cid, condition)
	// 查询同等级分类
	// var ccid int
	category, _ := models.CategoryGetById(cid)
	if category.Pid == 0 {
		cid = category.Id
	} else {
		cid = category.Pid
	}

	categorylist := models.CategoryBypid(cid)
	categorymap := make([]map[string]interface{}, len(categorylist))

	for k, v := range categorylist {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["name"] = v.Name
		categorymap[k] = row
	}

	list := make([]map[string]interface{}, len(result))

	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["title"] = v.Title
		row["created"] = v.Created
		list[k] = row
	}

	a.Data["title"] = category.Name
	a.Data["cid"] = cid
	a.Data["alist"] = list
	a.Data["category"] = category
	a.Data["categorymap"] = categorymap
	a.TplName = "home/list.html"
}

func (a *ArticleController) getmenu() {
	bannerresult, lens := models.BannerLists(1, 10)
	bannerlist := make([]map[string]interface{}, lens)
	for k, v := range bannerresult {
		row := make(map[string]interface{})
		row["image"] = v.Image
		bannerlist[k] = row
	}

	result := models.CategoryList()
	list := make([]map[string]interface{}, 0)
	for _, v := range result {
		if v.Pid != 0 || v.Index != 1 {
			continue
		}
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["name"] = v.Name
		row["pid"] = v.Pid
		row["status"] = v.Status
		row["level"] = v.Level
		row["create_time"] = v.Created
		row["update_time"] = v.Updated
		list = append(list, row)
	}

	a.Data["bannerlist"] = bannerlist
	a.Data["list"] = list
}

func (a *ArticleController) getfooter() {
	indexresult, _ := models.IndexList(1)
	newsspecial := make([]map[string]interface{}, 0)
	prodcenter := make([]map[string]interface{}, 0)
	interact := make([]map[string]interface{}, 0)
	distributor := make([]map[string]interface{}, 0)
	cases := make([]map[string]interface{}, 0)
	designer := make([]map[string]interface{}, 0)
	aboutus := make([]map[string]interface{}, 0)

	for _, v := range indexresult {
		row := make(map[string]interface{})

		row["id"] = v.Id
		row["cid"] = v.Cid
		row["image"] = v.Image
		row["title"] = v.Title
		row["sub_title"] = v.Sub_title
		if v.Cid == 11 {
			// 新闻专区
			if len(newsspecial) >= 6 {
				continue
			}
			newsspecial = append(newsspecial, row)
		} else if v.Cid == 12 {
			// 产品中心
			if len(prodcenter) >= 6 {
				continue
			}
			prodcenter = append(prodcenter, row)
		} else if v.Cid == 13 {
			// 互动体验
			if len(interact) >= 6 {
				continue
			}
			interact = append(interact, row)
		} else if v.Cid == 14 {
			// 经销商专区
			if len(distributor) >= 6 {
				continue
			}
			distributor = append(distributor, row)
		} else if v.Cid == 15 {
			if len(cases) >= 6 {
				continue
			}
			cases = append(cases, row)
		} else if v.Cid == 16 {
			if len(designer) >= 6 {
				continue
			}
			designer = append(designer, row)
		} else if v.Cid == 17 {
			if len(aboutus) >= 6 {
				continue
			}
			aboutus = append(aboutus, row)
		}

		continue
	}

	a.Data["newsspecial"] = newsspecial
	a.Data["prodcenter"] = prodcenter
	a.Data["interact"] = interact
	a.Data["distributor"] = distributor
	a.Data["cases"] = cases
	a.Data["designer"] = designer
	a.Data["aboutus"] = aboutus
}
