package cores

import (
	"base/cores/global"
	"errors"
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
	RequestID []string
	Token     []string `valid:"Required"`
	IP        []string `valid:"Required"`
}

/**
 *	@auther		jream.lu
 *	@intro		入参验证
 *	@logic
 *	@todo		返回值
 *	@meta		meta map[string][]string	   rawMetaHeader
 *	@data		data ...interface{}	切片指针	rawDataBody
 *	@return 	返回 true, metaMap, error
 */
func InputParamsCheck(meta map[string][]string, data ...interface{}) (inputCheckResult interface{}, err error) {
	//DataParams check
	valid := validation.Validation{}

	for _, val := range data {
		is, err := valid.Valid(val)

		//日志

		//检查参数
		if err != nil {
			// handle error
			log.Println(i18n.Tr(global.Lang, "outputParams.SYSTEMILLEGAL"), err)
		}

		if !is {
			for _, err := range valid.Errors {
				log.Println(i18n.Tr(global.Lang, "outputParams.DATAPARAMSILLEGAL"), err.Key, ":", err.Message)
				result := i18n.Tr(global.Lang, "outputParams.DATAPARAMSILLEGAL") + " " + err.Key + ":" + err.Message
				return result, errors.New(i18n.Tr(global.Lang, "outputParams.DATAPARAMSILLEGAL"))
			}
		}
	}

	//MetaHeader check
	metaCheckResult, err := MetaHeaderCheck(meta)
	if err != nil {
		return metaCheckResult, err
	}

	return metaCheckResult, nil
}

/**
 * meta参数验证
 * 1.map转json
 * 2.json转slice
 * 3.解析到struct
 * 4.将header 放入map 返回
 *
 * @meta 	meta  map[string][]string 	header信息 map格式
 */
func MetaHeaderCheck(meta map[string][]string) (metaCheckResult interface{}, err error) {
	rawMetaHeader, _ := ffjson.Marshal(meta)
	beego.Trace("入参meta:" + string(rawMetaHeader))
	var metaHeader MetaHeader
	ffjson.Unmarshal(rawMetaHeader, &metaHeader)

	//日志
	fmt.Println("meta json解析:", metaHeader)
	for key, val := range meta {
		fmt.Println("meta 解析", key, ":", val[0])
	}

	valid := validation.Validation{}

	is, err := valid.Valid(&metaHeader)

	//日志

	//检查参数
	if err != nil {
		// handle error
		log.Println(i18n.Tr(global.Lang, "outputParams.SYSTEMILLEGAL"), err)
	}

	if !is {
		for _, err := range valid.Errors {
			log.Println(i18n.Tr(global.Lang, "outputParams.METAPARAMSILLEGAL"), err.Key, ":", err.Message)
			result := i18n.Tr(global.Lang, "outputParams.METAPARAMSILLEGAL") + " " + err.Key + ":" + err.Message
			return result, errors.New(i18n.Tr(global.Lang, "outputParams.METAPARAMSILLEGAL "))
		}
	}

	//把meta参数放入新的struct 返回
	var metaMap = make(map[string]string)
	for key, val := range meta {
		metaMap[key] = val[0]
	}

	//日志
	if len(metaMap["Requestid"]) == 0 {
		metaMap["Requestid"] = getRequestID()
	}

	return metaMap, nil
}

//request id增加
func getRequestID() string {
	return "RRRRRRRRRRRRRRRRRRRR"
}

//Token 验证
