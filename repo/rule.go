package repo

type Rule struct {
	Id     int    `gorm:"column:id;primary_key;type:int;not null"`
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

func (r *Rule) Add() error {
	return nil
}
