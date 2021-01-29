package route

import (
	handler "pickup/handler/err"
	"pickup/model"

	"github.com/gin-gonic/gin"
)

func DeleteRoute(c *gin.Context) {
	id := c.GetInt("id")
	if err := model.DeleteRoute(id); err != nil {
		handler.ErrServerError(c, err)
		return
	}
	c.JSON(200, gin.H{
		"msg": "success",
	})
}
