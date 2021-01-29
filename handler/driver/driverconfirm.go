package driver

import (
	handler "pickup/handler/err"
	"pickup/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DriverConfirm(c *gin.Context) {
	prt_id := c.Param("prt_id") //前端发送要确认的乘客订单id
	id, _ := strconv.Atoi(prt_id)
	uid := c.Query("uid")
	err := model.ConfirmD(id, uid)
	if err != nil {
		handler.ErrServerError(c, err)
		return
	}

	c.JSON(200, gin.H{
		"msg": "success",
	})

}
