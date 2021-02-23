package user

import (
	handler "pickup/handler/err"
	"pickup/model"

	"github.com/gin-gonic/gin"
)

// @Summary 显示用户信息
// @Description 显示用户信息，点击“我的”和“修改信息”的时候调用
// @Tags user
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} model.Users "成功"
// @Failure 401 {object} handler.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} handler.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} handler.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /users [get]
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
		"sid":       tmpUser.Sid,
		"nick_name": tmpUser.NickName,
		"gender":    GetGender(tmpUser.Gender),
		"picture":   tmpUser.Picture,
		"phone":     tmpUser.Phone,
		"notes":     tmpUser.Notes,
	})
	return
}

// @Summary 修改用户信息
// @Description 修改用户信息，点击“修改信息”的时候调用
// @Tags user
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param user body model.VeUsers true "userinfo"
// @Success 200 {object} model.Users "成功"
// @Failure 401 {object} handler.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} handler.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} handler.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /users/ [put]
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
