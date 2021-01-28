package passenger

import (
	handler "pickup/handler/err"
	"pickup/model"

	"github.com/gin-gonic/gin"
)

func DeletePassengerRequirement(c *gin.Context) {
	uid := c.Param("uid")

	err := model.DeletePassengerRt(uid)
	if err != nil {
		handler.ErrUnauthorized(c, err)
		return
	}

	c.JSON(200, gin.H{
		"msg": "success",
	})

}
