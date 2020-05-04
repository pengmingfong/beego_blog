package models

import (
	// "time"

	"github.com/astaxie/beego/orm"
)

type Article struct {
	Id          int
	Cid         int
	Title       string
	Sub_title   string
	Small_title string
	Image       string
	Content     string
	Tag         string
	Keyword     string
	Status      int
	Index       int
	Created     string
	Updated     string
}

func (a *Article) TableName() string {
	return TableName("article")
}

func ArticleById(id int) (*Article, error) {
	c := new(Article)

	if err := orm.NewOrm().QueryTable(TableName("article")).Filter("id", id).One(c); err != nil {
		return nil, err
	}
	return c, nil

}

func ArticleByCid(cid int, limit int) []*Article {
	list := make([]*Article, 0)
	orm.NewOrm().QueryTable(TableName("article")).Filter("cid", cid).Limit(limit).All(list)
	return list
}

func ArticleByCids(page, pagesize, cid int, condition interface{}) ([]*Article, int64) {
	offset := (page - 1) * pagesize
	list := make([]*Article, 0)
	query := orm.NewOrm().QueryTable(TableName("article")).Filter("cid", cid).Filter("status", 1)

	total, _ := query.Count()
	query.OrderBy("-id").Limit(pagesize, offset).All(&list)

	return list, total
}

func IndexList(page int) ([]*Article, int64) {
	list := make([]*Article, 0)
	query := orm.NewOrm().QueryTable(TableName("article")).Filter("index", 1).Filter("status", 1)
	total, _ := query.Count()
	query.OrderBy("-id").All(&list)

	return list, total
}

func List(page, pagesize int) ([]*Article, int64) {
	offset := (page - 1) * pagesize
	list := make([]*Article, 0)
	query := orm.NewOrm().QueryTable(TableName("article"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pagesize, offset).All(&list)

	return list, total
}

func Addarticle(article *Article) (int64, error) {
	id, err := orm.NewOrm().Insert(article)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Article) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(r, fields...); err != nil {
		return err
	}
	return nil
}
