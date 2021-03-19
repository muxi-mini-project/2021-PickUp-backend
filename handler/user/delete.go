package user

import (
	handler "pickup/handler/err"
	"pickup/model"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// @Summary 注销
// @Description 注销用户用户
// @Tags user
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} model.Res "{"msg":"success"}"
// @Failure 401 {object} handler.Error "{"error_code":"20001", "message":"Fail."} 注销失败
// @Failure 500 {object} handler.Error "{"error_code":"30001", "message":"用户不存在"} 失败"
// @Router /users [delete]
func UserDelete(c *gin.Context) {
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

	if _, err := model.FindUser(uid); err != nil {
		c.JSON(404, gin.H{
			"msg": "用户不存在!",
		})
		return
	}

	if err := model.DeleteUser(uid); err != nil {
		c.JSON(401, gin.H{
			"msg": "注销失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg": "success",
	})
	return
}
