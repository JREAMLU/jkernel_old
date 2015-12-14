package main

import (
	_ "base/docs"
	_ "base/routers"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}

	beego.AddFuncMap("i18n", i18n.Tr)

	beego.Run()
}
