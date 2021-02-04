package passenger

import (
	handler "pickup/handler/err"
	"pickup/model"

	"github.com/gin-gonic/gin"
)

//直接发布订单
func AddPassengerRequirement(c *gin.Context) {
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
	var tmprequirement model.RequirePassenger
	tmprequirement.PassengerID = token.UID
	if err := c.BindJSON(&tmprequirement); err != nil {
		handler.ErrBadRequest(c, err)
		//fmt.Println("1")
		return
	}
	if err := model.CreatePassengerRt(tmprequirement); err != nil {
		handler.ErrServerError(c, err)
		//fmt.Println("2")
		return
	}

	rt, err := model.FindPassengerRt(tmprequirement.PassengerID)
	if err != nil {
		handler.ErrUnauthorized(c, err)
		return
	}

	c.JSON(200, gin.H{
		"msg": "success",
		"drt": rt.ID,
	})
}
