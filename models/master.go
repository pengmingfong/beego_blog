package models

// "github.com/astaxie/beego/orm"
import (
	"github.com/astaxie/beego/orm"
)

type Master struct {
	Id    int
	Image string
	Url   string
}

func (a *Master) TableName() string {
	return TableName("master")
}

func MasterGetById(id int) (*Master, error) {
	c := new(Master)

	if err := orm.NewOrm().QueryTable(TableName("master")).Filter("id", id).One(c); err != nil {
		return nil, err
	}
	return c, nil
}

func MasterLists(page, pagesize int) ([]*Master, int64) {
	offset := (page - 1) * pagesize
	list := make([]*Master, 0)
	query := orm.NewOrm().QueryTable(TableName("master"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pagesize, offset).All(&list)

	return list, total
}

func MasterDelete(id int) (int64, error) {
	query := orm.NewOrm().QueryTable(TableName("master"))
	return query.Filter("id", id).Delete()
}

func AddMaster(master *Master) (int64, error) {
	id, err := orm.NewOrm().Insert(master)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Master) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(r, fields...); err != nil {
		return err
	}
	return nil
}
