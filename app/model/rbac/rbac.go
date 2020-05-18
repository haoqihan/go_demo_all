package rbac

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"go_demo_all/app/model/role"
	"go_demo_all/app/model/users"
)

// Common 定义全局对象
type Common struct {
	Enforcer *casbin.Enforcer `inject:""`
}

// LoadPolicyData 注入权限策略
func (a *Common) LoadPolicyData(id int, username string) error {
	fmt.Println("注入策略")
	// 查询用户角色数据
	userResults, err := users.Model.M.As("u").
		InnerJoin("user_role as ur", "u.id = ur.user_id").
		LeftJoin("role as r", "ur.role_id = r.id").
		Fields("u.username,ur.role_id,r.name").
		FindAll("u.id = ?", 1)
	if err != nil {
		glog.Error("查询用户、角色数据错误", err)
		return err
	}
	// 清除已有的权限数据
	_, _ = a.Enforcer.DeleteRolesForUser(username)
	// roleID 保存角色id的切片
	var roleID []uint
	// 注册用户、角色到CasBin 配置中
	for _, v := range userResults.List() {
		roleID = append(roleID, gconv.Uint(v["role_id"]))
		fmt.Println(gconv.String(v["username"]), gconv.String(v["name"]))
		_, _ = a.Enforcer.AddRoleForUser(gconv.String(v["username"]), gconv.String(v["name"]))
	}

	// 查询角色与路由对应的关系
	roleMeunResult, err := role.Model.M.As("r").
		InnerJoin("role_menu as rm", "r.id = rm.role_id").
		LeftJoin("menu as m", "rm.menu_id = m.id").
		Fields("r.name as role_name,m.name,m.path,m.method").FindAll("r.id in (?)", roleID)
	if err != nil {
		glog.Error("查询角色与路由错误", err)
		return err

	}

	for _, v := range roleMeunResult.List() {
		if v["path"] == "" || v["method"] == "" {
			continue
		}
		_, _ = a.Enforcer.AddPermissionForUser(gconv.String(v["role_name"]), gconv.String(v["path"]), gconv.String(v["method"]))
	}
	return nil
}
