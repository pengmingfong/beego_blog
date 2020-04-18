package models

// "github.com/astaxie/beego/orm"
import (
	"github.com/astaxie/beego/orm"
)

type Banner struct {
	Id    int
	Image string
	Level int
}

func (a *Banner) TableName() string {
	return TableName("banner")
}

func BannerGetById(id int) (*Banner, error) {
	c := new(Banner)

	if err := orm.NewOrm().QueryTable(TableName("banner")).Filter("id", id).One(c); err != nil {
		return nil, err
	}
	return c, nil
}

func BannerLists(page, pagesize int) ([]*Banner, int64) {
	offset := (page - 1) * pagesize
	list := make([]*Banner, 0)
	query := orm.NewOrm().QueryTable(TableName("banner"))
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pagesize, offset).All(&list)

	return list, total
}

func BannerDelete(id int) (int64, error) {
	query := orm.NewOrm().QueryTable(TableName("banner"))
	return query.Filter("id", id).Delete()
}

func AddBanner(banner *Banner) (int64, error) {
	id, err := orm.NewOrm().Insert(banner)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Banner) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(r, fields...); err != nil {
		return err
	}
	return nil
}
