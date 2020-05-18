package inject

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/facebookgo/inject"
	"github.com/gogf/gf/net/ghttp"
	"go_demo_all/app/model/rbac"
	"go_demo_all/library/base"
	"go_demo_all/library/e"
)

type CasBinObj struct {
	Common   *rbac.Common
	Enforcer *casbin.Enforcer
}

// Obj CasBinObj 实例化变量
var Obj *CasBinObj

// 初始化CasbinRBAC配置
func init() {
	i := new(inject.Graph)
	fmt.Println("初始化CasbinRBAC配置")
	var r *ghttp.Request
	var path = "config/rbac/rbac.conf"
	enforcer, err := casbin.NewEnforcer(path, false)
	if err != nil {
		base.Error(r, e.Error)
	}
	_ = i.Provide(&inject.Object{Value: enforcer})

	common := new(rbac.Common)
	_ = i.Provide(&inject.Object{Value: common})

	if err := i.Populate(); err != nil {
		base.Error(r, e.Error)
	}
	Obj = &CasBinObj{
		Common:   common,
		Enforcer: enforcer,
	}
	return
}

func LoadCasBinPolicyData() error {
	m := Obj.Common

	// 注入用户和角色信息
	err := m.LoadPolicyData(1, "admin")
	if err != nil {
		return err
	}
	return nil
}

