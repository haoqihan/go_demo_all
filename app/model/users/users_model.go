package users

import (
	"database/sql"
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

// FindOne按Model.WherePri和Model.One检索并返回单个记录。
// 另请参见Model.WherePri和Model.One。
func (m *arModel) FindOne(where ...interface{}) (*Entity, error) {
	one, err := m.M.FindOne(where...)
	if err != nil {
		return nil, err
	}
	var entity *Entity
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

//FindCount按Model.WherePri和Model.Count检索并返回记录号。
//具体另请参见Model.WherePri和Model.Count。
func (m *arModel) FindCount(where ...interface{}) (int, error) {
	return m.M.FindCount(where...)
}
