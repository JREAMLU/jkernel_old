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
	i18n.Locale
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

	// beego.Trace("00000000000000:", i18n.Tr("zh-CN", "PARAMSILLEGAL"))

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

func (this *baseController) Tr(format string) string {
	var language = ""
	al := this.Ctx.Request.Header.Get("Accept-Language")
	if len(al) > 4 {
		al = al[:5] // Only compare first 5 letters.
		if i18n.IsExist(al) {
			language = al
		}
	}

	if len(language) == 0 {
		language = "en-US"
	}

	fmt.Println("6666666666666666", language)
	return i18n.Tr(language, format)
}
