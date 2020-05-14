package users

type Entity struct {
	Id int64 `orm:"id,primary" json:"id"`
}
