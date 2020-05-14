package token

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"go_demo_all/app/api/admin/login"
	"go_demo_all/library/base"
	"go_demo_all/library/e"
	"go_demo_all/library/redis"
)

// 验证token信息
func Validator(r *ghttp.Request) {
	login.GfJWTMiddleware.MiddlewareFunc()(r)
	// 解析token
	parseToken, _ := login.GfJWTMiddleware.ParseToken(r)
	var token = parseToken.Raw
	// 解析token中保存的信息
	var claims = gconv.Map(parseToken.Claims)
	r.SetParam("username", claims["username"])
	if !GetRedisToken(gconv.String(claims["uuid"]), token) {
		base.Fail(r, e.ErrorAuthCheckTokenFail)
	}
	r.Middleware.Next()
}

// GetRedisToken 获取缓存中的token与客户端进行对比
func GetRedisToken(uuid string, oldToken string) bool {
	redisPrefix := gconv.String(g.Cfg("redis").Get("APP.LOGIN_PREFIX"))
	key := redisPrefix + uuid
	if redis.Get(key) != oldToken {
		return false
	}
	return true
}
