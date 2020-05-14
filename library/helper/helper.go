package helper

import (
	"fmt"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"strings"
	"time"
)

// PageParam 检测分页参数，给定默认值与limit最大值
func PageParam(page, limit int) (int, int) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	if limit >= 50 {
		limit = 50
	}
	return page, limit
}

// TimeToString 将系统时间转换为 2020-05-14 22:28:05 形式
func TimeToString(expire time.Time) string {
	t, err := gtime.StrToTime(expire.Format(time.RFC3339))
	if err != nil {
		glog.Error("服务器内部错误", err)
		return ""
	}
	return gconv.String(t)
}

// OrderByParam 把map的interface 转换为字符串键值对
func OrderByParam(param []map[string]interface{}) string {
	var claims = param[0]
	var orderBy string
	for k, v := range claims {
		orderBy += fmt.Sprintf("%s %s ,", k, v)

	}
	orderBy = strings.TrimRight(orderBy, ",")
	return orderBy
}
