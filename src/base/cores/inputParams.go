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
	Source    []string `json:"source" valid:"Required"`
	Version   []string `json:"version" valid:"Required"`
	SecretKey []string `json:"secret_key" valid:"Required"`
	RequestID []string `json:"request_id" valid:"Required"`
	Token     []string `json:"token" valid:"Required"`
	IP        []string `json:"ip" valid:"Required"`
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
func InputParamsCheck(meta map[string][]string, data ...interface{}) (result map[string]string, msg string, err error) {
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
				msg := i18n.Tr(global.Lang, "outputParams.DATAPARAMSILLEGAL") + " " + err.Key + ":" + err.Message
				return nil, msg, errors.New(i18n.Tr(global.Lang, "outputParams.DATAPARAMSILLEGAL"))
			}
		}
	}

	//MetaHeader check
	metaCheckResult, msg, err := MetaHeaderCheck(meta)
	if err != nil {
		return nil, msg, err
	}

	return metaCheckResult, "", nil
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
func MetaHeaderCheck(meta map[string][]string) (result map[string]string, msg string, err error) {
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
			msg := i18n.Tr(global.Lang, "outputParams.METAPARAMSILLEGAL") + " " + err.Key + ":" + err.Message
			return nil, msg, errors.New(i18n.Tr(global.Lang, "outputParams.METAPARAMSILLEGAL "))
		}
	}

	//把meta参数放入新的struct 返回
	var metaMap = make(map[string]string)
	for key, val := range meta {
		metaMap[key] = val[0]
	}

	//日志
	if len(metaMap["request_id"]) == 0 {
		metaMap["request_id"] = getRequestID()
	}

	return metaMap, "", nil
}

//request id增加
func getRequestID() string {
	return "RRRRRRRRRRRRRRRRRRRR"
}

//Token 验证
