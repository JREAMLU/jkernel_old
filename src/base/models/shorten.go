package models

import (
	"core/db/mysql"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	alias      = beego.AppConfig.String("db::alias")
	driver     = beego.AppConfig.String("db::driver")
	username   = beego.AppConfig.String("db::username")
	password   = beego.AppConfig.String("db::password")
	database   = beego.AppConfig.String("db::database")
	charset    = beego.AppConfig.String("db::charset")
	maxIdle, _ = beego.AppConfig.Int("db::maxIdle")
)

func init() {
	orm.RegisterDataBase(alias, driver, username+":@/"+database+"?charset="+charset, maxIdle)
}

func GetUrlOne() orm.Params {
	params := []interface{}{1, 2, 3, "http://jream.lu"}
	sql := "SELECT * FROM redirect WHERE redirect_id IN (?, ?, ?) AND long_url = ? "
	a, _ := mysql.Select(params, sql)
	return a[0]
}
