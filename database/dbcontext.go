package database

import (
	_ "github.com/go-sql-driver/mysql" //
	"github.com/thg6060/exercise5/protoc"
	"xorm.io/xorm"
)

var db = DbConn()

//DbConn function for connect to database
func DbConn() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:123@tcp(0.0.0.0:7070)/exercise5")
	if err != nil {
		panic(err)
	}
	return engine
}

//CreateTable creates tables


func CreateTable() error {
	err := db.CreateTables(&protoc.UserPartner{})
	if err != nil {
		return err
	}
	err = db.Sync2(&protoc.UserPartner{})
	if err != nil {
		return err
	}
	return nil
}
