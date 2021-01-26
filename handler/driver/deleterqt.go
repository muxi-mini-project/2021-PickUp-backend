package driver

import (
	handler "pickup/handler/err"
	"pickup/model"

	"github.com/gin-gonic/gin"
)

func DeleteDriverRequirement(c *gin.Context) {
	uid := c.Param("uid")

	err := model.DeleteDriverRt(uid)
	if err != nil {
		handler.ErrUnauthorized(c, err)
		return
	}

}
