package user

/*import (
	"github.com/gin-gonic/gin"
)

func ViewUser(c *gin.Context) {
	uid := c.GetString("uid")
	tmpUser, err := model.FindUser(uid)
	if err != nil {
		ErrServerError(c, error2.ServerError)
		return
	}
	c.JSON(200, gin.H{
		"msg":      "Success",
		"sid":      tmpUser.Sid,
		"nickname": tmpUser.NickName,
		"college":  tmpUser.College,
		"gender":   tmpUser.Gender,
		"grade":    tmpUser.Grade,
		"portrait": tmpUser.Portrait,
	})
	return
}*/
