package router

import (
    "go_demo_all/app/api/admin/login"
    "github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/",login.GfJWTMiddleware.LoginHandler)
	})
}
