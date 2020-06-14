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

type TreeList struct {
	Id       int         `json:"id"`
	Name     string      `json:"name"`
	Pid      int         `json:"pid"`
	Children []*TreeList `json:"children"`
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

func CategoryBypid(pid int) []*Category {
	list := make([]*Category, 0)
	orm.NewOrm().QueryTable(TableName("category")).Filter("pid", pid).All(&list)
	return list
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

func GetChild(pid int) []*TreeList {
	o := orm.NewOrm()
	var menu []Category
	_, _ = o.QueryTable(TableName("category")).Filter("pid", pid).OrderBy("level").All(&menu)
	treeList := []*TreeList{}
	for _, v := range menu {
		if v.Index == 0 {
			continue
		}

		child := GetChild(v.Id)
		node := &TreeList{
			Id:   v.Id,
			Name: v.Name,
			Pid:  v.Pid,
		}
		node.Children = child
		treeList = append(treeList, node)
	}
	return treeList
}
