package controllers

import (
	"base/service"
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
	//接受参数 json raw
	params := r.Ctx.Input.RequestBody

	//调用servcie方法, 将参数传递过去
	var su service.Url
	shorten, err := su.GoShorten(params)

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
