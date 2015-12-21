package inout

import (
	"core/global"
	"time"

	"github.com/beego/i18n"
)

const (
	SUCCESS           = 0
	DATAPARAMSILLEGAL = 10000
	METAPARAMSILLEGAL = 15000
	LOGICILLEGAL      = 20000
	SYSTEMILLEGAL     = 30000
)

type Output struct {
	Meta       MetaList    `json:"meta"`
	StatusCode int         `json:"status_code"`
	Message    interface{} `json:"message"`
	Data       interface{} `json:"data"`
}

type MetaList struct {
	RequestId string `json:"request_id"`
	UpdatedAt string `json:"updated_at"`
}

/*
type dataList struct {
	Total int                    `json:"total"`
	List  map[string]interface{} `json:"list"`
}
*/

/**
 *	@auther		jream.lu
 *	@intro		出参成功
 *	@logic
 *	@todo		返回值
 *	@params		params ...interface{}	切片指针
 *	@return 	?
 */
func OutputSuccess(data interface{}, requestID string) Output {
	var op Output
	op.Meta.RequestId = requestID
	op.Meta.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	op.StatusCode = SUCCESS

	op.Message = i18n.Tr(global.Lang, "outputParams.SUCCESS")

	op.Data = data

	return op
}

func OutputFail(msg interface{}, status string, requestID string) Output {
	var op Output
	op.Meta.RequestId = requestID
	op.Meta.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	switch status {
	case "SUCCESS":
		op.StatusCode = SUCCESS
	case "DATAPARAMSILLEGAL":
		op.StatusCode = DATAPARAMSILLEGAL
	case "METAPARAMSILLEGAL":
		op.StatusCode = METAPARAMSILLEGAL
	case "LOGICILLEGAL":
		op.StatusCode = LOGICILLEGAL
	case "SYSTEMILLEGAL":
		op.StatusCode = SYSTEMILLEGAL
	}

	op.Message = msg

	op.Data = make(map[string]interface{})

	return op
}
