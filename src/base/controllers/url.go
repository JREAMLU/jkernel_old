package controllers

import (
	"base/cores/global"
	"base/services"
	"fmt"

	"github.com/astaxie/beego"

	//"encoding/json"
)

type UrlController struct {
	global.BaseController
}

/**
{
    "data": {
        "urls": [
            {
                "LongUrl": "http://o9d.cn",
                "IP": "127.0.0.1"
            },
            {
                "LongUrl": "http://huiyimei.com",
                "IP": "192.168.1.1   "
            }
        ]
    }
}
*/
/**
 *	@auther		jream.lu
 *	@url		https://base.jream.lu/v1/url/goshorten.json
 *	@todo 		参数验证, 封装返回
 */
func (r *UrlController) GoShorten() {
	//接受参数 json raw
	rawMetaHeader := r.Ctx.Input.Request.Header
	rawDataBody := r.Ctx.Input.RequestBody

	//记录参数日志
	beego.Trace("入参body:" + string(rawDataBody))

	//调用servcie方法, 将参数传递过去
	var su services.Url
	shorten, err := su.GoShorten(rawMetaHeader, rawDataBody)

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
