package passenger

import (
	handler "pickup/handler/err"
	"pickup/model"

	"github.com/gin-gonic/gin"
)

func PassengerConfirm(c *gin.Context) {
	uid := c.Query("uid")
	err := model.ConfirmP(uid)
	if err != nil {
		handler.ErrServerError(c, err)
		return
	}

	c.JSON(200, gin.H{
		"msg": "success",
	})

}
