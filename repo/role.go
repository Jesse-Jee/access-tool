package repo

type Role struct {
	Id   int
	Name string `gorm:"column:name;type:varchar(100)"`
	Auth []int  `gorm:"column:auth;type:varchar(100)"`
}

func RoleTableName() string {
	return "role"
}

func NewRole(id int, name string, auth []int) *Role {
	return &Role{Id: id, Name: name, Auth: auth}
}

func (r *Role) Add(re AccessRepo) error {
	return Db.Create(&re.Role).Error
}

func (r *Role) Delete(re AccessRepo) error {
	return Db.Delete(&re).Error
}

func (r *Role) Update(info map[string]interface{}) error {
	role := Role{}
	return Db.Model(&role).Updates(info).Error
}

func (r *Role) Query() (*[]Role, error) {
	role := make([]Role, 0)
	db := Db.Find(&role)
	return db.Value.(*[]Role), db.Error
}
