package user

import (
	handler "pickup/handler/err"
	"pickup/model"

	"github.com/gin-gonic/gin"
)

// @Summary 显示匹配度
// @Description 显示乘客和用户,在乘客或司机查看订单,进行确认的时候进行调用
// @Tags user
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param pid path string true "Passenger_id"
// @Param did path string true "driver_id"
// @Success 200 {array} model.Route "成功"
// @Failure 401 {object} handler.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} handler.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} handler.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /user/match/:pid/did [get]
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
