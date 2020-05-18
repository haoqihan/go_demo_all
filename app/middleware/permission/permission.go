package permission

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"go_demo_all/app/middleware/inject"
	"go_demo_all/library/base"
	"go_demo_all/library/e"
)

// CasBinMiddleware 检测当前用户是否具有访问权限
func CasBinMiddleware(r *ghttp.Request) {
	fmt.Println(r.GetParam("username"), r.RequestURI, r.Method)
	if ok, err := inject.Obj.Enforcer.Enforce(r.GetParam("username"), r.RequestURI, r.Method); err != nil {
		base.Error(r, e.Error)
	} else if !ok {
		base.Error(r, e.Forbidden)
	}
	fmt.Println("xxxxxxxxxxxxxx")
	r.Middleware.Next()
}


