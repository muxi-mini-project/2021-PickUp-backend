package driver

import (
	handler "pickup/handler/err"
	"pickup/model"

	"github.com/gin-gonic/gin"
)

func ViewDriverRequirement(c *gin.Context) {
	uid := c.Param("uid")
	tmpRt, err := model.FindDriverRt(uid)
	if err != nil {
		handler.ErrServerError(c, err)
		return
	}
	c.JSON(200, gin.H{
		"msg":           "Success",
		"start_time":    tmpRt.StartTime,
		"end_time":      tmpRt.EndTime,
		"start_spot":    tmpRt.StartSpot,
		"passing_spots": tmpRt.PassingSpots,
		"status":        tmpRt.Status,
		"notes":         tmpRt.Notes,
	})
	return
}
