package login

import (
	jwt "github.com/gogf/gf-jwt"
	"github.com/gogf/gf-jwt/example/auth"
	"github.com/gogf/gf/os/glog"
	"time"
)

var (
	GfJWTMiddleware *jwt.GfJWTMiddleware // 声明jwt包的全局变量
)

type SignRequest struct {
	Username string `v:required#账号不能为空 json:"username"`
	Password string `v:required#密码不能为空 json:"password"`
}

func init() {
	authMiddleWare, err := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           "realm",                                            // 用于展示中间件名称
		Key:             []byte("秘钥key"),                                    // 秘钥
		Timeout:         5 * time.Minute,                                    // token 过期时间
		MaxRefresh:      5 * time.Minute,                                    // 这是在任何刷新令牌端点上调用的提供的函数。如果传入的令牌是在该MaxRefreshTime时间范围内发出的，则此处理程序将创建/设置一个类似于的新令牌LoginHandler，并将该令牌传递到RefreshResponse
		IdentityKey:     "id",                                               // 身份验证的key值
		TokenLookup:     "header: Authorization, query: token, cookie: jwt", // token检索模式，用于提取token-> Authorization
		TokenHeadName:   "Bearer",                                           // token在请求头时的名称,客户端在header中传入Authorization 对一个值是Bearer + 空格 + token
		TimeFunc:        time.Now,                                           // 测试或服务器在其他时区可设置该属性
		Authenticator:   Authenticator,                                      // 根据登录信息对用户进行身份验证的回调函数
		LoginResponse:   LoginResponse,                                      // 完成登录后返回的信息，用户可自定义返回数据，默认返回
		RefreshResponse: auth.RefreshResponse,                               // 刷新token后返回的信息，用户可自定义返回数据，默认返回
		Unauthorized:    auth.Unauthorized,                                  // 处理不进行授权的逻辑
		IdentityHandler: auth.IdentityHandler,                               // 解析并设置用户身份信息
		PayloadFunc:     auth.PayloadFunc,                                   // 登录期间的回调的函数

	})
	if err != nil {
		glog.Error("JWT ERROR" + err.Error())

	}
	GfJWTMiddleware = authMiddleWare
}
