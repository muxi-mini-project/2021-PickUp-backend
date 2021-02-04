package passenger

import (
	handler "pickup/handler/err"
	"pickup/model"

	"github.com/gin-gonic/gin"
)

func ViewPassengerRequirement(c *gin.Context) {
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
	pid := token.UID
	tmpRt, err := model.FindPassengerRt(pid)
	if err != nil {
		handler.ErrServerError(c, err)
		return
	}
	c.JSON(200, gin.H{
		"msg":        "Success",
		"ymd":        tmpRt.Ymd,
		"start_time": tmpRt.StartTime,
		"end_time":   tmpRt.EndTime,
		"start_spot": tmpRt.StartSpot,
		"end_spot":   tmpRt.EndSpot,
		"status":     tmpRt.Status,
		"notes":      tmpRt.Notes,
		"urgent":     tmpRt.Urgent,
		"phone":      model.GetPhone(pid),
	})
	return
}

//通过常用路线显示将要发布的订单部分信息
func ViewPRequirement(c *gin.Context) {
	id := c.GetInt("id")
	route, err := model.FindRoute2(id)
	if err != nil {
		handler.ErrBadRequest(c, err)
		return
	}
	user, err2 := model.FindUser(route.UserID)
	if err2 != nil {
		handler.ErrBadRequest(c, err2)
		return
	}
	c.JSON(200, gin.H{
		"msg":        "success",
		"start_spot": route.StartSpot,
		"end_spot":   route.EndSpot,
		"user_phone": user.Phone,
	})

}
