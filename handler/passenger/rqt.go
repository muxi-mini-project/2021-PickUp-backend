package passenger

import (
	handler "pickup/handler/err"
	"pickup/model"

	"github.com/gin-gonic/gin"
)

// @Summary 显示订单信息
// @Description 显示订单信息,点击 订单详情页 时调用
// @Tags passenger
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} model.VeRequirePassenger "成功"
// @Failure 401 {object} handler.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} handler.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} handler.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /passenger [get]
func ViewPassengerRequirement(c *gin.Context) {
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
	pid := token.UID
	tmpRt, err := model.FindPassengerRt(pid)
	if err != nil {
		handler.ErrBadRequest(c, err)
		c.JSON(400, gin.H{"error_code": "00002", "message": "Failed"})
		return
	}
	c.JSON(200, gin.H{
		"ymd":        tmpRt.Ymd,
		"start_time": tmpRt.StartTime,
		"end_time":   tmpRt.EndTime,
		"start_spot": tmpRt.StartSpot,
		"end_spot":   tmpRt.EndSpot,
		"status":     ChangeStatus(tmpRt.Status),
		"notes":      tmpRt.Notes,
		"urgent":     tmpRt.Urgent,
		"phone":      model.GetPhone(pid),
	})
	return
}

// @Summary 通过常用路径来添加部分订单信息
// @Description 在路径中,点击 创建新订单 时调用
// @Tags passenger
// @Accept json
// @Produce json
// @Param id path int true "常用路径的id"
// @Param token header string true "token"
// @Success 200 {object} model.VeRoute "成功"
// @Failure 401 {object} handler.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} handler.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} handler.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /passenger/:id [get]
//通过常用路线显示将要发布的订单部分信息
func ViewPRequirement(c *gin.Context) {
	id := c.GetInt("id")
	route, err := model.FindRoute2(id)
	if err != nil {
		handler.ErrBadRequest(c, err)
		c.JSON(400, gin.H{"error_code": "00002", "message": "Lack Param Or Param Not Satisfiable."})
		return
	}
	user, err2 := model.FindUser(route.UserID)
	if err2 != nil {
		handler.ErrBadRequest(c, err)
		c.JSON(400, gin.H{"error_code": "00002", "message": "Lack Param Or Param Not Satisfiable."})
		return
	}
	c.JSON(200, gin.H{
		"msg":        "success",
		"start_spot": route.StartSpot,
		"end_spot":   route.EndSpot,
		"user_phone": user.Phone,
	})

}

func ChangeStatus(status int) string {
	if status == 1 {
		return "待确认"
	}

	if status == 2 {
		return "已确认"
	}

	return "获取信息失败"
}
