package controllers

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

type NestPreparer interface {
	NestPrepare()
}

type baseController struct {
	beego.Controller
}

type langType struct {
	Lang string
	Name string
}

func (this *baseController) Prepare() {
	// Initialized language type list.
	langs := strings.Split(beego.AppConfig.String("lang::types"), "|")
	names := strings.Split(beego.AppConfig.String("lang::names"), "|")
	langTypes := make([]*langType, 0, len(langs))
	for i, v := range langs {
		langTypes = append(langTypes, &langType{
			Lang: v,
			Name: names[i],
		})
	}

	for _, lang := range langs {
		beego.Trace("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			beego.Error("Fail to set message file: " + err.Error())
			return
		}
	}
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
