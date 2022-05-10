package echo

import "reflect"

// 中文

type Code struct {
	SUCCESS       string `code:"1" msg:"操作成功"`
	FAILD         string `code:"0" msg:"操作失败"`
	LOGIN_INVALID string `code:"204" msg:"未登录或者登录失效"`
	PARAMS_LOST   string `code:"104" msg:"参数必传"`
	SYSTEM_ERROR  string `code:"944" msg:"系统异常"`
}

func GetCode(code string) (string, string) {
	t := reflect.TypeOf(&Code{}).Elem()
	field, ok := t.FieldByName(code)
	if !ok {
		return "0", "unknown mistake"
	}
	return field.Tag.Get("code"), field.Tag.Get("msg")
}
