package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"go_demo_all/app/api/admin/login"
	"go_demo_all/app/middleware/token"
)

func init() {
	s := g.Server()
	s.Group("/v1", func(group *ghttp.RouterGroup) {
		group.POST("/admin/loginSubmit",login.GfJWTMiddleware.LoginHandler)
		group.Group("/admin", func(group *ghttp.RouterGroup) {
			//中间件检查token是否有效
			group.Middleware(token.Validator)
			//刷新token令牌
			group.GET("/refresh",login.GfJWTMiddleware.RefreshHandler)
		})
	})
}
