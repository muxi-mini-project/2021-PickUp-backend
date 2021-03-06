package driver

import (
	handler "pickup/handler/err"
	"pickup/model"

	"github.com/gin-gonic/gin"
)

// @Summary 司机订单
// @Description 司机删除订单
// @Tags driver
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} model.Res "{"msg":"success"}"
// @Failure 401 {object} handler.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} handler.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} handler.Error "{"error_code":"30001", "message":"database does not open successful"} 失败"
// @Router /driver [delete]
func DeleteDriverRequirement(c *gin.Context) {
	var token *model.JwtClaims
	var s string
	s = c.GetHeader("token")
	//fmt.Println(s)
	token, err2 := model.VerifyToken(s)
	//fmt.Println(token)
	if err2 != nil {
		handler.ErrTokenInvalid(c, err2)
		return
	}
	uid := token.UID

	err := model.DeleteDriverRt(uid)
	if err != nil {
		handler.ErrUnauthorized(c, err)
		return
	}

	c.JSON(200, gin.H{
		"msg": "success",
	})

}
