package model

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

const dsn = "root:123456@/PICKUP?charset=utf8&parseTime=True&loc=Local"

func CreateUser(tmpUser Users) error {

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
	}
	var tooluser Users
	err2 := db.Where(&Users{Sid: tmpUser.Sid}).Find(&tooluser).Error
	if err2 != nil {
		log.Println("creat user err: ", err2)
		return err2
	}
	db.Create(&tmpUser)
	defer db.Close()

	return nil
}
