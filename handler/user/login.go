package user

import (
	"fmt"
	"pickup/model"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	var tmpLoginInfo model.LoginInfo

	if err := c.BindJSON(&tmpLoginInfo); err != nil {
		fmt.Fprintln(c.Writer, "bad request!")
		return
	}

	ok := model.Login(tmpLoginInfo)
	if ok == 1 {
		c.JSON(500, gin.H{
			"msg": "database does not open successful",
		})
		return
	}

	if ok == 2 {
		c.JSON(404, gin.H{
			"msg": "user does not exist",
		})
		return
	}

	if ok == 3 {
		c.JSON(200, gin.H{
			"msg": "password does not correct",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg":   "success",
		"token": produceToken(tmpLoginInfo.Sid),
	})

}
