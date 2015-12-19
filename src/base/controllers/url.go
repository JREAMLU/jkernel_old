package controllers

import (
	"base/cores/global"
	"base/services"

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
                "long_url": "http://o9d.cn",
                "IP": "127.0.0.1"
            },
            {
                "long_url": "http://huiyimei.com",
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
	shorten := su.GoShorten(rawMetaHeader, rawDataBody)

	r.Data["json"] = shorten
	r.ServeJson()

}

func (r *UrlController) GoExpand() {
}
