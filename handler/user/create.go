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

// @Summary 注册
// @Description 注册新用户,通过一站式登录来注册
// @Tags user
// @Accept json
// @Produce json
// @Param loginInfo body model.LoginInfo true "学号和密码"
// @Success 200 {object} model.Res "{"msg":"success", "token": string}"
// @Failure 401 {object} handler.Error "{"error_code":"20001", "message":"Fail."} 注册失败
// @Failure 400 {object} handler.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} handler.Error "{"error_code":"30001", "message":"database does not open successful"} 失败"
// @Router /users [post]
func UserCreate(c *gin.Context) {
	var tmpUser model.SuInfo
	var tmpLoginInfo model.LoginInfo

	if err := c.BindJSON(&tmpLoginInfo); err != nil {
		handler.ErrBadRequest(c, err)
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
		c.JSON(401, gin.H{
			"msg": ok,
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
	claims := &model.JwtClaims{
		UID: uid,
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

func genToken(claims model.JwtClaims) (string, error) {
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

var (
	key        = "miniProject" //salt
	ExpireTime = 3600          //token expire time
)
