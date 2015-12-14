package controllers

import (
	"base/services"
	"fmt"

	//"encoding/json"
)

type UrlController struct {
	baseController
}

/**
 *	@auther		jream.lu
 *	@url		https://base.jream.lu/v1/url/goshorten.json
 *	@todo 		参数验证, 封装返回
 */
func (r *UrlController) GoShorten() {
	fmt.Println("9999999999999999999999999", r.Tr("outputParams.PARAMSILLEGAL"))

	//接受参数 json raw
	rawDataBody := r.Ctx.Input.RequestBody
	rawMetaHeader := r.Ctx.Input.Request.Header
	for key, val := range rawMetaHeader {
		fmt.Println(key, ": ", val[0])
	}

	//调用servcie方法, 将参数传递过去
	var su services.Url
	shorten, err := su.GoShorten(rawDataBody, rawMetaHeader)

	var data = make(map[string]interface{})
	if err.Err != nil {
		fmt.Println(err.Status, "-", err.Message)

		data["meta"] = ""
		data["status"] = err.Status
		data["data"] = err.Message
	} else {
		//定义json输出值
		data["meta"] = ""
		data["status"] = 0
		data["data"] = shorten
	}

	r.Data["json"] = data
	r.ServeJson()

}

func (r *UrlController) GoExpand() {
}
