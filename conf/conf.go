package conf

import (
	"fmt"
	"go_demo_all/cache"
	"go_demo_all/model"
	"go_demo_all/util"
	"os"


	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))
	fmt.Println(os.Getenv("LOG_LEVEL"))

	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("翻译文件加载失败", err)
	}

	// 连接数据库
	fmt.Println(os.Getenv("MYSQL_DSN"))
	model.Database(os.Getenv("MYSQL_DSN"))
	cache.Redis()
}
