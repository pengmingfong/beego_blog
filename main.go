package main

import (
	"gold/models"
	_ "gold/routers"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	models.Init()
	beego.Run()
}
