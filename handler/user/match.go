package user

import (
	handler "pickup/handler/err"
	"pickup/model"

	"github.com/gin-gonic/gin"
)

// @Summary 显示匹配度
// @Description 显示司机与乘客的匹配度和司机订单信息,在乘客查看订单,进行确认的时候进行调用
// @Tags user
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {array} model.VRequireDriver "成功"
// @Failure 401 {object} handler.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} handler.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} handler.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /user/recommend [get]
//向客户推荐司机时使用
func Recommend(c *gin.Context) {
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
	tmpuser, err := model.FindPassengerRt(uid)
	if err != nil {
		handler.ErrBadRequest(c, err)
		c.JSON(400, gin.H{"error_code": "00001", "message": "Failed"})
		return
	}

	Drt, err3 := model.FindDriverRt2(tmpuser)
	if err3 != nil {
		handler.ErrBadRequest(c, err)
		c.JSON(400, gin.H{"error_code": "00001", "message": "Failed"})
		return
	}
	model.EachMap(Drt, func(key int, value model.RequireDriver) {
		c.JSON(200, gin.H{
			"Drt":   value,
			"match": key,
		})
	})

}
