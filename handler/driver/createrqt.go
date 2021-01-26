package driver

import (
	"fmt"
	handler "pickup/handler/err"
	"pickup/model"

	"github.com/gin-gonic/gin"
)

func AddDriverRequirement(c *gin.Context) {
	var tmprequirement model.RequireDriver
	if err := c.BindJSON(&tmprequirement); err != nil {
		handler.ErrBadRequest(c, err)
		fmt.Println("1")
		return
	}
	if err := model.CreateDriverRt(tmprequirement); err != nil {
		handler.ErrServerError(c, err)
		fmt.Println("2")
		return
	}

	c.JSON(200, gin.H{
		"msg": "success",
	})
}
