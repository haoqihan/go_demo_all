package boot

import (
	"github.com/gogf/gf/frame/g"
	"go_demo_all/app/middleware/inject"
)

func init() {
	// 增加自定义配置文件
	g.Cfg("redis").SetFileName("redis.toml")
	// 初始化权限配置加载
	_ = inject.LoadCasBinPolicyData()

}
