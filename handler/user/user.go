package user

import (
	handler "pickup/handler/err"
	"pickup/model"

	"github.com/gin-gonic/gin"
)

func ViewUser(c *gin.Context) {
	var token *model.JwtClaims
	var s string
	s = c.GetHeader("token")
	//fmt.Println(s)
	token, err2 := model.VerifyToken(s)
	//fmt.Println(token)
	if err2 != nil {
		handler.ErrTokenInvalid(c, err2)
		return
	}
	uid := token.UID
	tmpUser, err := model.FindUser(uid)
	if err != nil {
		handler.ErrServerError(c, err)
		return
	}
	c.JSON(200, gin.H{
		"msg":       "Success",
		"sid":       tmpUser.Sid,
		"nick_name": tmpUser.NickName,
		"gender":    GetGender(tmpUser.Gender),
		"picture":   tmpUser.Picture,
		"phone":     tmpUser.Phone,
		"notes":     tmpUser.Notes,
	})
	return
}

func UpdateUser(c *gin.Context) {
	var token *model.JwtClaims
	var s string
	s = c.GetHeader("token")
	//fmt.Println(s)
	token, err2 := model.VerifyToken(s)
	//fmt.Println(token)
	if err2 != nil {
		handler.ErrTokenInvalid(c, err2)
		return
	}
	uid := token.UID
	var tmpchange model.Users
	err := c.BindJSON(&tmpchange)
	if err != nil {
		handler.ErrBadRequest(c, err)
		return
	}
	if err = model.Updateuser(tmpchange, uid); err != nil {
		handler.ErrUnauthorized(c, err)
		return
	}

	c.JSON(200, gin.H{
		"msg": "success",
	})
	return

}

func UpdatePassword(c *gin.Context) {
	var token *model.JwtClaims
	var s string
	s = c.GetHeader("token")
	//fmt.Println(s)
	token, err2 := model.VerifyToken(s)
	//fmt.Println(token)
	if err2 != nil {
		handler.ErrTokenInvalid(c, err2)
		return
	}
	uid := token.UID
	var tmpchange model.UpdatePwdInfo
	err := c.BindJSON(&tmpchange)
	if err != nil {
		handler.ErrBadRequest(c, err)
		return
	}

	err = model.UpdatePwd(tmpchange, uid)
	if err != nil {
		handler.ErrBadRequest(c, err)
		return
	}

	c.JSON(200, gin.H{
		"msg": "success",
	})

}

func GetGender(num int) string {
	if num == 1 {
		return "女"
	} else if num == 2 {
		return "男"
	} else {
		return "未知"
	}

}
