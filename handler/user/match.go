package user

import (
	handler "pickup/handler/err"
	"pickup/model"

	"github.com/gin-gonic/gin"
)

func MatchDegree(c *gin.Context) {
	pid := c.Query("pid")
	did := c.Query("did")
	tmpRt, err := model.FindPassengerRt(pid)
	if err != nil {
		handler.ErrServerError(c, err)
		return
	}
	tmprt, err2 := model.FindDriverRt(did)
	if err2 != nil {
		handler.ErrServerError(c, err2)
	}

	c.JSON(200, gin.H{
		"match_degree": model.PercentChange(model.MatchDegree(tmpRt, tmprt)),
	})
}
