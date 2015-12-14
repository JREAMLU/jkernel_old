package cores

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

const (
	SUCCESS       = 0
	PARAMSILLEGAL = 10000 //参数错误
	LOGICILLEGAL  = 20000 //逻辑错误
	SYSTEMILLEGAL = 30000 //系统错误
)

type outputParams struct {
	beego.Controller
}

var Lang string

var StatusCode = map[int]string{SUCCESS: i18n.Tr(Lang, "outputParams.SUCCESS"), PARAMSILLEGAL: "参数错误", LOGICILLEGAL: "逻辑错误", SYSTEMILLEGAL: "系统错误"}

type Output struct {
	Meta       MetaList
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       dataList
}

type MetaList struct {
	RequestId string `json:"request_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type dataList struct {
	Total int                    `json:"total"`
	List  map[string]interface{} `json:"list"`
}

func (op *outputParams) init() {
	al := op.Ctx.Request.Header.Get("Accept-Language")
	if len(al) > 4 {
		al = al[:5] // Only compare first 5 letters.
		if i18n.IsExist(al) {
			Lang = al
		}
	}

	if len(Lang) == 0 {
		Lang = "en-US"
	}
}

/**
 *	@auther		jream.lu
 *	@intro		出参成功
 *	@logic
 *	@todo		返回值
 *	@params		params ...interface{}	切片指针
 *	@return 	?
 */
func OutputSucc(o Output) Output {
	o.Meta.RequestId = "abc-111"
	return o
}

func OutputFail() {

}
