package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Category struct {
	Id      int
	Pid     int
	Status  int
	Level   int
	Name    string
	Index   int
	Created time.Time
	Updated time.Time
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

// func CategoryGetByPid(pid int) []*Category {
// 	list := make([]*Category, 0)
// 	orm.NewOrm().QueryTable(TableName("category").Filter("pid", pid)).All(&list)
// 	return list
// }
