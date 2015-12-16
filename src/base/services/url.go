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
	Urls []UrlsParams `valid:"Required"`
}

type UrlsParams struct {
	LongUrl string `valid:"Required"`
	IP      string `valid:"IP"`
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
	cores.InputParamsCheck(rawMetaHeader, &u.Data)

	/*
		//------------------------验证参数start------------------------
		//初始化验证
		validMetaData := validation.Validation{}
		validMeta := validation.Validation{}
		validData := validation.Validation{}

		//验证url下的meta,data
		vMetaData, errMD := validMetaData.Valid(&u)
		if errMD != nil {
			fmt.Println("meta,data: ", errMD)
			errParams.Err = errMD
			errParams.Message = "error"
			errParams.Status = 30000
			return "", errParams
		}
		if !vMetaData {
			for _, err := range validMetaData.Errors {
				fmt.Println("meta,data: ", err.Key, ":", err.Message)
				errParams.Status = 10000
				errParams.Err = errors.New(err.Key + ":" + err.Message)
				errParams.Message = err.Key + ":" + err.Message
				return "", errParams
			}
		}

		//验证meta信息

		vMeta, errM := validMeta.Valid(&u.Meta)
		if errM != nil {
			fmt.Println("meta: ", errM)
			errParams.Err = errM
			errParams.Message = "error"
			errParams.Status = 30000
			return "", errParams
		}
		if !vMeta {
			for _, err := range validMeta.Errors {
				fmt.Println("meta: ", err.Key, ":", err.Message)
				errParams.Status = 10000
				errParams.Err = errors.New(err.Key + ":" + err.Message)
				errParams.Message = err.Key + ":" + err.Message
				return "", errParams
			}
		}

		//验证data信息
		vData, errD := validData.Valid(&u.Data)
		if errD != nil {
			fmt.Println("data: ", errD)
			errParams.Err = errD
			errParams.Message = "error"
			errParams.Status = 30000
			return "", errParams
		}
		if !vData {
			for _, err := range validData.Errors {
				fmt.Println("data: ", err.Key, ":", err.Message)
				errParams.Status = 10000
				errParams.Err = errors.New(err.Key + ":" + err.Message)
				errParams.Message = err.Key + ":" + err.Message
				return "", errParams
			}
		}

		//------------------------验证参数end------------------------
	*/

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
