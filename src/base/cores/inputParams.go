package cores

import (
	"base/cores/global"
	"fmt"
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/beego/i18n"
	"github.com/pquerna/ffjson/ffjson"
)

type MetaHeader struct {
	Source    []string `valid:"Required"`
	Version   []string `valid:"Required"`
	SecretKey []string `valid:"Required"`
	RequestID []string `valid:"Required"`
	Token     []string `valid:"Required"`
	IP        []string `valid:"IP"`
}

/**
 *	@auther		jream.lu
 *	@intro		入参验证
 *	@logic
 *	@todo		返回值
 *	@data		data ...interface{}	切片指针
 *	@return 	?
 */
func InputParamsCheck(meta map[string][]string, data ...interface{}) int {
	//DataParams check
	valid := validation.Validation{}

	for _, val := range data {
		is, err := valid.Valid(val)

		//日志

		if err != nil {
			// handle error
			log.Println(i18n.Tr(global.Lang, "outputParams.SYSTEMILLEGAL"), err)
		}

		if !is {
			for _, err := range valid.Errors {
				log.Println(i18n.Tr(global.Lang, "outputParams.PARAMSILLEGAL"), err.Key, ":", err.Message)
			}
		}
	}

	//MetaHeader check
	MetaHeaderCheck(meta)
	return 1
}

/**
 * meta参数验证
 * 1.map转json
 * 2.json转slice
 * 3.解析到struct
 */
func MetaHeaderCheck(meta map[string][]string) int {
	rawMetaHeader, _ := ffjson.Marshal(meta)
	beego.Trace("入参meta:" + string(rawMetaHeader))
	var metaHeader MetaHeader
	ffjson.Unmarshal(rawMetaHeader, &metaHeader)
	fmt.Println("meta json解析:", metaHeader)

	return 1
}

//request id增加
