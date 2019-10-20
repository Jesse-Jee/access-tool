package repo

import (
	"database/sql"
	"github.com/alecthomas/log4go"
)
import _ "github.com/go-sql-driver/mysql"

type Model interface {
	Add() error
	Delete() error
	Modify() error
	Get() error
}

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:Ro0T123456@/ACCESSTOOL")
	if err != nil {
		_ = log4go.Error(err)
		panic(err)
	}
	defer db.Close()
	return db
}
