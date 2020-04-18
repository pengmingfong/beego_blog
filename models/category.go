package models

import (
	"github.com/astaxie/beego/orm"
)

type Category struct {
	Id      int
	Pid     int
	Status  int
	Level   int
	Name    string
	Index   int
	Created string
	Updated string
}

func (a *Category) TableName() string {
	return TableName("category")
}

func CategoryGetById(id int) (*Category, error) {
	c := new(Category)

	if err := orm.NewOrm().QueryTable(TableName("category")).Filter("id", id).One(c); err != nil {
		return nil, err
	}
	return c, nil
}

func CategoryListParent() []*Category {
	list := make([]*Category, 0)
	orm.NewOrm().QueryTable(TableName("category")).Filter("pid", 0).All(&list)

	return list
}

func CategoryList() []*Category {
	list := make([]*Category, 0)
	orm.NewOrm().QueryTable(TableName("category")).All(&list)

	return list
}

func CategoryAllList() []*Category {
	list := make([]*Category, 0)
	orm.NewOrm().QueryTable(TableName("category")).All(&list)

	return list
}

func CategoryLists(page, pagesize int) ([]*Category, int64) {
	offset := (page - 1) * pagesize
	list := make([]*Category, 0)
	query := orm.NewOrm().QueryTable(TableName("category"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pagesize, offset).All(&list)

	return list, total
}

func CategoryDelete(id int) (int64, error) {
	query := orm.NewOrm().QueryTable(TableName("category"))
	return query.Filter("id", id).Delete()
}

func AddCategory(category *Category) (int64, error) {
	id, err := orm.NewOrm().Insert(category)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Category) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(r, fields...); err != nil {
		return err
	}
	return nil
}
