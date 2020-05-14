package users

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

type arModel struct {
	M *gdb.Model
}

var (
	Table = "users"
	Model = &arModel{g.DB("default").Table(Table).Safe()}
)

func FindOne(where ...interface{}) (*Entity, error) {

}
