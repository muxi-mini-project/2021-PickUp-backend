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

func UpdateCommentD(tmpComment CommentDriver, uid string) error {
	exist, err := FindCommentD(uid)
	if err != nil {
		if err := Db.Self.Model(&CommentDriver{}).Create(&tmpComment).Error; err != nil {
			return err
		}
		return nil
	}
	exist.DriverScore = (exist.DriverScore + tmpComment.DriverScore) / 2.0
	exist.Words = exist.Words + tmpComment.Words
	if err := Db.Self.Model(&CommentDriver{}).Where(CommentDriver{DriverID: uid}).Update(&exist).Error; err != nil {
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

func UpdateCommentP(tmpComment CommentPassenger, uid string) error {
	exist, err := FindCommentP(uid)
	if err != nil {
		if err := Db.Self.Model(&CommentPassenger{}).Create(&tmpComment).Error; err != nil {
			return err
		}
		return nil
	}
	exist.PassengerScore = (exist.PassengerScore + tmpComment.PassengerScore) / 2.0
	exist.Words = exist.Words + tmpComment.Words
	if err := Db.Self.Model(&CommentPassenger{}).Where(CommentPassenger{PassengerID: uid}).Update(&exist).Error; err != nil {
		return err
	}

	return nil
}

func FindCommentP(uid string) (CommentPassenger, error) {
	var tmpComment CommentPassenger
	if err := Db.Self.Model(&CommentPassenger{}).Where(CommentPassenger{PassengerID: uid}).First(&tmpComment).Error; err != nil {
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

func ConfirmP(uid string) error {
	tmpRt, err := FindPassengerRt(uid)
	if err != nil {
		return err
	}
	tmpRt.Status = 2
	if err := Db.Self.Model(&RequirePassenger{}).Where(RequirePassenger{PassengerID: uid}).Update(&tmpRt).Error; err != nil {
		return err
	}

	return nil
}
func ConfirmD(pid string, uid string) error {
	var tmpRt RequirePassenger
	var myerr handler.Error
	myerr.ErrorCode = "passenger does not confirm"
	myerr.Message = "wait!"
	err := Db.Self.Model(&RequirePassenger{}).Where(RequirePassenger{PassengerID: pid}).First(&tmpRt).Error
	fmt.Println(1)
	if err != nil {
		return err
	}
	if tmpRt.Status == 1 {
		return &myerr
	}
	fmt.Println(2)
	tmprt, err2 := FindDriverRt(uid)
	if err2 != nil {
		return err2
	}
	fmt.Println(3)
	tmprt.Status = 2
	if err := Db.Self.Model(&RequireDriver{}).Where(RequireDriver{DriverID: uid}).Update(&tmprt).Error; err != nil {
		return err
	}

	match := Match{
		DriverID:    tmprt.DriverID,
		PassengerID: tmpRt.PassengerID,
		StartTime:   tmprt.StartTime,
		EndTime:     tmprt.EndTime,
		StartSpot:   tmprt.StartSpot,
		EndSpot:     tmpRt.EndSpot,
		DriverPhone: GetPhone(tmprt.DriverID),
	}

	if err := CreateRoute(match); err != nil {
		return err
	}

	if err := DeleteDriverRt(tmprt.DriverID); err != nil {
		return err
	}
	if err := DeletePassengerRt(tmpRt.PassengerID); err != nil {
		return err
	}

	return nil
}

func GetPhone(uid string) string {
	user, err := FindUser(uid)
	if err != nil {
		return "null"
	}
	return user.Phone
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

func CreateRoute(tmpRoute Match) error {
	err := Db.Self.Model(&Match{}).Create(&tmpRoute).Error
	if err != nil {
		return err
	}
	return nil
}

func FindRoute(uid string) ([]Match, error) {
	var routes []Match
	err := Db.Self.Model(&Match{}).Where(Match{UserID: uid}).Find(&routes).Error
	if err != nil {
		return routes, err
	}
	return routes, nil
}

func FindRoute2(id int) (Match, error) {
	var route Match
	err := Db.Self.Model(&Match{}).Where(Match{ID: id}).Find(&route).Error
	if err != nil {
		return route, err
	}
	return route, nil
}

func DeleteRoute(id int) error {
	var route Match
	err := Db.Self.Model(&Match{}).Where(Match{ID: id}).Delete(&route).Error
	if err != nil {
		return err
	}

	return nil
}
