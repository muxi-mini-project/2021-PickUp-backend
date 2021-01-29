package user

import (
	handler "pickup/handler/err"
	"pickup/model"

	"github.com/gin-gonic/gin"
)

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
