package route

import (
	handler "pickup/handler/err"
	"pickup/model"

	"github.com/gin-gonic/gin"
)

func ViewRoute(c *gin.Context) {
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
	routes, err := model.FindRoute(uid)
	if err != nil {
		handler.ErrBadRequest(c, err)
		return
	}
	c.JSON(200, routes)

}
