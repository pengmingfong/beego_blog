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

	categorys := models.GetChild(0)

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
	dongtai := make([]map[string]interface{}, 0)

	for _, v := range indexresult {
		row := make(map[string]interface{})

		row["id"] = v.Id
		row["cid"] = v.Cid
		row["image"] = v.Image
		row["title"] = v.Title
		row["sub_title"] = v.Sub_title
		row["desc"] = v.Desc
		row["indexs"] = v.Index
		if v.Index == 1 && v.Image != "" {
			// 产品鉴赏文章
			if len(prodlist) >= 6 {
				continue
			}

			prodlist = append(prodlist, row)
			// dongtai = append(dongtai, row)
		} else if v.Cid == 52 && v.Status == 1 {
			if len(goldlist) >= 3 {
				continue
			}
			goldlist = append(goldlist, row)
		}
		continue
	}

	this.Data["categorys"] = categorys
	this.Data["title"] = "安徽黄金-打造黄金百年企业.谱写徽韵文化传奇"
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
	this.Data["dongtai"] = dongtai
	this.Data["view"] = 1
	this.TplName = "home/index.html"
}

func (this *IndexController) Index1() {
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

	categorys := models.GetChild(0)

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
	dongtai := make([]map[string]interface{}, 0)
	xinwen := make([]map[string]interface{}, 0)

	for _, v := range indexresult {
		row := make(map[string]interface{})

		row["id"] = v.Id
		row["cid"] = v.Cid
		row["image"] = v.Image
		row["title"] = v.Title
		row["sub_title"] = v.Sub_title
		row["desc"] = v.Desc
		row["indexs"] = v.Index
		if v.Index == 1 && v.Image != "" {
			// 产品鉴赏文章
			if len(prodlist) >= 6 {
				continue
			}

			prodlist = append(prodlist, row)
			// dongtai = append(dongtai, row)
		} else if v.Cid == 52 && v.Status == 1 {
			if len(goldlist) >= 3 {
				continue
			}
			goldlist = append(goldlist, row)
		} else if v.Index == 3 && v.Status == 1 {
			if len(xinwen) >= 9 {
				continue
			}
			xinwen = append(xinwen, row)
		}
		continue
	}

	this.Data["categorys"] = categorys
	this.Data["title"] = "安徽黄金-打造黄金百年企业.谱写徽韵文化传奇"
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
	this.Data["dongtai"] = dongtai
	this.Data["xinwen"] = xinwen
	this.Data["view"] = 2
	this.TplName = "home/index1.html"
}
