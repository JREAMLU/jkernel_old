package models

import (
	"core/db/mysql"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:@/base?charset=utf8", 30)
}

func GetUrlOne() orm.Params {
	params := []interface{}{1, 2, 3, "http://jream.lu"}
	sql := "SELECT * FROM redirect WHERE redirect_id IN (?, ?, ?) AND long_url = ? "
	a, _ := mysql.Select(params, sql)
	return a[0]
}
