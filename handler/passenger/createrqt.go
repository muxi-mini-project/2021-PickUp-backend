package passenger

import (
	handler "pickup/handler/err"
	"pickup/model"

	"github.com/gin-gonic/gin"
)

// @Summary 乘客订单
// @Description 乘客创造新的订单
// @Tags passenger
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param tmprequirement body model.VeRequirePassenger true "乘客需要填写的信息,注意年月日需要以xx年xx月xx日的形式填写,status表示该订单是否完成,1为未完成,2为已完成"
// @Success 200 {object} model.Res "{"msg":"success", "pid":"string"}"
// @Failure 401 {object} handler.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} handler.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} handler.Error "{"error_code":"30001", "message":"database does not open successful"} 失败"
// @Router /passenger [post]
func AddPassengerRequirement(c *gin.Context) {
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
	var tmprequirement model.RequirePassenger
	tmprequirement.PassengerID = token.UID
	if err := c.BindJSON(&tmprequirement); err != nil {
		handler.ErrBadRequest(c, err)
		c.JSON(400, gin.H{"error_code": "00002", "message": "Lack Param Or Param Not Satisfiable."})
		//fmt.Println("1")
		return
	}
	if err := model.CreatePassengerRt(tmprequirement); err != nil {
		handler.ErrBadRequest(c, err)
		c.JSON(400, gin.H{"error_code": "00002", "message": "Failed"})
		//fmt.Println("2")
		return
	}

	rt, err := model.FindPassengerRt(tmprequirement.PassengerID)
	if err != nil {
		handler.ErrBadRequest(c, err)
		c.JSON(400, gin.H{"error_code": "00001", "message": "Failed"})
		return
	}

	c.JSON(200, gin.H{
		"msg": "success",
		"pid": rt.PassengerID,
	})
}
