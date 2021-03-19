package route

import (
	handler "pickup/handler/err"
	"pickup/model"

	"github.com/gin-gonic/gin"
)

// @Summary 常用路径
// @Description 用户创造新的常用路径
// @Tags route
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param tmprequirement body model.Route true "用户需要填写的信息"
// @Success 200 {object} model.Res "{"msg":"success", "pid":"string"}"
// @Failure 401 {object} handler.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} handler.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} handler.Error "{"error_code":"30001", "message":"database does not open successful"} 失败"
// @Router /route [post]
func AddNewRoute(c *gin.Context) {
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
	var tmproute model.Route
	tmproute.UserID = uid
	if err := c.BindJSON(&tmproute); err != nil {
		handler.ErrBadRequest(c, err)
		c.JSON(400, gin.H{"error_code": "00002", "message": "Lack Param Or Param Not Satisfiable."})
		return
	}
	route := model.Match{
		StartSpot: tmproute.StartSpot,
		EndSpot:   tmproute.EndSpot,
		StartTime: tmproute.StartTime,
		EndTime:   tmproute.EndTime,
	}
	if err := model.CreateRoute(route); err != nil {
		handler.ErrBadRequest(c, err)
		c.JSON(400, gin.H{"error_code": "00002", "message": "Failed"})
		return
	}

	c.JSON(200, gin.H{
		"msg": "success",
	})

}
