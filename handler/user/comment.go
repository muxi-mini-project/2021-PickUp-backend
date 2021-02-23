package user

import (
	handler "pickup/handler/err"
	"pickup/model"

	"github.com/gin-gonic/gin"
)

// @Summary 用户评论
// @Description 用户给予乘客和用户评分和评论
// @Tags user
// @Accept json
// @Produce json
// @Param Comment body model.Comment true "用户给予司机和乘客评分和评论,需要对应乘客和司机的id"
// @Success 200 {object} model.Res "{"msg":"success"}"
// @Failure 401 {object} handler.Error "{"error_code":"20001", "message":"Fail."} 评论失败
// @Failure 400 {object} handler.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} handler.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /users/comment [put]
func UsersComment(c *gin.Context) {
	//uid := c.Param("uid")
	var tmpcomment model.Comment
	var tmpdriver model.CommentDriver
	var tmppassenger model.CommentPassenger

	err := c.BindJSON(&tmpcomment)
	if err != nil {
		handler.ErrBadRequest(c, err)
		return
	}
	tmpdriver.DriverID = tmpcomment.DriverID
	tmpdriver.DriverScore = tmpcomment.DriverScore
	tmpdriver.Words = tmpcomment.DriverWords
	tmppassenger.PassengerID = tmpcomment.PassengerID
	tmppassenger.PassengerScore = tmpcomment.PassengerScore
	tmppassenger.Words = tmpcomment.PassengerWords

	if err := model.UpdateCommentD(tmpdriver, tmpcomment.PassengerID); err != nil {
		handler.ErrServerError(c, err)
		return
	}
	if err := model.UpdateCommentP(tmppassenger, tmpcomment.DriverID); err != nil {
		handler.ErrServerError(c, err)
		return
	}
	//fmt.Println(uid)
	c.JSON(200, gin.H{
		"msg": "success",
	})
}
