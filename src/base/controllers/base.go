package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type NestPreparer interface {
	NestPrepare()
}

type baseController struct {
	beego.Controller
}

func (this *baseController) Prepare() {

}

func GoGetAllAppConfig() map[string]interface{} {
	var core = make(map[string]interface{})
	core["author"] = beego.AppConfig.String("core::author")
	core["ProjectName"] = beego.AppConfig.String("core::ProjectName")

	for key, value := range core {
		fmt.Println(key, ":", value)
	}

	return core
}

func (this *baseController) GoGetAuthor() string {
	return beego.AppConfig.String("core::author")
}

func (this *baseController) GoGetProjectName() string {
	return beego.AppConfig.String("core::ProjectName")
}

func (this *baseController) GogetUrlDomain() string {
	return beego.AppConfig.String("ShortUrl")
}
