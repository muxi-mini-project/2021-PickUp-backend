package user

import (
	"fmt"
	"log"
	handler "pickup/handler/err"
	"pickup/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

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
		NickName: tmpUser.User.Name,
		Password: tmpLoginInfo.Pwd,
		Gender:   3,
		Phone:    "null",
		Picture:  tmpLoginInfo.Sid,
		Notes:    "null",
	}

	ok := model.CreateUser(user)
	if ok != nil {
		c.JSON(200, gin.H{
			"msg": "try_again",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg":   "success",
		"token": produceToken(tmpUser.User.Usernumber),
	})
	return
}

func produceToken(uid string) string {
	//id, _ := strconv.Atoi(uid)
	claims := &jwtClaims{
		Uid: uid,
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	singedToken, err := genToken(*claims)
	//fmt.Println(singedToken, err)
	if err != nil {
		log.Print("produceToken err:")
		fmt.Println(err)
		return ""
	}
	return singedToken
}

func genToken(claims jwtClaims) (string, error) {
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

type jwtClaims struct {
	jwt.StandardClaims
	Uid string `json:"uid"`
}

var (
	key        = "miniProject" //salt
	ExpireTime = 3600          //token expire time
)
