package repo

type Rule struct {
	Id     int
	Name   string `gorm:"column:name;type:varchar(100)"`
	TeamId int    `gorm:"column:team_id;type:int"`
	Role   []int  `gorm:"column:role;type:varchar(200)"`
	Custom int    `gorm:"column:custom;type:varchar(100)"`
}

func RuleTableName() string {
	return "rule"
}

func NewRule(id int, name string, teamId int, role []int, custom int) *Rule {
	return &Rule{Id: id, Name: name, TeamId: teamId, Role: role, Custom: custom}
}

func (r *Rule) Add(re AccessRepo) error {
	return Db.Create(&re.Rule).Error
}

func (r *Rule) Delete(re AccessRepo) error {
	return Db.Delete(&re).Error
}

func (r *Rule) Update(info map[string]interface{}) error {
	rule := Rule{}
	return Db.Model(&rule).Updates(info).Error
}

func (r *Rule) Query() (*[]Rule, error) {
	rule := make([]Rule, 0)
	db := Db.Find(&rule)
	return db.Value.(*[]Rule), db.Error
}

func (r *Rule) QueryOne(id int) (*Rule, error) {
	rule := Rule{}
	if err := Db.First(&rule, id).Error; err != nil {
		return nil, err
	}
	return &rule, nil
}
