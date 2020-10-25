package models

// "github.com/astaxie/beego/orm"
import (
	"github.com/astaxie/beego/orm"
)

type Product struct {
	Id    int
	Image string
	Url   string
}

func (a *Product) TableName() string {
	return TableName("product")
}

func ProductGetById(id int) (*Product, error) {
	c := new(Product)

	if err := orm.NewOrm().QueryTable(TableName("product")).Filter("id", id).One(c); err != nil {
		return nil, err
	}
	return c, nil
}

func ProductLists(page, pagesize int) ([]*Product, int64) {
	offset := (page - 1) * pagesize
	list := make([]*Product, 0)
	query := orm.NewOrm().QueryTable(TableName("product"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pagesize, offset).All(&list)

	return list, total
}

func ProductDelete(id int) (int64, error) {
	query := orm.NewOrm().QueryTable(TableName("product"))
	return query.Filter("id", id).Delete()
}

func AddProduct(product *Product) (int64, error) {
	id, err := orm.NewOrm().Insert(product)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Product) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(r, fields...); err != nil {
		return err
	}
	return nil
}
