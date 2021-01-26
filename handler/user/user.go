package user

import (
	handler "pickup/handler/err"
	"pickup/model"

	"github.com/gin-gonic/gin"
)

func ViewUser(c *gin.Context) {
	uid := c.Param("uid")
	tmpUser, err := model.FindUser(uid)
	if err != nil {
		handler.ErrServerError(c, err)
		return
	}
	c.JSON(200, gin.H{
		"msg":       "Success",
		"sid":       tmpUser.Sid,
		"nick_name": tmpUser.NickName,
		"gender":    tmpUser.Gender,
		"picture":   tmpUser.Picture,
		"phone":     tmpUser.Phone,
		"notes":     tmpUser.Notes,
	})
	return
}

func UpdateUser(c *gin.Context) {
	uid := c.Param("uid")
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
	uid := c.Param("uid")
	var tmpchange model.UpdatePwdinfo
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
