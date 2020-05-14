package response

import (
	"encoding/json"
	"github.com/gogf/gf/util/gconv"
	"go_demo_all/library/e"
)

// Resp 返回结构体
type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// JSON 返回json字符串
func (resp Resp) JSON() string {
	str, _ := json.Marshal(resp)
	return string(str)
}

// Success 成功
func Success(data interface{}) Resp {
	return Resp{200, "success", data}
}

// Fail 失败
func Fail(code int) Resp {
	msg := e.GetMsg(code)
	return Resp{code, msg, []int{}}
}

// Error 错误
func Error(code int) Resp {
	msg := e.GetMsg(code)
	return Resp{code, gconv.String(msg), []int{}}
}

// FailParam 校验参数提示信息
func FailParam(msg string) Resp {
	return Resp{400, msg, []int{}}
}
