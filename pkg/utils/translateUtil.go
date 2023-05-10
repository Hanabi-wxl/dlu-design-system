package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

var (
	msgMap = map[string]string{}
)

func init() {
	msgMap["Name.required"] = "姓名不可以为空"
	msgMap["Name.min"] = "姓名长度大于{min}位"
	msgMap["Name.max"] = "姓名长度小于{max}位"

	msgMap["Password.required"] = "密码不可以为空"
	msgMap["Password.min"] = "密码长度大于{min}位"
	msgMap["Password.max"] = "密码长度小于{max}位"

	msgMap["Number.required"] = "学/工号不可以为空"
	msgMap["Number.min"] = "学/工号长度为{min}位"
	msgMap["Number.max"] = "学/工号长度为{max}位"

	msgMap["Phone.required"] = "手机号不可以为空"

	msgMap["IsStu.required"] = "用户类型不可以为空"
}

// TranslateOverride 翻译错误信息
func TranslateOverride(err error) string {
	errStr := strings.Builder{}
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			key := fmt.Sprintf("%v.%v", e.Field(), e.Tag())
			if _, ok := msgMap[key]; ok {
				if e.Param() != "" {
					errStr.WriteString(strings.Replace(msgMap[key], "{"+e.Tag()+"}", e.Param(), -1) + " ")
				} else {
					errStr.WriteString(msgMap[key] + " ")
				}
			} else {
				errStr.WriteString(key + "未定义翻译字段" + " ")
			}
		}
	}
	return errStr.String()
}
