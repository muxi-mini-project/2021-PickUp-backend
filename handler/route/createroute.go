package route

import (
	handler "pickup/handler/err"
	"pickup/model"

	"github.com/gin-gonic/gin"
)

func AddNewRoute(c *gin.Context) {
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
	var tmproute model.Route
	tmproute.UserID = uid
	if err := c.BindJSON(&tmproute); err != nil {
		handler.ErrBadRequest(c, err)
		return
	}
	route := model.Match{
		StartSpot: tmproute.StartSpot,
		EndSpot:   tmproute.EndSpot,
		StartTime: tmproute.StartTime,
		EndTime:   tmproute.EndTime,
	}
	if err := model.CreateRoute(route); err != nil {
		handler.ErrServerError(c, err)
		return
	}

	c.JSON(200, gin.H{
		"msg": "success",
	})

}
