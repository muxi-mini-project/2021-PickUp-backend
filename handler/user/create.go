package user

import (
	"fmt"
	handler "pickup/handler/err"
	"pickup/model"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const dsn = "root:123456@/PICKUP?charset=utf8&parseTime=True&loc=Local"

func UserCreate(c *gin.Context) {
	var tmpUser model.SuInfo
	var tmpLoginInfo model.LoginInfo

	if err := c.BindJSON(&tmpLoginInfo); err != nil {
		fmt.Fprintln(c.Writer, "bad request!")
		return
	}

	tmpUser, err := model.GetUserInfoFormOne(tmpLoginInfo.Sid, tmpLoginInfo.Pwd)
	if err != nil {
		handler.ErrUnauthorized(c, &handler.Error{})
		return
	}

	user := model.Users{
		Sid:      tmpUser.User.Usernumber,
		NickName: tmpUser.User.DeptName,
		Password: tmpLoginInfo.Pwd,
		Gender:   3,
		Phone:    "null",
		Picture:  tmpLoginInfo.Pwd,
		Notes:    "null",
	}

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
	}

	db.Create(&user)
	ok := db.NewRecord(user)
	if ok {
		c.JSON(200, gin.H{
			"msg": "try_again",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg":   "success",
		"token": produceToken(tmpUser.User.Usernumber),
	})
	defer db.Close()
	return
}

func produceToken(s string) string {
	token := s + "233"

	return token
}
