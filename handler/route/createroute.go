package route

import (
	handler "pickup/handler/err"
	"pickup/model"

	"github.com/gin-gonic/gin"
)

func AddNewRoute(c *gin.Context) {
	var tmproute model.Route
	if err := c.BindJSON(&tmproute); err != nil {
		handler.ErrBadRequest(c, err)
		return
	}
	route := model.Match{
		UserID:    tmproute.UserID,
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
