package route

import (
	handler "pickup/handler/err"
	"pickup/model"

	"github.com/gin-gonic/gin"
)

// @Summary 用户路径
// @Description 用户删除路径
// @Tags route
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param route_id path int true "路径id"
// @Success 200 {object} model.Res "{"msg":"success"}"
// @Failure 401 {object} handler.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} handler.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} handler.Error "{"error_code":"30001", "message":"database does not open successful"} 失败"
// @Router /route/:route_id [delete]
func DeleteRoute(c *gin.Context) {
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
	id := c.GetInt("id")
	uid := token.UID
	route, err3 := model.FindRoute2(id)
	if err3 != nil {
		handler.ErrBadRequest(c, err3)
		return
	}

	if route.UserID != uid {
		handler.ErrBadRequest(c, err2)
		return
	}
	if err := model.DeleteRoute(id); err != nil {
		handler.ErrServerError(c, err)
		return
	}

	c.JSON(200, gin.H{
		"msg": "success",
	})
}
