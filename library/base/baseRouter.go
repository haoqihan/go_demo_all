package base

import (
	"github.com/gogf/gf/net/ghttp"
	"go_demo_all/library/response"
)

// Success 返回成功
func Success(r *ghttp.Request, data interface{}) {
	_ = r.Response.WriteJson(response.Success(data))
	r.Exit()
}

// Fail 操作失败返回
func Fail(r *ghttp.Request, code int) {
	_ = r.Response.WriteJson(response.Fail(code))
	r.Exit()
}

// Error 请求错误返回
func Error(r *ghttp.Request, code int) {
	_ = r.Response.WriteJson(response.Error(code))
	r.Exit()
}

// FailParam 校验返回错误提示信息
func FailParam(r *ghttp.Request, msg string) {
	_ = r.Response.WriteJson(response.FailParam(msg))
	r.Exit()
}
