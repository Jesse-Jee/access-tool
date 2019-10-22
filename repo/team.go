package repo

type Team struct {
	Id       int
	TeamName string `gorm:"column:name;type:varchar(200)"`
}

func TeamTableName() string {
	return "team"
}

func NewTeam(id int, teamName string) *Team {
	return &Team{Id: id, TeamName: teamName}
}

func (t *Team) Add(re AccessRepo) error {
	return Db.Create(&re.Team).Error
}

func (t *Team) Delete(re AccessRepo) error {
	return Db.Delete(&re).Error
}

func (t *Team) Update(info map[string]interface{}) error {
	team := Team{}
	return Db.Model(&team).Updates(info).Error
}

func (t *Team) Query() (*[]Team, error) {
	teams := make([]Team, 0)
	db := Db.Find(&teams)
	return db.Value.(*[]Team), db.Error
}

func (t *Team) QueryTeamById(teamId int) (*Team, error) {
	team := Team{}
	if err := Db.First(&team, teamId).Error; err != nil {
		return nil, err
	}
	return &team, nil
}

func (t *Team) QueryTeamByName(name string) (*Team, error) {
	team := Team{}
	if err := Db.First(&team, name).Error; err != nil {
		return nil, err
	}
	return &team, nil
}
