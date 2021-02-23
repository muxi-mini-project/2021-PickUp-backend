package driver

import (
	handler "pickup/handler/err"
	"pickup/model"

	"github.com/gin-gonic/gin"
)

func DriverConfirm(c *gin.Context) {
	pid := c.Query("pid") //前端发送要确认的乘客id
	did := c.Query("did")
	err := model.ConfirmD(pid, did)
	if err != nil {
		handler.ErrServerError(c, err)
		//fmt.Println(2)
		return
	}

	c.JSON(200, gin.H{
		"msg": "success",
	})

}
