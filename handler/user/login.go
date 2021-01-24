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

}
