package user

/*import (
	"pickup/model"

	"github.com/gin-gonic/gin"
)

// @Summary 登录
// @Description Login
// @Tags user
// @Accept json
// @Produce json
// @Param loginInfo body model.LoginInfo true "学号和密码"
// @Success 200 {object} model.Res "{"msg":"success", "token": string}"
// @Failure 401 {object} handler.Error "{"error_code":"20001", "message":"user does not exist"} 登录失败 or {"error_code":"20002", "message":"password does not correct."}
// @Failure 400 {object} handler.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} handler.Error "{"error_code":"30001", "message":"database does not open successful"} 失败"
// @Router /users/login [post]
func UserLogin(c *gin.Context) {
	var tmpLoginInfo model.LoginInfo

	if err := c.BindJSON(&tmpLoginInfo); err != nil {
		c.JSON(400, gin.H{
			"msg": "Lack Param Or Param Not Satisfiable.",
		})
		return
	}

	ok := model.Login(tmpLoginInfo)
	if ok == 1 {
		c.JSON(500, gin.H{
			"msg": "database does not open successful",
		})
		return
	}

	if ok == 2 {
		c.JSON(401, gin.H{
			"msg": "user does not exist",
		})
		return
	}

	if ok == 3 {
		c.JSON(401, gin.H{
			"msg": "password does not correct",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg":   "success",
		"token": produceToken(tmpLoginInfo.Sid),
	})

}
*/
