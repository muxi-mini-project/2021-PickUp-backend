package router

import (
	"pickup/handler"
	"pickup/handler/driver"
	"pickup/handler/passenger"
	"pickup/handler/route"
	"pickup/handler/user"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {

	r.GET("/text", handler.Text)

	//此函数为一站式登录函数,若用户登录成功(首次),则将其学号和密码保存起来创造一个新用户
	r.POST("/users", user.UserCreate)

	//登录
	r.POST("/users/login", user.UserLogin)

	//用户主页及其信息修改
	r.GET("/users/:uid", user.ViewUser)
	r.PUT("/users/:uid/update", user.UpdateUser)

	//用户修改密码
	r.PUT("/users/:uid/password", user.UpdatePassword)

	//用户评价
	r.PUT("/users/:uid/comment", user.UsersComment)
	//r.GET("/users/comment/{comment_id}", handler.ViewComment)

	//司机订单
	r.POST("/driver", driver.AddDriverRequirement)
	r.POST("/driver/route")
	r.GET("/driver", driver.ViewDriverRequirement)
	r.PUT("​driver/comfirm​", driver.DriverConfirm)
	r.DELETE("/driver", driver.DeleteDriverRequirement)

	//乘客订单
	r.POST("/passenger", passenger.AddPassengerRequirement)
	r.PATCH("/passenger/route", passenger.AddPRequirement)
	r.GET("/passenger", passenger.ViewPassengerRequirement)
	r.DELETE("/passenger", passenger.DeletePassengerRequirement)

	/*	//乘客确认
		r.POST("/passenger/confirm/{confirm_id}", handler.PassengerConfirm)
	*/
	//常用路径
	r.POST("/route", route.AddNewRoute)
	r.GET("/route", route.ViewRoute)
	r.DELETE("/route", route.DeleteRoute)

}
