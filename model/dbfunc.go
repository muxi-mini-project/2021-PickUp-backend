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

func UpdateCommentD(score float64, words string, uid string) error {
	tmpComment, err := FindCommentD(uid)
	if err != nil {
		return err
	}
	tmpComment.DriverScore = (tmpComment.DriverScore + score) / 2.0
	tmpComment.Words = words + tmpComment.Words
	if err := Db.Self.Model(&CommentDriver{}).Where(CommentDriver{DriverID: uid}).Update(&tmpComment).Error; err != nil {
		return err
	}

	return nil
}

func FindCommentD(uid string) (CommentDriver, error) {
	var tmpComment CommentDriver
	if err := Db.Self.Model(&CommentDriver{}).Where(CommentDriver{DriverID: uid}).First(&tmpComment).Error; err != nil {
		return tmpComment, err
	}
	return tmpComment, nil
}

func CreateDriverRt(tmpRt RequireDriver) error {
	if _, err := FindDriverRt(tmpRt.DriverID); err == nil {
		if ok := DeleteDriverRt(tmpRt.DriverID); ok != nil {
			return ok
		}
	}

	if err := Db.Self.Create(&tmpRt).Error; err != nil {
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

func ConfirmD(rqt_id int, uid string) error {
	var tmpRt RequirePassenger
	var myerr handler.Error
	myerr.ErrorCode = "passenger does not confirm"
	myerr.Message = "wait!"
	err := Db.Self.Model(&RequirePassenger{}).Where(RequirePassenger{ID: rqt_id}).First(&tmpRt).Error
	if err != nil {
		return err
	}
	if tmpRt.Status == 1 {
		return &myerr
	}
	tmprt, err2 := FindDriverRt(uid)
	if err2 != nil {
		return err2
	}
	tmprt.Status = 2
	if err := Db.Self.Model(&RequireDriver{}).Where(RequireDriver{DriverID: uid}).Update(&tmprt).Error; err != nil {
		return err
	}

	return nil
}

func CreatePassengerRt(tmpRt RequirePassenger) error {
	if _, err := FindPassengerRt(tmpRt.PassengerID); err == nil {
		if ok := DeletePassengerRt(tmpRt.PassengerID); ok != nil {
			return ok
		}
	}

	if err := Db.Self.Create(&tmpRt).Error; err != nil {
		return err
	}

	return nil
}

func FindPassengerRt(uid string) (RequirePassenger, error) {
	var tmpRt RequirePassenger
	if err := Db.Self.Model(&RequirePassenger{}).Where(RequirePassenger{PassengerID: uid}).First(&tmpRt).Error; err != nil {
		return tmpRt, err
	}

	return tmpRt, nil
}

func DeletePassengerRt(uid string) error {
	tmpRt, err := FindPassengerRt(uid)
	if err != nil {
		return err
	}
	if err = Db.Self.Delete(&tmpRt).Error; err != nil {
		return err
	}
	return nil
}
