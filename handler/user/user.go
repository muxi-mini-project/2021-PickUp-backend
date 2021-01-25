package user

import (
	handler "pickup/handler/err"
	"pickup/model"

	"github.com/gin-gonic/gin"
)

func ViewUser(c *gin.Context) {
	uid := c.GetString("uid")
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
