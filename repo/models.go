package repo

import (
	"fmt"
	"github.com/alecthomas/log4go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	dbUser   = "root"
	dbPwd    = "Ro0T123456"
	database = "ACCESSTOOL"
	dbHost   = "localhost"
)

var Db *gorm.DB
var repository *AccessRepo

type AccessRepo struct {
	Rule
	Role
	Team
}

type Model interface {
	Add(repo AccessRepo) error
	Delete(repo AccessRepo) error
	Update(info map[string]interface{}) error
}

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local", dbUser, dbPwd, dbHost, 3306, database))
	if err != nil {
		_ = log4go.Error(err)
		panic(err)
	}
	defer db.Close()

	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(5)
	return db, nil
}

func InitRepo() {
	var err error
	Db, err = InitDB()
	if err != nil {
		_ = log4go.Error(err)
		panic(err)
	}
	Db.AutoMigrate(&Team{}, &Role{}, &Rule{})

}

func init() {
	repository = new(AccessRepo)
}
