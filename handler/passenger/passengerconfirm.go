package passenger

import (
	handler "pickup/handler/err"
	"pickup/model"

	"github.com/gin-gonic/gin"
)

// @Summary 乘客确认
// @Description 乘客对司机发出请求
// @Tags passenger
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param driver_id path string true "司机的id"
// @Success 200 {object} model.Res "{"msg":"success","pid":"string","did":"string"}" 如果成功,则发送给对应司机乘客的id
// @Failure 401 {object} handler.Error "{"error_code":"20001", "message":"Fail."} 确认失败
// @Failure 400 {object} handler.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} handler.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /passenger/confirm [put]
func PassengerConfirm(c *gin.Context) {
	did := c.Query("driver_id")
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
	uid := token.UID
	err := model.ConfirmP(uid)
	if err != nil {
		handler.ErrBadRequest(c, err)
		c.JSON(400, gin.H{"error_code": "00002", "message": "Lack Param Or Param Not Satisfiable."})
		return
	}

	c.JSON(200, gin.H{
		"msg": "success",
		"pid": uid,
		"did": did,
	})

}
