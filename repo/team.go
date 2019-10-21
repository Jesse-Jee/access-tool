package repo

type Team struct {
	Id       int    `gorm:"column:id;primary_key;type:int;not null"`
	TeamName string `gorm:"column:name;type:varchar(200)"`
}

func TeamTableName() string {
	return "team"
}

func NewTeam(id int, teamName string) *Team {
	return &Team{Id: id, TeamName: teamName}
}

func (t *Team) Add() error {

	return nil
}
