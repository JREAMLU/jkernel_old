package cores

const (
	SUCCESS       = 0
	PARAMSILLEGAL = 10000 //参数错误
	LOGICILLEGAL  = 20000 //逻辑错误
	SYSTEMILLEGAL = 30000 //系统错误
)

var StatusCode = map[int]string{SUCCESS: "SUCCESS", PARAMSILLEGAL: "参数错误", LOGICILLEGAL: "逻辑错误", SYSTEMILLEGAL: "系统错误"}

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
