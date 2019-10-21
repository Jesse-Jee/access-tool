package repo


type Role struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Auth []int  `json:"auth"`
}

func RoleTableName() string {
	return "rule"
}

func NewRole(id int, name string, auth []int) *Role {
	return &Role{Id: id, Name: name, Auth:auth}
}

func (r *Role) Add() error {
	return nil
}
