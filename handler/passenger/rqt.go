package passenger

import (
	handler "pickup/handler/err"
	"pickup/model"

	"github.com/gin-gonic/gin"
)

func ViewPassengerRequirement(c *gin.Context) {
	uid := c.Param("uid")
	tmpRt, err := model.FindPassengerRt(uid)
	if err != nil {
		handler.ErrServerError(c, err)
		return
	}
	c.JSON(200, gin.H{
		"msg":        "Success",
		"start_time": tmpRt.StartTime,
		"end_time":   tmpRt.EndTime,
		"start_spot": tmpRt.StartSpot,
		"end_spot":   tmpRt.EndSpot,
		"status":     tmpRt.Status,
		"notes":      tmpRt.Notes,
		"urgent":     tmpRt.Urgent,
	})
	return
}
