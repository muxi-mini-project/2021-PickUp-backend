package user

import (
	handler "pickup/handler/err"
	"pickup/model"
	"strings"

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
		c.JSON(400, gin.H{"error_code": "00002", "message": "Lack Param Or Param Not Satisfiable."})
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

// @Summary 显示用户评论
// @Description 显示乘客和用户评分和评论
// @Tags user
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {arrey}  "{"msg":"success"}"
// @Failure 401 {object} handler.Error "{"error_code":"20001", "message":"Fail."} 评论失败
// @Failure 400 {object} handler.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} handler.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /users/comment [get]
func ViewUsersComment(c *gin.Context) {
	//uid := c.Param("uid")
	var token *model.JwtClaims
	var s string
	s = c.GetHeader("token")
	//fmt.Println(s)
	token, err2 := model.VerifyToken(s)
	//fmt.Println(token)
	if err2 != nil {
		handler.ErrTokenInvalid(c, err2)
		c.JSON(401, gin.H{"error_code": "10001", "message": "Token Invalid."})
		return
	}
	uid := token.UID

	Comment1, err1 := model.FindCommentD(uid)
	Comment2, err2 := model.FindCommentP(uid)
	var words []string

	if err1 != nil && err2 != nil {
		handler.ErrBadRequest(c, err1)
		c.JSON(401, gin.H{"error_code": "10001", "message": "Token Invalid."})
		return
	}
	if err1 != nil || err2 == nil {
		words = strings.SplitN(Comment2.Words, ",", -1)
	} else {
		words = strings.SplitN(Comment1.Words, ",", -1)
	}
	word := make(map[string]int)
	for _, value := range words {
		if value != "" {
			word[value]++
		}
	}

	//fmt.Println(uid)
	c.JSON(200, gin.H{
		"msg":              "success",
		"words_and_number": word,
	})
}

//新增举报功能
// @Summary 用户举报
// @Description 用户举报别人
// @Tags user
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param report_id path string true "用户举报别人,需要对应被举报人员的id"
// @Success 200 {object} model.Res "{"msg":"success"}"
// @Failure 401 {object} handler.Error "{"error_code":"20001", "message":"Fail."} 举报失败
// @Failure 400 {object} handler.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} handler.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /users/report/:report_id [put]
func Report(c *gin.Context) {
	report_id := c.Query("report_id")
	var token *model.JwtClaims
	var s string
	s = c.GetHeader("token")
	//fmt.Println(s)
	token, err2 := model.VerifyToken(s)
	//fmt.Println(token)
	if err2 != nil {
		handler.ErrTokenInvalid(c, err2)
		c.JSON(401, gin.H{"error_code": "10001", "message": "Token Invalid."})
		return
	}
	uid := token.UID
	_, err := model.FindUser(uid)
	if err != nil {
		handler.ErrBadRequest(c, err)
		c.JSON(400, gin.H{"error_code": "00001", "message": "Failed"})
		return
	}

	err3 := model.Deduct(report_id)
	if err3 != nil {
		//举报失败
		handler.ErrBadRequest(c, err)
		c.JSON(400, gin.H{"error_code": "00001", "message": "Failed"})
		return
	}

	c.JSON(200, gin.H{
		"msg": "success",
	})
}
