package model

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const dns = "root:1234@/PICKUP?charset=utf8&parseTime=True"

var Db *Database

func getDatabase() (*gorm.DB, error) {
	/*dns := fmt.Sprintf("%s:%s@tcp(%s)/PICKUP",
	viper.GetString("db.username"),
	viper.GetString("db.password"),
	viper.GetString("db.addr"))*/
	db, err := gorm.Open("mysql", dns)
	if err != nil {
		fmt.Print("getDatabase")
		log.Println(err)
		db = nil
	}
	db.SingularTable(true)
	db.LogMode(true)
	return db, err
}

func (db *Database) Init() {
	newDb, err := getDatabase()
	if err != nil {
		log.Printf("DBInit")
		fmt.Println(err)
	}
	Db = &Database{Self: newDb}
}

func (db *Database) Close() error {
	if err := Db.Self.Close(); err != nil {
		return err
	}
	return nil
}
