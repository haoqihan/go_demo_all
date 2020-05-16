package login

import (
	"errors"
	"fmt"
	jwt "github.com/gogf/gf-jwt"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
	"go_demo_all/app/model/users"
	"go_demo_all/library/base"
	"go_demo_all/library/helper"
	"go_demo_all/library/input"
	"go_demo_all/library/redis"
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
		LoginResponse:   PostLogin,                                          // 完成登录后返回的信息，用户可自定义返回数据，默认返回
		RefreshResponse: RefreshResponse,                                    // 刷新token后返回的信息，用户可自定义返回数据，默认返回
		Unauthorized:    Unauthorized,                                       // 处理不进行授权的逻辑
		IdentityHandler: IdentityHandler,                                    // 解析并设置用户身份信息
		PayloadFunc:     PayloadFunc,                                        // 登录期间的回调的函数

	})
	if err != nil {
		glog.Error("JWT ERROR" + err.Error())

	}
	GfJWTMiddleware = authMiddleWare
}

// Authenticator 检测身份信息是否正常
func Authenticator(r *ghttp.Request) (interface{}, error) {
	var req *SignRequest
	// 接收参数
	input.JSONTOStruct(r, &req)
	// 校验数据参数
	if err := gvalid.CheckStruct(req, nil); err != nil {
		base.FailParam(r, err.String())
	}
	// 查询数据
	res := users.GetOne(g.Map{"username": req.Username})

	if res.Id <= 0 {
		return nil, errors.New("用户名或密码错误")
	}

	reqPwd, errPwd := gmd5.Encrypt(req.Password + res.Salt)
	if errPwd != nil {
		glog.Error("md5加密异常", errPwd)
		return nil, errors.New("服务器异常")
	}
	if reqPwd != res.Password {
		return nil, errors.New("用户名或密码错误")
	}
	// 设置参数保存到请求中
	r.SetParam("uuid", res.Uuid)
	fmt.Println("auth 认证")

	return g.Map{"username": res.Username, "uuid": res.Uuid}, nil
}

// PostLogin
func PostLogin(r *ghttp.Request, code int, token string, expire time.Time) {
	j, _ := r.GetJson()
	// 格式化时间
	t := helper.TimeToString(expire)
	fmt.Printf("%+v",r)

	// 获取配置文件中redis前缀
	var loginPrefix = g.Cfg("redis").Get("APP.LOGIN_PREFIX")
	redis.Set(gconv.String(loginPrefix)+gconv.String(r.GetParam("uuid")), token)
	base.Success(r, g.Map{
		"username": j.GetString("username"),
		"token":    token,
		"expire":   t,
	})

}

// RefreshResponse 刷新token信息
func RefreshResponse(r *ghttp.Request, code int, token string, expire time.Time) {
	fmt.Println(code,token,expire)
	var loginPrefix = g.Cfg("redis").Get("APP.LOGIN_PREFIX")
	// 重新设置该用户的token信息
	redis.Set(gconv.String(loginPrefix)+gconv.String(r.GetParam("uuid")), token)
	base.Success(r, g.Map{"token": token, "expire": helper.TimeToString(expire)})
}

// Unauthorized 返回验证错误的信息
func Unauthorized(r *ghttp.Request, code int, message string) {
	// TODO 错误提示英文转换为中文，最好做一个配置文件
	fmt.Println("失败",message)
	base.FailParam(r, message)
}

// IdentityHandler 设置JWT的标识。
func IdentityHandler(r *ghttp.Request) interface{} {
	claims := jwt.ExtractClaims(r)
	return claims["id"]
}

// PayloadFunc 给token添加其他字段信息
func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}
/*

paramsMap:map[
		JWT_PAYLOAD:map[
				exp:1.589642243e+09
				orig_iat:1.589641943e+09
				username:admin
				uuid:94091b1fa6ac4a27a06c0b92155aea6a
				]
		JWT_TOKEN:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODk2NDIyNDMsIm9yaWdfaWF0IjoxNTg5NjQxOTQzLCJ1c2VybmFtZSI6ImFkbWluIiwidXVpZCI6Ijk0MDkxYjFmYTZhYzRhMjdhMDZjMGI5MjE1NWFlYTZhIn0.9aL3Qgy-XR6_3Ym9iXM6wmXrb1dSg8UeqACNIDKxV5s
		username:admin
]



*/