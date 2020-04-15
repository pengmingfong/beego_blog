package models

import (
	"time"
	// "github.com/astaxie/beego/orm"
)

type User struct {
	Id             int
	Username       string
	Password       string
	Last_logintime time.Time
}
