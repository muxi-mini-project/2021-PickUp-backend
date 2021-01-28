package driver

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func DriverConfirm(c *gin.Context) {
	prt_id := c.Param("prt_id") //前端发送要确认的乘客订单id

	fmt.Println(prt_id)

}
