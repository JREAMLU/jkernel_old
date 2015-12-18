package services

import (

	//"encoding/json"

	"fmt"

	"base/cores"
	"base/cores/url"

	"github.com/pquerna/ffjson/ffjson"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type Url struct {
	Data DataParams `valid:"Required"`
}

type DataParams struct {
	Urls []UrlsParams `json:"urls" valid:"Required"`
}

type UrlsParams struct {
	LongUrl string `json:"long_url" valid:"Required"`
	IP      string `json:"ip" valid:"IP"`
}

type ErrParams struct {
	Status  int
	Err     error
	Message string
}

type dataList struct {
	Total int                    `json:"total"`
	List  map[string]interface{} `json:"list"`
}

func GetParams(url Url) Url {
	return url
}

/**
 *	自定义多valid
 */
func (r *Url) Valid(v *validation.Validation) {}

/**
 *	@auther		jream.lu
 *	@intro		原始链接=>短链接
 *	@logic
 *	@todo		参数验证抽出去
 *	@params		params []byte	参数
 *	@return 	slice
 */
func (r *Url) GoShorten(rawMetaHeader map[string][]string, rawDataBody []byte) (shortUrl interface{}, errParams ErrParams) {
	//将传递过来多json raw解析到struct
	var u Url
	ffjson.Unmarshal(rawDataBody, &u)
	// ffjson.Unmarshal(rawDataBody, &u.Data.Urls)

	//日志
	fmt.Println("Url json解析:", u)

	//测试嵌套验证
	checkedMessage, err := cores.InputParamsCheck(rawMetaHeader, &u.Data)
	if err != nil {
		fmt.Println("err----------", checkedMessage)
		cores.OutputSuccess(checkedMessage)
	}
	fmt.Println("checked----------", checkedMessage)

	//进行shorten
	var list = make(map[string]interface{})
	for _, val := range u.Data.Urls {
		list[val.LongUrl] = url.GetShortenUrl(val.LongUrl, beego.AppConfig.String("ShortenDomain"))
	}

	var dl dataList
	dl.List = list
	dl.Total = len(list)

	//持久化到mysql

	errParams.Status = 0
	return dl, errParams
}
