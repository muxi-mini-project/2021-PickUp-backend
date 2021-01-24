package model

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

const dns = "root:123456@tcp(127.0.0.1:3306)/users"

var Db *Database

func getDatabase() (*gorm.DB, error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s)/mini_project",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"))
	//dns := fmt.Sprintf("%s:%s@tcp(localhost:3306)/mini_project", os.Getenv("DBUser"), os.Getenv("DBPassword"))
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
