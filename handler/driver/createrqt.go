package driver

import (
	handler "pickup/handler/err"
	"pickup/model"

	"github.com/gin-gonic/gin"
)

// @Summary 司机订单
// @Description 司机创造新的订单
// @Tags driver
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param tmprequirement body model.VeRequireDriver true "司机需要填写的信息,注意年月日需要以xx年xx月xx日的形式填写,status表示该订单是否完成,1为未完成,2为已完成"
// @Success 200 {object} model.Res "{"msg":"success", "drt":"string"}"
// @Failure 401 {object} handler.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} handler.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} handler.Error "{"error_code":"30001", "message":"database does not open successful"} 失败"
// @Router /driver [post]
func AddDriverRequirement(c *gin.Context) {

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
		c.JSON(401, gin.H{"error_code": "00001", "message": "Fail."})
		return
	}

	var tmprequirement model.RequireDriver

	tmprequirement.DriverID = token.UID
	tmprequirement.Status = 1
	if err := c.BindJSON(&tmprequirement); err != nil {
		handler.ErrBadRequest(c, err)
		c.JSON(400, gin.H{"error_code": "00001", "message": "Fail."})
		return
	}
	if err := model.CreateDriverRt(tmprequirement); err != nil {
		handler.ErrServerError(c, err)
		c.JSON(500, gin.H{"error_code": "30001", "message": "Fail."})
		return
	}

	rt, err := model.FindDriverRt(tmprequirement.DriverID)
	if err != nil {
		handler.ErrUnauthorized(c, err)
		c.JSON(400, gin.H{"error_code": "30001", "message": "Fail."})
		return
	}

	c.JSON(200, gin.H{
		"msg": "success",
		"drt": rt.DriverID,
	})
}
