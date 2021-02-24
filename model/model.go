package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type Users struct {
	Sid      string `json:"-" gorm:"sid"`
	NickName string `json:"nick_name" gorm:"nick_name"`
	Password string `json:"password" gorm:"password"`
	Gender   int    `json:"gender" gorm:"gender"`
	Phone    string `json:"phone" gorm:"gender"`
	Picture  string `json:"picture" gorm:"picture"`
	Notes    string `json:"notes" gorm:"notes"`
}

type VeUsers struct {
	NickName string `json:"nick_name" gorm:"nick_name"`
	Gender   int    `json:"gender" gorm:"gender"`
	Phone    string `json:"phone" gorm:"gender"`
	Picture  string `json:"picture" gorm:"picture"`
	Notes    string `json:"notes" gorm:"notes"`
}

type Database struct {
	Self *gorm.DB
}

type LoginInfo struct {
	Sid string `json:"sid"`
	Pwd string `json:"pwd"`
}

/*type User struct {
	Sid  string `grom:"sid"`
	Name string `grom:"name"`
}*/

type UpdatePwdInfo struct {
	Old string `json:"old"`
	New string `json:"new"`
}

type RequireDriver struct {
	ID        int    `json:"id" gorm:"id"`
	DriverID  string `json:"driver_id" gorm:"driver_id"`
	StartSpot string `json:"start_spot" gorm:"start_spot"`
	//年月日
	Ymd          string `json:"ymd" gorm:"ymd"`
	StartTime    string `json:"start_time" gorm:"start_time"`
	EndTime      string `json:"end_time" gorm:"end_time"`
	PassingSpots string `json:"passing_spots" gorm:"passing_spots"`
	Notes        string `json:"notes" gorm:"notes"`
	Status       int    `json:"status" gorm:"status"`
}

type VeRequireDriver struct {
	StartSpot string `json:"start_spot" gorm:"start_spot"`
	//年月日
	Ymd          string `json:"ymd" gorm:"ymd"`
	StartTime    string `json:"start_time" gorm:"start_time"`
	EndTime      string `json:"end_time" gorm:"end_time"`
	PassingSpots string `json:"passing_spots" gorm:"passing_spots"`
	Notes        string `json:"notes" gorm:"notes"`
}

type VeRequireDriver1 struct {
	DriverID    string `json:"driver_id" gorm:"driver_id"`
	DriverPhone string `json:"driver_phone" gorm:"driver_phone"`
	Status      string `json:"status" gorm:"status"`
	StartSpot   string `json:"start_spot" gorm:"start_spot"`
	//年月日
	Ymd          string `json:"ymd" gorm:"ymd"`
	StartTime    string `json:"start_time" gorm:"start_time"`
	EndTime      string `json:"end_time" gorm:"end_time"`
	PassingSpots string `json:"passing_spots" gorm:"passing_spots"`
	Notes        string `json:"notes" gorm:"notes"`
}

type RequirePassenger struct {
	ID          int    `json:"id" gorm:"id"`
	PassengerID string `json:"passenger_id" gorm:"passenger_id"`
	StartSpot   string `json:"start_spot" gorm:"start_spot"`
	EndSpot     string `json:"end_spot" gorm:"end_spot"`
	//年月日
	Ymd       string `json:"ymd" gorm:"ymd"`
	StartTime string `json:"start_time" gorm:"start_time"`
	EndTime   string `json:"end_time" gorm:"end_time"`
	Urgent    int    `json:"urgent" gorm:"urgent"`
	Notes     string `json:"notes" gorm:"notes"`
	Status    int    `json:"status" gorm:"status"`
}

type VeRequirePassenger struct {
	StartSpot string `json:"start_spot" gorm:"start_spot"`
	EndSpot   string `json:"end_spot" gorm:"end_spot"`
	//年月日
	Ymd       string `json:"ymd" gorm:"ymd"`
	StartTime string `json:"start_time" gorm:"start_time"`
	EndTime   string `json:"end_time" gorm:"end_time"`
	//紧急情况,1为紧急,2为非紧急,如果没有则视为非紧急
	Urgent int    `json:"urgent" gorm:"urgent"`
	Notes  string `json:"notes" gorm:"notes"`
	Status int    `json:"status" gorm:"status"`
}

type CommentDriver struct {
	ID          int     `json:"id" gorm:"id"`
	DriverID    string  `json:"driver_id" gorm:"driver_id"`
	DriverScore float64 `json:"driver_score" gorm:"driver_score"`
	Words       string  `json:"words" gorm:"words"`
}

type CommentPassenger struct {
	ID             int     `json:"id" gorm:"id"`
	PassengerID    string  `json:"passenger_id" gorm:"passenger_id"`
	PassengerScore float64 `json:"passenger_score" gorm:"passenger_score"`
	Words          string  `json:"words" gorm:"words"`
}

type Comment struct {
	DriverID       string  `json:"driver_id" `
	DriverScore    float64 `json:"driver_score" `
	DriverWords    string  `json:"words" `
	PassengerID    string  `json:"passenger_id" `
	PassengerScore float64 `json:"passenger_score" `
	PassengerWords string  `json:"passenger_words"`
}

type Match struct {
	ID          int    `json:"id" gorm:"id"`
	UserID      string `json:"user_id" gorm:"user_id"`
	DriverID    string `json:"driver_id" gorm:"driver_id"`
	PassengerID string `json:"passenger_id" gorm:"passenger_id"`
	Ymd         string `json:"ymd" gorm:"ymd"`
	StartTime   string `json:"start_time" gorm:"start_time"`
	EndTime     string `json:"end_time" gorm:"end_time"`
	StartSpot   string `json:"start_spot" gorm:"start_spot"`
	EndSpot     string `json:"end_spot" gorm:"end_spot"`
	DriverPhone string `json:"driver_phone" gorm:"driver_phone"`
}

type Route struct {
	UserID    string `json:"user_id" gorm:"user_id"`
	StartTime string `json:"start_time" gorm:"start_time"`
	EndTime   string `json:"end_time" gorm:"end_time"`
	StartSpot string `json:"start_spot" gorm:"start_spot"`
	EndSpot   string `json:"end_spot" gorm:"end_spot"`
}

type VeRoute struct {
	StartSpot string `json:"start_spot" gorm:"start_spot"`
	EndSpot   string `json:"end_spot" gorm:"end_spot"`
	UserPhone string `json:"user_phone" gorm:"user_phone"`
}

type JwtClaims struct {
	jwt.StandardClaims
	UID string `json:"uid"`
}

var Secret = "miniProject"

type Res struct {
	Msg string `json:"msg"`
}
