package driver

import (
	handler "pickup/handler/err"
	"pickup/model"

	"github.com/gin-gonic/gin"
)

// @Summary 司机确认
// @Description 司机对接受到的订单,发出确认,注意,只能确认一个
// @Tags driver
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param passenger_id path string true "乘客的id"
// @Success 200 {object} model.Res "{"msg":"success","pid":"string","did":"string"}" 如果成功,则将对应司机和乘客的订单存到匹配成功的数据库中,并且删除这两个订单请求.
// @Failure 401 {object} handler.Error "{"error_code":"20001", "message":"Fail."} 确认失败
// @Failure 400 {object} handler.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} handler.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /passenger/confirm [put]
func DriverConfirm(c *gin.Context) {
	pid := c.Query("pid") //前端发送要确认的乘客id
	var token *model.JwtClaims
	var s string
	s = c.GetHeader("token")
	//fmt.Println(s)
	token, err2 := model.VerifyToken(s)
	//fmt.Println(token)
	if err2 != nil {
		handler.ErrTokenInvalid(c, err2)
		c.JSON(401, gin.H{"error_code": "10001", "message": "Token Invalid."})
		return
	}

	if _, err := model.FindUser(token.UID); err != nil {
		handler.ErrBadRequest(c, err)
		c.JSON(400, gin.H{"error_code": "00001", "message": "Fail."})
		return
	}

	did := token.UID
	err := model.ConfirmD(pid, did)
	if err != nil {
		handler.ErrBadRequest(c, err)
		c.JSON(400, gin.H{"error_code": "00001", "message": "Fail."})
		return
	}

	c.JSON(200, gin.H{
		"msg": "success",
		"pid": pid,
		"did": did,
	})

}
