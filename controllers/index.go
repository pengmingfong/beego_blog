package controllers

import (
	"gold/models"

	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Index() {
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

	bannerresult, lens := models.BannerLists(1, 10)
	bannerlist := make([]map[string]interface{}, lens)
	for k, v := range bannerresult {
		row := make(map[string]interface{})
		row["image"] = v.Image
		bannerlist[k] = row
	}

	indexresult, _ := models.IndexList(1)
	prodlist := make([]map[string]interface{}, 0)
	goldlist := make([]map[string]interface{}, 0)
	nationallist := make([]map[string]interface{}, 0)
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
		if v.Cid == 4 {
			// 产品鉴赏文章
			if len(prodlist) >= 6 {
				continue
			}
			prodlist = append(prodlist, row)
		} else if v.Cid == 9 {
			// 黄金头条
			if len(goldlist) >= 10 {
				continue
			}
			goldlist = append(goldlist, row)
		} else if v.Cid == 10 {
			// 国际新闻
			if len(nationallist) >= 10 {
				continue
			}
			nationallist = append(nationallist, row)
		} else if v.Cid == 11 {
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

	this.Data["title"] = "首页"
	this.Data["list"] = list
	this.Data["bannerlist"] = bannerlist
	this.Data["productarticle"] = prodlist
	this.Data["goldlist"] = goldlist
	this.Data["nationallist"] = nationallist
	this.Data["newsspecial"] = newsspecial
	this.Data["prodcenter"] = prodcenter
	this.Data["interact"] = interact
	this.Data["distributor"] = distributor
	this.Data["cases"] = cases
	this.Data["designer"] = designer
	this.Data["aboutus"] = aboutus
	this.TplName = "home/index.html"
}
