package model

import (
	"fmt"
	"log"
	handler "pickup/handler/err"
	"strings"

	"github.com/jinzhu/gorm"
)

//const dsn = "root:123456@/PICKUP?charset=utf8&parseTime=True&loc=Local"

func CreateUser(tmpUser Users) error {

	var tooluser Users
	var myerr handler.Error
	myerr.ErrorCode = "users is exist!!"
	myerr.Message = "bad request!"
	err2 := Db.Self.Where(&Users{Sid: tmpUser.Sid}).Find(&tooluser).Error
	if err2 == nil {
		log.Println("creat user err: ", myerr)
		return &myerr
	}
	Db.Self.Create(&tmpUser)

	return nil
}

func Login(tmpLogin LoginInfo) int {
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		return 1
	}
	var user Users
	var logininfo LoginInfo
	err1 := db.Where(&Users{Sid: tmpLogin.Sid}).Find(&user).Error
	if err1 != nil {
		log.Println(err1)
		return 2
	}
	logininfo.Pwd = user.Password

	if logininfo.Pwd != tmpLogin.Pwd {
		return 3
	}

	return 4
}

func FindUser(uid string) (Users, error) {
	var tmpUser Users
	if err := Db.Self.Model(&Users{}).Where(Users{Sid: uid}).Find(&tmpUser).Error; err != nil {
		return tmpUser, err
	}
	return tmpUser, nil

}

func Updateuser(tmpUser Users, uid string) error {
	if err := Db.Self.Model(&Users{}).Where(Users{Sid: uid}).Error; err != nil {
		return err
	}

	if err := Db.Self.Model(&Users{}).Where(Users{Sid: uid}).Update(&tmpUser).Error; err != nil {
		return err
	}
	return nil

}

func UpdatePwd(tmpChange UpdatePwdInfo, uid string) error {
	tmpUser, err := FindUser(uid)
	if err != nil {
		return err
	}
	num := strings.Compare(tmpChange.Old, tmpUser.Password)

	//fmt.Println(num, tmpChange.Old)
	var myerr handler.Error
	myerr.ErrorCode = "password is not corret"
	myerr.Message = "bad request!"
	if num != 0 {
		return &myerr
	}
	tmpUser.Password = tmpChange.New

	err = Db.Self.Model(&Users{}).Where(Users{Sid: uid}).Update(&tmpUser).Error
	if err != nil {
		return err
	}

	return nil

}

func CreateDriverRt(tmpRt RequireDriver) error {

	err := Db.Self.Create(&tmpRt).Error
	if err != nil {
		return err
	}

	return nil
}

func FindDriverRt(uid string) (RequireDriver, error) {
	var tmpRt RequireDriver
	if err := Db.Self.Model(&RequireDriver{}).Where(RequireDriver{DriverID: uid}).First(&tmpRt).Error; err != nil {
		return tmpRt, err
	}

	return tmpRt, nil
}

func DeleteDriverRt(uid string) error {
	tmpRt, err := FindDriverRt(uid)
	if err != nil {
		return err
	}
	if err = Db.Self.Delete(&tmpRt).Error; err != nil {
		return err
	}
	return nil
}
