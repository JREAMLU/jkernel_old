package validate

import (
	"fmt"
	"log"

	"github.com/astaxie/beego/validation"
)

func InputParams(params ...interface{}) int {

	valid := validation.Validation{}

	for key, val := range params {
		is, err := valid.Valid(val)

		if err != nil {
			// handle error
			fmt.Println(key, "测试验证报错", err)
		}

		if !is {
			for _, err := range valid.Errors {
				log.Println(key, "测试验证不通过", err.Key, "-", err.Message)
			}
		}
	}
	return 1
}
