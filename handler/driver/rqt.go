package driver

import (
	handler "pickup/handler/err"
	"pickup/model"

	"github.com/gin-gonic/gin"
)

func ViewDriverRequirement(c *gin.Context) {
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
	did := token.UID
	tmpRt, err2 := model.FindDriverRt(did)
	if err2 != nil {
		handler.ErrServerError(c, err2)
	}
	c.JSON(200, gin.H{
		"msg":           "Success",
		"ymd":           tmpRt.Ymd,
		"start_time":    tmpRt.StartTime,
		"end_time":      tmpRt.EndTime,
		"start_spot":    tmpRt.StartSpot,
		"passing_spots": tmpRt.PassingSpots,
		"status":        tmpRt.Status,
		"notes":         tmpRt.Notes,
		"phone":         model.GetPhone(did),
	})
	return
}

//通过常用路线发布订单
func ViewDRequirement(c *gin.Context) {
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
