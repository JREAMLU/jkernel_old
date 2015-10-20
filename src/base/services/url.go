package service

import (
	"crypto/md5"
	"encoding/hex"
	//"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/pquerna/ffjson/ffjson"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type MetaParams struct {
	Source    string `valid:"Required"`
	Version   string `valid:"Required"`
	SecretKey string `valid:"Required"`
	RequestID string `valid:"Required"`
	IP        string `valid:"IP"`
}

type UrlsParams struct {
	LongUrl string `valid:"Required"`
	IP      string `valid:"IP"`
}

type DataParams struct {
	Urls []UrlsParams `valid:"Required"`
}

type Url struct {
	Meta MetaParams `valid:"Required"`
	Data DataParams `valid:"Required"`
}

//10000开始参数错误
//20000开始逻辑错误
//30000开始系统错误
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
func (r *Url) Valid(v *validation.Validation) {

}

/**
 *	@auther		jream.lu
 *	@intro		原始链接=>短链接
 *	@logic
 *	@todo		参数验证抽出去
 *	@params		params []byte	参数
 *	@return 	slice
 */
func (r *Url) GoShorten(params []byte) (shortUrl interface{}, errParams ErrParams) {
	fmt.Println("接受的参数:", string(params))

	//将传递过来多json raw解析到struct
	var u Url
	ffjson.Unmarshal(params, &u)
	ffjson.Unmarshal(params, &u.Data.Urls)
	fmt.Println("json解析:", u)

	//测试嵌套验证
	vtest := validation.Validation{}
	b, err := vtest.Valid(&u)
	if err != nil {
		// handle error
		fmt.Println("测试验证报错", err)
	}
	if !b {
		for _, err := range vtest.Errors {
			log.Println("测试验证不通过", err.Key, "-", err.Message)
		}
	}

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

	//进行shorten
	var list = make(map[string]interface{})
	for _, val := range u.Data.Urls {
		list[val.LongUrl] = GetShortenUrl(val.LongUrl, beego.AppConfig.String("ShortUrl"))
	}

	var dl dataList
	dl.List = list
	dl.Total = len(list)

	//持久化到mysql

	errParams.Status = 0
	return dl, errParams
}

//----------------------------------------------短网址算法--------------------------------------------------
func GetShortenUrl(url, mainUrl string) string {
	//随机
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	random := r.Intn(4)

	shortUrl := ShortUrl(url)
	sUrl := shortUrl[random]

	lastUrl := mainUrl + sUrl
	fmt.Println(lastUrl)
	return lastUrl
}

func ShortUrl(url string) []string {
	//62位的密码
	base62 := [...]string{"a", "b", "c", "d", "e", "f", "g", "h",
		"i", "j", "k", "l", "m", "n", "o", "p",
		"q", "r", "s", "t", "u", "v", "w", "x",
		"y", "z", "A", "B", "C", "D", "E", "F",
		"G", "H", "I", "J", "K", "L", "M", "N",
		"O", "P", "Q", "R", "S", "T", "U", "V",
		"W", "X", "Y", "Z", "0", "1", "2", "3",
		"4", "5", "6", "7", "8", "9"}

	//	fmt.Println(base62)
	//传进来的url进行md5
	h := md5.New()
	h.Write([]byte(url))

	//hex.EncodeToString(h.Sum(nil)) -> md5
	hex := hex.EncodeToString(h.Sum(nil))
	hexLen := len(hex)
	subHexLen := hexLen / 8

	var outPut []string
	for i := 0; i < subHexLen; i++ {
		//截取md5的长度 按8个 截4段
		subHex := SubString(hex, i*8, 8)

		//每段与0x组成16进制
		strHex := "0x" + subHex

		//将string转换成int 16进制
		sH, _ := strconv.ParseInt(strHex, 0, 64)

		//做运算
		valInt := 0x3FFFFFFF & (1 * sH)

		var out string
		for j := 0; j < 6; j++ {
			val := 0x0000003D & valInt
			out += base62[val]
			valInt = valInt >> 5
			//			fmt.Println(valInt)
		}

		//添加到outPut 切片里
		outPut = append(outPut, out)
		//		fmt.Println(outPut)
	}

	return outPut
}

/**
 * 截取字符串长度
 */
func SubString(str string, begin, length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)

	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}

	// 返回子串
	return string(rs[begin:end])
}
