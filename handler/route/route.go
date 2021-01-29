package route

import (
	handler "pickup/handler/err"
	"pickup/model"

	"github.com/gin-gonic/gin"
)

func ViewRoute(c *gin.Context) {
	uid := c.Query("uid")
	routes, err := model.FindRoute(uid)
	if err != nil {
		handler.ErrBadRequest(c, err)
		return
	}
	c.JSON(200, routes)

}
