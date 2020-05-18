package user

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"go_demo_all/app/model/users"
	"go_demo_all/library/base"
	"go_demo_all/library/helper"
	"go_demo_all/library/input"
)

// Controller 定义操作的结构体
type Controller struct {}

// Index 显示用户列表信息
func (c *Controller) Index(r *ghttp.Request) {
	fmt.Println("显示用户列表信息")
	var req = input.ListParams(r)

	// 设置分页的默认值和limit的最大值
	page, limit := helper.PageParam(req.Page, req.Limit)

	// 获取数据
	total, result := users.GetList(page, limit, req.Where, req.OrderBy)
	// 返回结果集
	base.Success(r, g.Map{"total": total, "page": req.Page, "lists": result})

}
