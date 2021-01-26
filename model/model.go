package model

import (
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

type Database struct {
	Self *gorm.DB
}

type latestAction struct {
	Sid        string `gorm:"sid"`
	LatestTime string `gorm:"latest_time"`
	RandNum    int    `gorm:"rand_num"`
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
	DriverID     string `json:"driver_id" gorm:"driver_id"`
	StartSpot    string `json:"start_spot" gorm:"start_spot"`
	StartTime    string `json:"start_time" gorm:"start_time"`
	EndTime      string `json:"end_time" gorm:"end_time"`
	PassingSpots string `json:"passing_spots" gorm:"passing_spots"`
	Notes        string `json:"notes" gorm:"notes"`
	Status       int    `json:"status" gorm:"status"`
}
