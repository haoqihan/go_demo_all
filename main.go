package main

import (
	_ "go_demo_all/boot"
	_ "go_demo_all/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
