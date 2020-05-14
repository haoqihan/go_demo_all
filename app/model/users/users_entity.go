package users

type Entity struct {
	Id         int64  `orm:"id,primary" json:"id"`            // 主键
	Uuid       string `orm:"uuid" json:"uuid"`                // UUID
	Username   string `orm:"username,unique" json:"username"` // 登录名
	Password   string `orm:"password" json:"password"`        // 密码
	Salt       string `orm:"salt" json:"salt"`                // 密码盐
	RealName   string `orm:"real_name" json:"real_name"`      // 真实姓名
	DepartId   int    `orm:"depart_id" json:"depart_id"`      // 部门
	UserType   int    `orm:"user_type" json:"user_type"`      // 类型 1，管理员 2，普通用户 3，前台用户 4，第三方用户 5，api用户
	Status     int    `orm:"status" json:"status"`            // 状态
	ThirdId    string `orm:"third_id" json:"third_id"`        // 第三方id
	EndTime    string `orm:"end_time" json:"end_time"`        // 结束时间
	Email      string `orm:"email" json:"email"`              // email
	Tel        string `orm:"tel" json:"tel"`                  // 手机号
	Address    string `orm:"address" json:"address"`          // 地址
	TitleUrl   string `orm:"title_url" json:"title_url"`      // 头像地址
	Remark     string `orm:"remark" json:"remark"`            // 说明
	Theme      string `orm:"theme" json:"theme"`              // 主题
	UpdateTime string `orm:"update_time" json:"update_time"`  // 更新时间
	UpdateId   int64  `orm:"update_id" json:"update_id"`      // 更新人
	CreateTime string `orm:"create_time" json:"create_time"`  // 创建时间
	CreateId   int64  `orm:"create_id" json:"create_id"`      // 创建者
}
